package lib

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type VariableCfg struct {
	Name string `hcl:"name,label"`
	Default hcl.Attributes `hcl:"default,remain"`
}

type RequestCfg struct {
	Name string `hcl:"name,label"`
	Body hcl.Body `hcl:",remain"`
}

type RequestCfgs []*RequestCfg

type EnvironmentCfg struct {
	Name      string `hcl:"name,label"`
	Default   bool `hcl:"default,optional"`
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
