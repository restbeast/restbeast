package lib

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

func parseVariables(rawVars []*VariableCfg, functions map[string]function.Function) (*map[string]cty.Value, error) {
	variables := map[string]cty.Value{}

	evalContext := &hcl.EvalContext{
		Functions: functions,
	}

	spec := &hcldec.AttrSpec{
		Name:     "value",
		Type:     cty.DynamicPseudoType,
		Required: false,
	}

	for _, varCfg := range rawVars {
		cfg, diags := hcldec.Decode(varCfg.Value, spec, evalContext)
		if len(diags) != 0 {
			var errTxt string
			for _, diag := range diags {
				errTxt += fmt.Sprintf("- %s\n", diag)
			}

			return nil, fmt.Errorf(errTxt)
		}

		variables[varCfg.Name] = cfg
	}

	return &variables, nil
}
