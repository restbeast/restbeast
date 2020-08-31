package lib

import (
	"errors"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/hashicorp/hcl/v2/gohcl"
	"strings"
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

func compareVersion(constraint, actual string) error {
	cnst, err := semver.NewConstraint(constraint)
	if err != nil {
		return fmt.Errorf("Invalid version,\n got: %s\n", constraint)
	}

	if actual == "" {
		return nil
	}

	actl, _ := semver.NewVersion(strings.TrimLeft(actual, "v"))

	if !cnst.Check(actl) {
		return fmt.Errorf("Invalid restbeast version,\n expected: %s\n got: %s\n", constraint, actual)
	}

	return nil
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

	if root.Version != "" {
		err := compareVersion(root.Version, execCtx.Version)

		if err != nil {
			return nil, err
		}
	}

	internalFunctions := getCtyFunctions()
	functions, fnErr := parseExternalFunctions(internalFunctions, root.ExternalFunctions, execCtx)
	if fnErr != nil {
		return nil, fnErr
	}

	envVars, envErr := parseEnv(env, root.Environments, execCtx)
	if envErr != nil {
		return nil, envErr
	}

	variables, varErr := parseVariables(root.Variables, *functions)
	if varErr != nil {
		return nil, varErr
	}

	return &EvalContext{
		Functions:     functions,
		Variables:     variables,
		Environment:   envVars,
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
