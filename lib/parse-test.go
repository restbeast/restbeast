package lib

import (
	. "fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"sort"
)

type ParsedAssertions map[string]cty.Value

// Find test by .Name property
func findTest(name string, rawTests TestCfgs) (*TestCfg, error) {
	for _, t := range rawTests {
		if name == t.Name {
			return t, nil
		}
	}

	return nil, Errorf("`%s` not found", name)
}

func updateEvalContextWithTestFns(evCtx *EvalContext) {
	fns := *evCtx.Functions
	for k, v := range assertionFunctionList {
		fns[k] = function.New(&function.Spec{
			Params: v.Params,
			Type:   v.Type,
			Impl:   v.Impl,
		})
	}
	evCtx.Functions = &fns
}

func prepareResults(name string, assertions map[string]cty.Value) *Test {
	var passResults Assertions
	var failResults Assertions
	for k, v := range assertions {
		if v.AsString() == "PASS" {
			passResults = append(passResults, Assertion{
				Name: k,
				Pass: true,
			})
		} else {
			failResults = append(failResults, Assertion{
				Name: k,
				Pass: false,
				Text: v.AsString(),
			})
		}
	}

	sort.Slice(passResults, func(i, j int) bool {
		return passResults[i].Name < passResults[j].Name
	})
	sort.Slice(failResults, func(i, j int) bool {
		return failResults[i].Name < failResults[j].Name
	})

	return &Test{
		Name:       name,
		Assertions: append(passResults, failResults...),
	}
}

func retryTestWithDependency(testCfg *TestCfg, diags hcl.Diagnostics, evCtx EvalContext, execCtx *ExecutionContext, responses []*Response, parsedAssertions *ParsedAssertions) error {
	dependencies, restDiagMsgs := getPossibleDependencies(diags)
	if len(restDiagMsgs) > 0 {
		errTxt := ""
		for _, diag := range restDiagMsgs {
			errTxt += Sprintf("- %s\n", diag)
		}

		return Errorf(errTxt)
	}

	if len(dependencies) > 0 {
		uniqueDeps := getUniqueDependencies(dependencies)
		sortedDeps, err := sortCrossDependency(uniqueDeps, evCtx)

		if err != nil {
			return err
		}

		for _, dependency := range sortedDeps {
			if _, ok := evCtx.RequestAsVars[dependency]; !ok {
				evCtxP, response, err := processDependency(dependency, &evCtx, execCtx)
				if err != nil {
					return err
				}

				responses = append(responses, response)
				evCtx = *evCtxP
			}
		}

		ctxEvalContext := getCtxEvalContext(evCtx)
		diags := gohcl.DecodeBody(testCfg.Body, &ctxEvalContext, parsedAssertions)

		if len(diags) > 0 {
			errTxt := ""
			for _, diag := range diags {
				errTxt += Sprintf("- %s\n", diag)
			}

			return Errorf(errTxt)
		}
	}

	return nil
}

func parseTest(name string, evCtx EvalContext, execCtx *ExecutionContext) (*Test, error) {
	testCfg, err := findTest(name, evCtx.RawTests)
	if err != nil {
		return nil, err
	}

	updateEvalContextWithTestFns(&evCtx)

	var responses []*Response
	ctxHclEvalContext := getCtxEvalContext(evCtx)
	var parsedAssertions ParsedAssertions
	diags := gohcl.DecodeBody(testCfg.Body, &ctxHclEvalContext, &parsedAssertions)
	depErr := retryTestWithDependency(testCfg, diags, evCtx, execCtx, responses, &parsedAssertions)
	if depErr != nil {
		return nil, depErr
	}

	return prepareResults(name, parsedAssertions), nil
}
