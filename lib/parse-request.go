package lib

import (
	"encoding/json"
	. "fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/gocty"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

var dependencyDiagMessageRegex = regexp.MustCompile(`This object does not have an attribute named "(?P<name>[\w\d-_]+)"`)
var requestDependencyRegex = regexp.MustCompile(`^request.([\w\d-_]+)`)

func getRequestObjSpec() hcldec.ObjectSpec {
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
		"cookies": &hcldec.AttrSpec{
			Name:     "cookies",
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
func findRequest(name string, rawRequests RequestCfgs) (*RequestCfg, error) {
	for _, r := range rawRequests {
		if name == r.Name {
			return r, nil
		}
	}

	return nil, Errorf("`%s` not found", name)
}

func getCtxEvalContext(evCtx EvalContext) hcl.EvalContext {
	var vars map[string]cty.Value

	if evCtx.Variables != nil {
		vars = map[string]cty.Value{
			"var":     cty.ObjectVal(*evCtx.Variables),
			"request": cty.ObjectVal(evCtx.RequestAsVars),
		}
	} else {
		vars = map[string]cty.Value{
			"vars":    cty.EmptyObjectVal,
			"request": cty.ObjectVal(evCtx.RequestAsVars),
		}
	}

	if evCtx.Environment != nil {
		vars["env"] = *evCtx.Environment
	}

	fns := map[string]function.Function{}
	if evCtx.Functions != nil {
		fns = *evCtx.Functions
	}

	return hcl.EvalContext{
		Variables: vars,
		Functions: fns,
	}
}

// Read the tea leaves
// Try to find dependencies from error messages.
// HCL diag will have "Unsupported attribute" as summary and
// detail wil be "This object does not have an attribute named xxx" for unknown attributes.
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
			} else if diag.Summary != "" {
				restDiagMsgs = append(restDiagMsgs, Sprint(diag))
			}
		}
	}

	return dependencies, restDiagMsgs
}

// Traverse dependencies
// Remove duplicates
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
		spec := getRequestObjSpec()

		request, findRequestErr := findRequest(dep, evCtx.RawRequests)

		if findRequestErr != nil {
			return nil, findRequestErr
		}

		// We don't care about the decoded body at this point
		_, diags := hcldec.Decode(request.Remain, spec, &ctxEvalContext)
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

func processResponseBody(contentType *string, body []byte) (cty.Value, error) {
	switch true {
	case len(body) == 0:
		return cty.Value{}, nil
	case contentType == nil || strings.HasPrefix(*contentType, "text/plain"):
		return cty.StringVal(string(body)), nil
	case strings.HasPrefix(*contentType, "application/json"):
		var decodedBody interface{}
		err := json.Unmarshal(body, &decodedBody)
		if err != nil {
			return cty.Value{}, Errorf("error decoding json response body for request\n%s\nResponse body:\n%s\n", err, string(body))
		}
		return walkThrough(reflect.ValueOf(decodedBody)), nil
	}

	return cty.Value{}, nil
}

func processDependency(dependency string, evCtx *EvalContext, execCtx *ExecutionContext) (*EvalContext, *Response, error) {
	repeatCount, parseErr := getRequestRepeatCount(dependency, *evCtx)
	if parseErr != nil {
		return nil, nil, parseErr
	}

	if repeatCount == 0 {
		repeatCount = 1
	}

	var request *Request

	// Request loaded and executed X repeat count
	// Only the last response will be assigned to request
	for i := 0; i < repeatCount; i++ {
		request, parseErr = parseRequest(dependency, *evCtx, execCtx)
		if parseErr != nil {
			return nil, nil, parseErr
		}

		requestErr := request.Exec()
		if requestErr != nil {
			return nil, nil, requestErr
		}
	}

	var responseAsCty = map[string]cty.Value{}
	headersAsCty := walkThrough(reflect.ValueOf(request.Response.Headers.GetAll()))
	bodyAsCty, err := processResponseBody(request.Response.Headers.Get("content-type"), request.Response.Body)
	if err != nil {
		return nil, nil, err
	}
	responseAsCty["body"] = bodyAsCty
	responseAsCty["headers"] = headersAsCty
	responseAsCty["status"] = cty.NumberIntVal(int64(request.Response.StatusCode))

	evCtx.RequestAsVars[dependency] = cty.ObjectVal(responseAsCty)

	return evCtx, request.Response, nil
}

func getHeadersAsMap(cfg cty.Value) (map[string]string, error) {
	headersMap := map[string]string{}
	if cfg.Type().HasAttribute("headers") {
		headerErr := gocty.FromCtyValue(cfg.GetAttr("headers"), &headersMap)
		if headerErr != nil {
			return headersMap, Errorf("Error: failed to parse headers, \n%s\n", headerErr)
		}
	}

	return headersMap, nil
}

func getCookiesAsMap(cfg cty.Value) (map[string]string, error) {
	cookiesMap := map[string]string{}
	if cfg.Type().HasAttribute("cookies") {
		cookieErr := gocty.FromCtyValue(cfg.GetAttr("cookies"), &cookiesMap)
		if cookieErr != nil {
			return cookiesMap, Errorf("Error: failed to parse cookies, \n%s\n", cookieErr)
		}
	}

	return cookiesMap, nil
}

func getRequest(cfg cty.Value, requestCfg RequestCfg, evCtx EvalContext, execCtx *ExecutionContext) (*Request, error, hcl.Diagnostics) {
	headers := Headers{}
	headersAsMap, err := getHeadersAsMap(cfg)
	if err != nil {
		return nil, err, nil
	}
	headers.AddBulk(headersAsMap)

	cookiesAsMap, err := getCookiesAsMap(cfg)
	if err != nil {
		return nil, err, nil
	}
	headers.SetCookies(cookiesAsMap)

	var body io.Reader
	if cfg.Type().HasAttribute("body") {
		var bodyError error
		body, bodyError = parseBody(cfg.GetAttr("body"), headers)
		if bodyError != nil {
			return nil, bodyError, nil
		}
	}

	roundTripper := http.DefaultTransport
	request := &Request{
		Headers:          headers,
		Body:             body,
		ExecutionContext: execCtx,
		RoundTripper:     roundTripper,
	}

	var method string
	if !cfg.GetAttr("method").IsNull() {
		method = cfg.GetAttr("method").AsString()
	}
	request.SetMethod(method)

	authBlockDiags := parseAuthBlock(request, requestCfg.Auth, getCtxEvalContext(evCtx))
	if authBlockDiags != nil {
		return nil, nil, authBlockDiags
	}

	paramsBlockDiags := parseParamsBlock(request, requestCfg.Params, getCtxEvalContext(evCtx))
	if paramsBlockDiags != nil {
		return nil, nil, paramsBlockDiags
	}

	// URL has to be set AFTER params processed
	if !cfg.GetAttr("url").IsWhollyKnown() {
		return nil, Errorf("Error: failed to parse url, possible unknown variable used.\n"), nil
	}
	url := cfg.GetAttr("url").AsString()
	request.SetUrl(url)

	return request, nil, nil
}

func retryWithDependency(requestCfg *RequestCfg, cfg cty.Value, diags hcl.Diagnostics, evCtx EvalContext, execCtx *ExecutionContext, responses []*Response) (cty.Value, []*Response, error) {
	dependencies, restDiagMsgs := getPossibleDependencies(diags)

	if len(restDiagMsgs) > 0 {
		errTxt := ""
		for _, diag := range restDiagMsgs {
			errTxt += Sprintf("- %s\n", diag)
		}

		return cfg, responses, Errorf(errTxt)
	}

	if len(dependencies) > 0 {
		uniqueDeps := getUniqueDependencies(dependencies)
		sortedDeps, err := sortCrossDependency(uniqueDeps, evCtx)

		if err != nil {
			return cfg, responses, err
		}

		for _, dependency := range sortedDeps {
			if _, ok := evCtx.RequestAsVars[dependency]; !ok {
				evCtxP, response, depErr := processDependency(dependency, &evCtx, execCtx)
				if depErr != nil {
					return cfg, responses, depErr
				}

				responses = append(responses, response)
				evCtx = *evCtxP
			}
		}

		spec := getRequestObjSpec()
		ctxEvalContext := getCtxEvalContext(evCtx)
		cfg, diags = hcldec.Decode(requestCfg.Remain, spec, &ctxEvalContext)

		if len(diags) > 0 {
			errTxt := ""
			for _, diag := range diags {
				errTxt += Sprintf("- %s\n", diag)
			}

			return cfg, responses, Errorf(errTxt)
		}

		return cfg, responses, nil
	}

	return cfg, responses, nil
}

func getRequestRepeatCount(name string, evCtx EvalContext) (int, error) {
	requestCfg, err := findRequest(name, evCtx.RawRequests)
	if err != nil {
		return 0, err
	}

	return requestCfg.Repeat, nil
}

func parseRequest(name string, evCtx EvalContext, execCtx *ExecutionContext) (*Request, error) {
	requestCfg, err := findRequest(name, evCtx.RawRequests)
	if err != nil {
		return nil, err
	}

	var responses []*Response

	// This feature hasn't been documented yet.
	if requestCfg.DependsOn != nil {
		for _, v := range requestCfg.DependsOn {
			findString := requestDependencyRegex.FindStringSubmatch(v)

			if len(findString) > 1 {
				if _, ok := evCtx.RequestAsVars[findString[1]]; !ok {
					evCtxP, response, depErr := processDependency(findString[1], &evCtx, execCtx)
					responses = append(responses, response)

					if depErr != nil {
						return nil, depErr
					}

					evCtx = *evCtxP
				}
			}
		}
	}

	ctxEvalContext := getCtxEvalContext(evCtx)
	spec := getRequestObjSpec()
	cfg, diags := hcldec.Decode(requestCfg.Remain, spec, &ctxEvalContext)

	cfg, responses, err = retryWithDependency(requestCfg, cfg, diags, evCtx, execCtx, responses)
	if err != nil {
		return nil, err
	}

	finalRequest, err, requestDiags := getRequest(cfg, *requestCfg, evCtx, execCtx)
	if err != nil {
		return nil, err
	}

	if requestDiags != nil {
		cfg, responses, err = retryWithDependency(requestCfg, cfg, requestDiags, evCtx, execCtx, responses)

		finalRequest, err, requestDiags = getRequest(cfg, *requestCfg, evCtx, execCtx)
		if err != nil {
			return nil, err
		}

		if len(requestDiags) > 0 {
			errTxt := ""
			for _, diag := range requestDiags {
				errTxt += Sprintf("- %s\n", diag)
			}

			return nil, Errorf(errTxt)
		}
	}

	finalRequest.PrecedingRequests = responses

	return finalRequest, nil
}
