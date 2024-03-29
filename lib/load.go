package lib

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/zclconf/go-cty/cty"
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

// LoadEvalCtx Get all internal functions
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
		RawDynamics:   root.Dynamics,
		RawRequests:   root.Requests,
		RawTests:      root.Tests,
		RequestAsVars: &RequestAsVars{},
	}, nil
}

func loadDynamics(evCtx *EvalContext) error {
	dynamics, err := parseVariables(evCtx.RawDynamics, *evCtx.Functions)
	if err != nil {
		return err
	}

	variables := map[string]cty.Value{}
	if evCtx.Variables != nil {
		variables = *evCtx.Variables
	}

	if dynamics != nil {
		for varName, value := range *dynamics {
			variables[varName] = value
		}
	}
	evCtx.Variables = &variables

	return nil
}

// LoadOnlyRequest Load only request with given EvalContext
func LoadOnlyRequest(name string, evCtx *EvalContext, execCtx *ExecutionContext) (request *Request, err error) {
	err = loadDynamics(evCtx)
	if err != nil {
		return nil, err
	}

	return parseRequest(name, evCtx, execCtx)
}

// LoadWhole Gather EvalContext and load given request
func LoadWhole(name, env string, execCtx *ExecutionContext) (request *Request, err error) {
	evCtx, err := LoadEvalCtx(env, execCtx)
	if err != nil {
		return nil, err
	}

	err = loadDynamics(evCtx)
	if err != nil {
		return nil, err
	}

	return parseRequest(name, evCtx, execCtx)
}

func ListRequestsAndTests(lsType string, execCtx *ExecutionContext) (list []ListObject, maxNameLen int, err error) {
	maxNameLen = 0
	root, err := readAndDecodeBody()
	if err != nil {
		return list, maxNameLen, err
	}

	if root.Version != "" {
		err = compareVersion(root.Version, execCtx.Version)

		if err != nil {
			return list, maxNameLen, err
		}
	}

	if lsType == "" || lsType == "request" {
		for _, requestCfg := range root.Requests {
			thisLen := len([]rune(requestCfg.Name))
			if thisLen > maxNameLen {
				maxNameLen = thisLen
			}
			list = append(list, ListObject{requestCfg.Name, "request"})
		}
	}

	if lsType == "" || lsType == "test" {
		for _, testCfg := range root.Tests {
			thisLen := len([]rune(testCfg.Name))
			if thisLen > maxNameLen {
				maxNameLen = thisLen
			}
			list = append(list, ListObject{testCfg.Name, "test"})
		}
	}

	sort.Slice(
		list, func(i, j int) bool {
			return list[i].Name < list[j].Name
		},
	)

	return list, maxNameLen, nil
}

func LoadTest(name, env string, execCtx *ExecutionContext) (request *Test, err error) {
	evCtx, err := LoadEvalCtx(env, execCtx)
	if err != nil {
		return nil, err
	}

	err = loadDynamics(evCtx)
	if err != nil {
		return nil, err
	}

	return parseTest(name, evCtx, execCtx)
}

func LoadAllTests(env string, execCtx *ExecutionContext) (tests Tests, err error) {
	evCtx, err := LoadEvalCtx(env, execCtx)
	if err != nil {
		return nil, err
	}

	err = loadDynamics(evCtx)
	if err != nil {
		return nil, err
	}

	for _, t := range evCtx.RawTests {
		uniqueCtx := &EvalContext{
			Functions:     evCtx.Functions,
			Variables:     evCtx.Variables,
			Environment:   evCtx.Environment,
			RequestAsVars: &RequestAsVars{},
			RawRequests:   evCtx.RawRequests,
			RawDynamics:   evCtx.RawDynamics,
			RawTests:      evCtx.RawTests,
		}
		test, err := parseTest(t.Name, uniqueCtx, execCtx)
		if err != nil {
			return nil, err
		}

		tests = append(tests, test)
	}

	return tests, err
}

func LoadRepeatCount(name string, evCtx *EvalContext) (int, error) {
	return getRequestRepeatCount(name, evCtx)
}
