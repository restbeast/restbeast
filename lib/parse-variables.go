package lib

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"os"
)

func parseVariables(rawVars []*VariableCfg, functions map[string]function.Function) map[string]cty.Value {
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
			for _, diag := range diags {
				fmt.Printf("- %s\n", diag)
			}

			os.Exit(1)
		}

		variables[varCfg.Name] = cfg
	}

	return variables
}
