package lib

import (
	. "fmt"
	"log"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type ParsedSecret map[string]string
type AllParsedSecrets map[string]ParsedSecret

func parseSecrets(secretCfgs SecretCfgs) AllParsedSecrets {
	parsedSecrets := make(AllParsedSecrets)

	for _, secretCfg := range secretCfgs {
		switch secretCfg.Type {
		case "env-var":
			parsedSecrets[secretCfg.Name] = secretEngineEnvVar(secretCfg.Paths)
		}
	}

	return parsedSecrets
}

func parseEnv(env string, rawEnvironments EnvironmentCfgs, execCtx *ExecutionContext) (*cty.Value, error) {
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
			if execCtx.Debug {
				log.Printf("parser: loading env '%s'", rawEnvironments[i].Name)
			}

			secrets := parseSecrets(rawEnvironments[i].Secrets)
			value, err := gocty.ToCtyValue(secrets, cty.Map(cty.Map(cty.String)))
			if err != nil {
				return nil, Errorf("failed to load secrets, %s\n", err)
			}

			evalContext := &hcl.EvalContext{
				Variables: map[string]cty.Value{
					"secret": value,
				},
			}

			cfg, diags := hcldec.Decode(rawEnvironments[i].Variables, spec, evalContext)
			if len(diags) != 0 {
				errTxt := ""
				for _, diag := range diags {
					errTxt += Sprintf("- %s\n", diag)
				}

				return nil, Errorf(errTxt)
			}

			vars := cfg.GetAttr("variables")
			return &vars, nil
		}
	}

	return nil, nil
}
