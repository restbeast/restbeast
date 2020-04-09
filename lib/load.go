package lib

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/gohcl"
)

func LoadOne(name, env, version string) (request Request, err error) {
	mergedBody, err := readFiles()
	if err != nil {
		return Request{}, err
	}

	var root RootCfg
	// Decode merged HCL body into root config struct
	diags := gohcl.DecodeBody(mergedBody, nil, &root)
	if len(diags) != 0 {
		for _, diag := range diags {
			fmt.Printf("decoding - %s\n", diag)
		}
		return
	}

	envVars, _ := parseEnv(env, root.Environments)
	variables := parseVariables(root.Variables)
	request = parseRequest(name, variables, envVars, version, root.Requests)

	return request, nil
}
