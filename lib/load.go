package lib

import (
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2/gohcl"
)

// Reads all files in present folder
// Merge all files into a single body
// First hcl decoding pass
// Return decoded body
func readAndDecodeBody() (*RootCfg, error) {
	mergedBody, err := readFiles()
	if err != nil {
		return nil, err
	}

	var root RootCfg
	// Decode merged HCL body into root config struct
	diags := gohcl.DecodeBody(mergedBody, nil, &root)

	if len(diags) != 0 {
		errTxt := ""

		for _, diag := range diags {
			errTxt += fmt.Sprintf("decoding - %s\n", diag)
		}

		return nil, errors.New(errTxt)
	}

	return &root, nil
}

// Get all internal functions
// Parse external functions
// Parse environment variables
// Parse variables
// Create EvalContext
func LoadEvalCtx(env string, execCtx *ExecutionContext) (*EvalContext, error) {
	root, err := readAndDecodeBody()
	if err != nil {
		return nil, err
	}

	internalFunctions := getCtyFunctions()
	functions := parseExternalFunctions(internalFunctions, root.ExternalFunctions, execCtx)

	envVars, envErr := parseEnv(env, root.Environments)
	if envErr != nil {
		return nil, envErr
	}

	variables := parseVariables(root.Variables, functions)

	return &EvalContext{
		Functions:     functions,
		Variables:     variables,
		Environment:   *envVars,
		RawRequests:   root.Requests,
		RequestAsVars: RequestAsVars{},
	}, nil
}

// Load only request with given EvalContext
func LoadOnlyRequest(name string, evCtx *EvalContext, execCtx *ExecutionContext) (request *Request, err error) {

	return parseRequest(name, *evCtx, *execCtx)
}

// Gather EvalContext and load given request
func LoadWhole(name, env string, execCtx *ExecutionContext) (request *Request, err error) {
	evCtx, err := LoadEvalCtx(env, execCtx)
	if err != nil {
		return nil, err
	}

	return parseRequest(name, *evCtx, *execCtx)
}
