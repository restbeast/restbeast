package lib

import (
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"os"
)

type Environments map[string]cty.Value

func parseEnv(env string, rawEnvironments []*EnvironmentCfg) (cty.Value, error) {
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

			secrets := map[string]map[string]string{}
			for _, secret := range rawEnvironments[i].Secrets {
				secrets[secret.Name] = map[string]string{}

				switch secret.Type {
				case "env-var":
					secrets[secret.Name] = secretEngineEnvVar(secret.Paths)
				}
			}

			value, err := gocty.ToCtyValue(secrets, cty.Map(cty.Map(cty.String)))

			if err != nil {
				fmt.Printf("Error: failed to load secrets, %s\n", err)
				os.Exit(1)
			}

			evalContext := &hcl.EvalContext{
				Variables: map[string]cty.Value{
					"secret": value,
				},
			}

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

	return cty.Value{}, nil
}
