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
	"os"
	"reflect"
	"regexp"
)

var dependencyDiagMessageRegex = regexp.MustCompile(`This object does not have an attribute named "(?P<name>[\w\d-_]+)"`)

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

func getCtxEvalContext(requestAsVars map[string]cty.Value, evCtx EvalContext) hcl.EvalContext {
	return hcl.EvalContext{
		Variables: map[string]cty.Value{
			"var":     cty.ObjectVal(evCtx.Variables),
			"env":     evCtx.Environment,
			"request": cty.ObjectVal(requestAsVars),
		},
		Functions: evCtx.Functions,
	}
}

func findRequest(name string, rawRequests []*RequestCfg) (err error, request RequestCfg) {
	for _, r := range rawRequests {
		if name == r.Name {
			return nil, *r
		}
	}

	return errors.New(Sprintf("`%s` not found", name)), RequestCfg{}
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

func processDependencies(dependencies []string, evCtx EvalContext, execCtx ExecutionContext) (requestAsVars map[string]cty.Value, err error) {
	requestAsVars = make(map[string]cty.Value)

	for _, dep := range getUniqueDependencies(dependencies) {
		request, parseErr := parseRequest(dep, evCtx, execCtx)
		if parseErr != nil {
			return nil, parseErr
		}

		response := DoRequest(*request, &execCtx)

		var decoded interface{}
		err := json.Unmarshal(response.Body, &decoded)

		if err != nil {
			return nil, errors.New(Sprintf("error decoding json response body, %s\n", err))
		}

		requestAsVars[dep] = walkThrough(reflect.ValueOf(decoded))
	}

	return requestAsVars, nil
}

func getRequest(cfg cty.Value) Request {
	body, jsonErr := json.MarshalIndent(ctyjson.SimpleJSONValue{cfg.GetAttr("body")}, "", "  ")
	if jsonErr != nil {
		Printf("Error: failed to parse body, %s\n", jsonErr)
		os.Exit(1)
	}

	var headers map[string]string
	headerErr := gocty.FromCtyValue(cfg.GetAttr("headers"), &headers)

	if headerErr != nil {
		Printf("Error: failed to parse headers, %s\n", headerErr)
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

	requestAsVars := map[string]cty.Value{}
	evalContext := getCtxEvalContext(requestAsVars, evCtx)
	spec := getObjSpec()

	cfg, diags := hcldec.Decode(request.Body, spec, &evalContext)
	dependencies, restDiagMsgs := getPossibleDependencies(diags)

	if len(restDiagMsgs) > 0 {
		errTxt := ""
		for _, diag := range restDiagMsgs {
			errTxt += Sprintf("- %s\n", diag)
		}

		return nil, errors.New(errTxt)
	}

	if len(dependencies) > 0 {
		requestAsVars, depErr := processDependencies(dependencies, evCtx, execCtx)
		if depErr != nil {
			return nil, err
		}

		evalContext = getCtxEvalContext(requestAsVars, evCtx)

		cfg, diags = hcldec.Decode(request.Body, spec, &evalContext)

		if len(diags) > 0 {
			errTxt := ""
			for _, diag := range diags {
				errTxt += Sprintf("- %s\n", diag)
			}

			return nil, errors.New(errTxt)
		}
	}

	finalRequest := getRequest(cfg)

	return &finalRequest, nil
}
