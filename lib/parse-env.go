package lib

import (
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type Environments map[string]cty.Value

func parseEnv(env string, rawEnvironments []*EnvironmentCfg) (cty.Value, error) {
	evalContext := &hcl.EvalContext{}
	spec := &hcldec.ObjectSpec{
		"default": &hcldec.AttrSpec{
			Name:     "default",
			Type:     cty.Bool,
			Required: false,
		},
		"variables": &hcldec.AttrSpec{
			Name:     "variables",
			Type:     cty.DynamicPseudoType,
			Required: false,
		},
	}

	for i := range rawEnvironments {
		if (env != "" && rawEnvironments[i].Name == env) || (env == "" && rawEnvironments[i].Default) {
			cfg, diags := hcldec.Decode(rawEnvironments[i].Variables, spec, evalContext)
			if len(diags) != 0 {
				for _, diag := range diags {
					fmt.Printf("- %s\n", diag)
				}
				return cty.Value{}, errors.New("environment definition contains errors")
			}

			return cfg.GetAttr("variables"), nil
		}
	}

	return cty.Value{}, errors.New("environment not found")
}
