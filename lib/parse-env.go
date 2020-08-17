package lib

import (
	"errors"
	. "fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type ParsedSecret map[string]string
type ParsedSecretGroup map[string]ParsedSecret
type AllParsedSecrets map[string]ParsedSecretGroup

func parseSecret(secretCfg *SecretCfg) ParsedSecretGroup {
	cfg := *secretCfg
	parsedSecret := make(ParsedSecretGroup)

	switch cfg.Type {
	case "env-var":
		parsedSecret[cfg.Name] = secretEngineEnvVar(cfg.Paths)
	}

	return parsedSecret
}

func parseSecrets(secretCfgs SecretCfgs) AllParsedSecrets {
	parsedSecrets := make(AllParsedSecrets)

	for _, secret := range secretCfgs {
		parseSecret(secret)
	}

	return parsedSecrets
}

func parseEnv(env string, rawEnvironments EnvironmentCfgs) (*cty.Value, error) {
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
				for _, diag := range diags {
					Printf("- %s\n", diag)
				}
				return nil, errors.New("environment definition contains errors")
			}

			vars := cfg.GetAttr("variables")
			return &vars, nil
		}
	}

	return nil, nil
}
