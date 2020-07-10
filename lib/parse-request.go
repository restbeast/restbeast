package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
	"github.com/zclconf/go-cty/cty/gocty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"os"
	"reflect"
	"regexp"
)

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

func walkThroughJson(v reflect.Value) cty.Value {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		length := v.Len()

		newSlice := make([]cty.Value, length)

		for i := 0; i < length; i++ {
			newSlice[i] = walkThroughJson(v.Index(i))
		}

		return cty.ListVal(newSlice)
	case reflect.Map:
		newMap := make(map[string]cty.Value)

		for _, k := range v.MapKeys() {
			newMap[k.String()] = walkThroughJson(v.MapIndex(k))
		}

		return cty.ObjectVal(newMap)
	case reflect.Bool:
		return cty.BoolVal(v.Bool())
	default:
		return cty.StringVal(v.String())
	}
}

func getEvalContext(variables map[string]cty.Value, envVars cty.Value, requestAsVars map[string]cty.Value) hcl.EvalContext {
	return hcl.EvalContext{
		Variables: map[string]cty.Value{
			"var": cty.ObjectVal(variables),
			"env": envVars,
			"request": cty.ObjectVal(requestAsVars),
		},
		Functions: map[string]function.Function{
			"upper": stdlib.UpperFunc,
			"uuid":  funcGenerateUuid(),
		},
	}
}

func findRequest(name string, rawRequests []*RequestCfg) (err error, request RequestCfg) {
	for _, r := range rawRequests {
		if name == r.Name {
			return nil, *r
		}
	}

	return errors.New("request not found"), RequestCfg{}
}

func getPossibleDependencies(diags hcl.Diagnostics) (dependencies []string) {
	if len(diags) != 0 {
		diagMessageRegex := regexp.MustCompile(`This object does not have an attribute named "(?P<name>[\w\d-_]+)"`)

		for _, diag := range diags {
			if diag.Summary == "Unsupported attribute" {
				findString := diagMessageRegex.FindStringSubmatch(diag.Detail)

				if len(findString) > 1 {
					dependencies = append(dependencies, findString[1])
				}
			}
		}

		if len(dependencies) == 0 {
			for _, diag := range diags {
				fmt.Printf("- %s\n", diag)
			}

			os.Exit(0)
		}
	}

	return dependencies
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

func processDependencies(dependencies []string, variables map[string]cty.Value, envVars cty.Value, version string, rawRequests RequestCfgs) (requestAsVars map[string]cty.Value) {
	requestAsVars = make(map[string]cty.Value)

	for _, dep := range getUniqueDependencies(dependencies) {
		request := parseRequest(dep, variables, envVars, version, rawRequests)
		response := DoRequest(request, version)

		var decoded interface{}
		err := json.Unmarshal(response.Body, &decoded)

		if err != nil {
			fmt.Println("Error: error decoding json response body", err)
			os.Exit(0)
		}

		requestAsVars[dep] = walkThroughJson(reflect.ValueOf(decoded))
	}

	return requestAsVars
}

func getRequest(cfg cty.Value) Request {
	body, _ := json.MarshalIndent(ctyjson.SimpleJSONValue{cfg.GetAttr("body")}, "", "  ")
	var headers map[string]string
	_ = gocty.FromCtyValue(cfg.GetAttr("headers"), &headers)

	request := Request{
		Method:  cfg.GetAttr("method").AsString(),
		Url:     cfg.GetAttr("url").AsString(),
		Headers: headers,
		Body:    string(body),
	}

	return request
}

func parseRequest(name string, variables map[string]cty.Value, envVars cty.Value, version string, rawRequests RequestCfgs) Request {
  err, request := findRequest(name, rawRequests)

  if err != nil {
  	fmt.Println("Error: Request not found")
  	os.Exit(1)
	}

	requestAsVars := map[string]cty.Value{}
	evalContext := getEvalContext(variables, envVars, requestAsVars)
	spec := getObjSpec()

	cfg, diags := hcldec.Decode(request.Body, spec, &evalContext)
	dependencies := getPossibleDependencies(diags)

	if len(dependencies) > 0 {
		requestAsVars := processDependencies(dependencies, variables, envVars, version, rawRequests)
		evalContext = getEvalContext(variables, envVars, requestAsVars)

		cfg, diags = hcldec.Decode(request.Body, spec, &evalContext)

		if len(diags) > 0 {
			for _, diag := range diags {
				fmt.Printf("- %s\n", diag)
			}

			os.Exit(0)
		}
	}

	return getRequest(cfg)
}
