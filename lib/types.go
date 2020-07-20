package lib

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type VariableCfg struct {
	Name string `hcl:"name,label"`
	Value string `hcl:"value,attr"`
}

type RequestCfg struct {
	Name string `hcl:"name,label"`
	Body hcl.Body `hcl:",remain"`
}

type RequestCfgs []*RequestCfg

type SecretCfg struct {
	Name string `hcl:"name,label"`
	Type string `hcl:"type,attr"`
	Paths map[string]string `hcl:"paths,attr"`
}

type EnvironmentCfg struct {
	Name      string `hcl:"name,label"`
	Default   bool `hcl:"default,optional"`
	Secrets   []*SecretCfg `hcl:"secrets,block"`
	Variables hcl.Body `hcl:",remain"`
}

type RootCfg struct {
	Requests     []*RequestCfg     `hcl:"request,block"`
	Variables    []*VariableCfg    `hcl:"variable,block"`
	Environments []*EnvironmentCfg `hcl:"env,block"`
}

type Request struct {
	Method string
	Url string
	Headers map[string]string
	Body string
}

type Requests map[string]map[string]cty.Value
