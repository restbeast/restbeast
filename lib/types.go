package lib

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

type VariableCfg struct {
	Name  string   `hcl:"name,label"`
	Value hcl.Body `hcl:"value,remain"`
}

type RequestCfg struct {
	Name string   `hcl:"name,label"`
	Body hcl.Body `hcl:",remain"`
}

type RequestCfgs []*RequestCfg

type SecretCfg struct {
	Name  string            `hcl:"name,label"`
	Type  string            `hcl:"type,attr"`
	Paths map[string]string `hcl:"paths,attr"`
}

type EnvironmentCfg struct {
	Name      string       `hcl:"name,label"`
	Default   bool         `hcl:"default,optional"`
	Secrets   []*SecretCfg `hcl:"secrets,block"`
	Variables hcl.Body     `hcl:",remain"`
}

type ExternalFunctionCfg struct {
	Name        string   `hcl:"name,label"`
	Interpreter string   `hcl:"interpreter,attr"`
	Script      string   `hcl:"script,attr"`
	Args        []string `hcl:"args,optional"`
}

type RootCfg struct {
	Requests          []*RequestCfg          `hcl:"request,block"`
	Variables         []*VariableCfg         `hcl:"variable,block"`
	Environments      []*EnvironmentCfg      `hcl:"env,block"`
	ExternalFunctions []*ExternalFunctionCfg `hcl:"external-function,block"`
}

type Request struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
	EvalContext
}

type EvalContext struct {
	Functions   map[string]function.Function
	Variables   map[string]cty.Value
	Environment cty.Value
	RawRequests RequestCfgs
}

type Requests map[string]map[string]cty.Value

type ExecutionContext struct {
	Version string
	Debug   bool
}
