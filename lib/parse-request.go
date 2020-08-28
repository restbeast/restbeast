package lib

import (
	"encoding/json"
	"errors"
	. "fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

var dependencyDiagMessageRegex = regexp.MustCompile(`This object does not have an attribute named "(?P<name>[\w\d-_]+)"`)
var requestDependencyRegex = regexp.MustCompile(`^request.([\w\d-_]+)`)

func getObjSpec() hcldec.ObjectSpec {
	return hcldec.ObjectSpec{
		"method": &hcldec.AttrSpec{
			Name:     "method",
			Required: true,
			Type:     cty.String,
		},
		"url": &hcldec.AttrSpec{
			Name:     "url",
			Required: true,
			Type:     cty.String,
		},
		"headers": &hcldec.AttrSpec{
			Name:     "headers",
			Required: false,
			Type:     cty.Map(cty.String),
		},
		"body": &hcldec.AttrSpec{
			Name:     "body",
			Required: false,
			Type:     cty.DynamicPseudoType,
		},
		"depends_on": &hcldec.AttrSpec{
			Name:     "depends_on",
			Required: false,
			Type:     cty.List(cty.String),
		},
	}
}

// Find a request by .Name property
func findRequest(name string, rawRequests RequestCfgs) (err error, request RequestCfg) {
	for _, r := range rawRequests {
		if name == r.Name {
			return nil, *r
		}
	}

	return Errorf("`%s` not found", name), RequestCfg{}
}

func getCtxEvalContext(evCtx EvalContext) hcl.EvalContext {
	vars := map[string]cty.Value{
		"var":     cty.ObjectVal(evCtx.Variables),
		"request": cty.ObjectVal(evCtx.RequestAsVars),
	}

	if evCtx.Environment != nil {
		vars["env"] = *evCtx.Environment
	}

	return hcl.EvalContext{
		Variables: vars,
		Functions: evCtx.Functions,
	}
}

func getPossibleDependencies(diags hcl.Diagnostics) (dependencies []string, restDiagMsgs []string) {
	if len(diags) != 0 {
		for _, diag := range diags {
			if diag.Summary == "Unsupported attribute" {
				findString := dependencyDiagMessageRegex.FindStringSubmatch(diag.Detail)

				if len(findString) > 1 {
					dependencies = append(dependencies, findString[1])
				} else {
					restDiagMsgs = append(restDiagMsgs, Sprint(diag))
				}
			}
		}
	}

	return dependencies, restDiagMsgs
}

func getUniqueDependencies(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// TODO: This needs a lot more thinking
// Right now it can handle chained dependencies based on dependency count
// This obviously might cause unexpected dependency problems
func sortCrossDependency(deps []string, evCtx EvalContext) ([]string, error) {
	if len(deps) < 2 {
		return deps, nil
	}

	depsWithDeps := make(map[string]int)

	for _, dep := range deps {
		ctxEvalContext := getCtxEvalContext(evCtx)
		spec := getObjSpec()

		err, request := findRequest(dep, evCtx.RawRequests)

		if err != nil {
			return nil, err
		}

		_, diags := hcldec.Decode(request.Body, spec, &ctxEvalContext)
		dependencies, _ := getPossibleDependencies(diags)

		var relevantDeps []string
		for _, innerDep := range dependencies {
			if sliceContains(deps, innerDep) {
				relevantDeps = append(relevantDeps, innerDep)
			}
		}

		depsWithDeps[dep] = len(relevantDeps)
	}

	sort.Slice(deps, func(i, j int) bool {
		return depsWithDeps[deps[i]] > depsWithDeps[deps[j]]
	})

	return deps, nil
}

// Make headers accessible with lowercase variations
func lowercaseHeaders(headers http.Header) http.Header {
	for k, v := range headers {
		headers[strings.ToLower(k)] = v
	}

	return headers
}

func processDependency(dependency string, evCtx *EvalContext, execCtx ExecutionContext) (*EvalContext, *Response, error) {
	request, parseErr := parseRequest(dependency, *evCtx, execCtx)
	if parseErr != nil {
		return nil, nil, parseErr
	}

	response, requestErr := DoRequest(*request, &execCtx)
	if requestErr != nil {
		return nil, nil, requestErr
	}

	var decodedBody interface{}
	err := json.Unmarshal(response.Body, &decodedBody)

	if err != nil {
		return nil, nil, Errorf("error decoding json response body\n%s\n", err)
	}

	var responseAsCty = map[string]cty.Value{}
	responseAsCty["body"] = walkThrough(reflect.ValueOf(decodedBody))

	convertedHeaders := lowercaseHeaders(response.Headers)
	headersAsCty := walkThrough(reflect.ValueOf(convertedHeaders))

	responseAsCty["headers"] = headersAsCty
	responseAsCty["status"] = cty.NumberIntVal(int64(response.StatusCode))

	evCtx.RequestAsVars[dependency] = cty.ObjectVal(responseAsCty)

	return evCtx, response, nil
}

func getRequest(cfg cty.Value) Request {
	body, jsonErr := json.MarshalIndent(ctyjson.SimpleJSONValue{cfg.GetAttr("body")}, "", "  ")
	if jsonErr != nil {
		Printf("Error: failed to parse request body, \n%s\n", jsonErr)
		os.Exit(1)
	}

	var headers map[string]string
	headerErr := gocty.FromCtyValue(cfg.GetAttr("headers"), &headers)

	if headerErr != nil {
		Printf("Error: failed to parse headers, \n%s\n", headerErr)
		os.Exit(1)
	}

	var method string
	if !cfg.GetAttr("method").IsNull() {
		method = cfg.GetAttr("method").AsString()
	} else {
		method = "GET"
	}

	if !cfg.GetAttr("url").IsWhollyKnown() {
		Printf("Error: failed to parse url, possible unknown variable used.\n")
		os.Exit(1)
	}

	request := Request{
		Method:  method,
		Url:     cfg.GetAttr("url").AsString(),
		Headers: headers,
		Body:    string(body),
	}

	return request
}

func parseRequest(name string, evCtx EvalContext, execCtx ExecutionContext) (*Request, error) {
	err, request := findRequest(name, evCtx.RawRequests)

	if err != nil {
		return nil, err
	}

	var responses []*Response

	if request.DependsOn != nil {
		for _, v := range request.DependsOn {
			findString := requestDependencyRegex.FindStringSubmatch(v)

			if len(findString) > 1 {
				if _, ok := evCtx.RequestAsVars[findString[1]]; !ok {
					evCtxP, response, err := processDependency(findString[1], &evCtx, execCtx)
					responses = append(responses, response)

					if err != nil {
						return nil, err
					}

					evCtx = *evCtxP
				}
			}
		}
	}

	ctxEvalContext := getCtxEvalContext(evCtx)
	spec := getObjSpec()
	cfg, diags := hcldec.Decode(request.Body, spec, &ctxEvalContext)
	dependencies, restDiagMsgs := getPossibleDependencies(diags)

	if len(restDiagMsgs) > 0 {
		errTxt := ""
		for _, diag := range restDiagMsgs {
			errTxt += Sprintf("- %s\n", diag)
		}

		return nil, errors.New(errTxt)
	}

	if len(dependencies) > 0 {
		uniqueDeps := getUniqueDependencies(dependencies)
		sortedDeps, err := sortCrossDependency(uniqueDeps, evCtx)

		if err != nil {
			return nil, err
		}

		for _, dependency := range sortedDeps {
			if _, ok := evCtx.RequestAsVars[dependency]; !ok {
				evCtxP, response, err := processDependency(dependency, &evCtx, execCtx)
				responses = append(responses, response)

				if err != nil {
					return nil, err
				}

				evCtx = *evCtxP
			}
		}

		ctxEvalContext = getCtxEvalContext(evCtx)
		cfg, diags = hcldec.Decode(request.Body, spec, &ctxEvalContext)

		if len(diags) > 0 {
			errTxt := ""
			for _, diag := range diags {
				errTxt += Sprintf("- %s\n", diag)
			}

			return nil, errors.New(errTxt)
		}
	}

	finalRequest := getRequest(cfg)
	finalRequest.PrecedingRequests = responses

	return &finalRequest, nil
}
