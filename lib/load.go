package lib

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/gohcl"
)

func LoadOne(name, env string, execCtx *ExecutionContext) (request Request, err error) {
	mergedBody, err := readFiles()
	if err != nil {
		return Request{}, err
	}

	internalFunctions := getCtyFunctions()

	var root RootCfg
	// Decode merged HCL body into root config struct
	diags := gohcl.DecodeBody(mergedBody, nil, &root)
	if len(diags) != 0 {
		for _, diag := range diags {
			fmt.Printf("decoding - %s\n", diag)
		}
		return
	}

	functions := parseExternalFunctions(internalFunctions, root.ExternalFunctions, execCtx)

	envVars, envErr := parseEnv(env, root.Environments)
	if envErr != nil {
		return Request{}, envErr
	}

	variables := parseVariables(root.Variables, functions)
	request = parseRequest(name, variables, envVars, execCtx, functions, root.Requests)

	return request, nil
}
