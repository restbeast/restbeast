package lib

import (
	"encoding/json"
	"errors"
	. "fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
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

func getEvalContext(variables map[string]cty.Value, envVars cty.Value, requestAsVars map[string]cty.Value, functions map[string]function.Function) hcl.EvalContext {
	return hcl.EvalContext{
		Variables: map[string]cty.Value{
			"var":     cty.ObjectVal(variables),
			"env":     envVars,
			"request": cty.ObjectVal(requestAsVars),
		},
		Functions: functions,
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

func processDependencies(dependencies []string, variables map[string]cty.Value, envVars cty.Value, execCtx *ExecutionContext, functions map[string]function.Function, rawRequests RequestCfgs) (requestAsVars map[string]cty.Value) {
	requestAsVars = make(map[string]cty.Value)

	for _, dep := range getUniqueDependencies(dependencies) {
		request := parseRequest(dep, variables, envVars, execCtx, functions, rawRequests)
		response := DoRequest(request, execCtx)

		var decoded interface{}
		err := json.Unmarshal(response.Body, &decoded)

		if err != nil {
			Println("Error: error decoding json response body", err)
			os.Exit(1)
		}

		requestAsVars[dep] = walkThrough(reflect.ValueOf(decoded))
	}

	return requestAsVars
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

func parseRequest(name string, variables map[string]cty.Value, envVars cty.Value, execCtx *ExecutionContext, functions map[string]function.Function, rawRequests RequestCfgs) Request {
	err, request := findRequest(name, rawRequests)

	if err != nil {
		Printf("Error: %s\n", err)
		os.Exit(1)
	}

	requestAsVars := map[string]cty.Value{}
	evalContext := getEvalContext(variables, envVars, requestAsVars, functions)
	spec := getObjSpec()

	cfg, diags := hcldec.Decode(request.Body, spec, &evalContext)
	dependencies, restDiagMsgs := getPossibleDependencies(diags)

	if len(restDiagMsgs) > 0 {
		for _, diag := range restDiagMsgs {
			Printf("- %s\n", diag)
		}

		os.Exit(1)
	}

	if len(dependencies) > 0 {
		requestAsVars := processDependencies(dependencies, variables, envVars, execCtx, functions, rawRequests)
		evalContext = getEvalContext(variables, envVars, requestAsVars, functions)

		cfg, diags = hcldec.Decode(request.Body, spec, &evalContext)

		if len(diags) > 0 {
			for _, diag := range diags {
				Printf("- %s\n", diag)
			}

			os.Exit(1)
		}
	}

	return getRequest(cfg)
}
