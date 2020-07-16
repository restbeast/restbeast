package lib

import (
	"fmt"
	"github.com/zclconf/go-cty/cty"
)

func parseVariables(parsedVars []*VariableCfg) map[string]cty.Value {
	variables := map[string]cty.Value{}
	for _, v := range parsedVars {
		if len(v.Value) == 0 {
			continue
		}

		val, diags := v.Value["default"].Expr.Value(nil)
		if len(diags) != 0 {
			for _, diag := range diags {
				fmt.Printf("decoding - %s\n", diag)
			}
			return nil
		}

		variables[v.Name] = val
	}

	return variables
}
