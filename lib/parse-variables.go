package lib

import (
	"github.com/zclconf/go-cty/cty"
)

func parseVariables(parsedVars []*VariableCfg) map[string]cty.Value {
	variables := map[string]cty.Value{}
	for _, v := range parsedVars {
		if len(v.Value) == 0 {
			continue
		}

		variables[v.Name] = cty.StringVal(v.Value)
	}

	return variables
}
