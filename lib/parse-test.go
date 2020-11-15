package lib

import (
	. "fmt"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

// Find test by .Name property
func findTest(name string, rawTests TestCfgs) (*TestCfg, error) {
	for _, t := range rawTests {
		if name == t.Name {
			return t, nil
		}
	}

	return nil, Errorf("`%s` not found", name)
}

func parseTest(name string, evCtx EvalContext, execCtx *ExecutionContext) (*Test, error) {
	testCfg, err := findTest(name, evCtx.RawTests)

	if err != nil {
		return nil, err
	}

	fns := *evCtx.Functions
	for k, v := range assertionFunctionList {
		fns[k] = function.New(&function.Spec{
			Params: v.Params,
			Type:   v.Type,
			Impl:   v.Impl,
		})
	}
	evCtx.Functions = &fns

	ctxEvalContext := getCtxEvalContext(evCtx)
	var assertions map[string]cty.Value
	diags := gohcl.DecodeBody(testCfg.Body, &ctxEvalContext, &assertions)
	if len(diags) > 0 {
		errTxt := ""
		for _, diag := range diags {
			errTxt += Sprintf("- %s\n", diag)
		}

		return nil, Errorf(errTxt)
	}

	var results []Assertion
	for k, v := range assertions {
		if v.AsString() == "PASS" {
			results = append(results, Assertion{
				Name: k,
				Pass: true,
			})
		} else {
			results = append(results, Assertion{
				Name: k,
				Pass: false,
				Text: v.AsString(),
			})
		}
	}

	return &Test{
		Name:       "name",
		Assertions: results,
	}, nil
}
