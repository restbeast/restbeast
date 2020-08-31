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

type VariableCfgs []*VariableCfg

type RequestCfg struct {
	Name      string   `hcl:"name,label"`
	DependsOn []string `hcl:"depends_on,optional"`
	Body      hcl.Body `hcl:",remain"`
}

type RequestCfgs []*RequestCfg

type SecretCfg struct {
	Name  string            `hcl:"name,label"`
	Type  string            `hcl:"type,attr"`
	Paths map[string]string `hcl:"paths,attr"`
}

type SecretCfgs []*SecretCfg

type EnvironmentCfg struct {
	Name      string     `hcl:"name,label"`
	Default   bool       `hcl:"default,optional"`
	Secrets   SecretCfgs `hcl:"secrets,block"`
	Variables hcl.Body   `hcl:",remain"`
}

type EnvironmentCfgs []*EnvironmentCfg

type ExternalFunctionCfg struct {
	Name        string   `hcl:"name,label"`
	Interpreter string   `hcl:"interpreter,attr"`
	Script      string   `hcl:"script,attr"`
	Args        []string `hcl:"args,optional"`
}

type ExternalFunctionCfgs []*ExternalFunctionCfg

type RootCfg struct {
	Requests          RequestCfgs          `hcl:"request,block"`
	Variables         VariableCfgs         `hcl:"variable,block"`
	Environments      EnvironmentCfgs      `hcl:"env,block"`
	ExternalFunctions ExternalFunctionCfgs `hcl:"external-function,block"`
	Version           string               `hcl:"version,optional"`
}

type Request struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
	EvalContext
	PrecedingRequests []*Response
}

type EvalContext struct {
	Functions     *map[string]function.Function
	Variables     *map[string]cty.Value
	Environment   *cty.Value
	RequestAsVars RequestAsVars
	RawRequests   RequestCfgs
}

type RequestAsVars map[string]cty.Value
type Requests map[string]map[string]cty.Value

type ExecutionContext struct {
	Version string
	Debug   bool
}
