package lib

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"net/http"
	"time"
)

type VariableCfg struct {
	Name  string   `hcl:"name,label"`
	Value hcl.Body `hcl:"value,remain"`
}

type VariableCfgs []*VariableCfg

type BasicAuthCfg struct {
	Body hcl.Body `hcl:",remain"`
}

type AuthCfg struct {
	BasicAuth *BasicAuthCfg `hcl:"basic,block"`
}

type RequestCfg struct {
	Name      string   `hcl:"name,label"`
	DependsOn []string `hcl:"depends_on,optional"`
	Body      hcl.Body `hcl:",remain"`
	Auth      *AuthCfg `hcl:"auth,block"`
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
	Dynamics				  VariableCfgs				 `hcl:"dynamic,block"`
	Environments      EnvironmentCfgs      `hcl:"env,block"`
	ExternalFunctions ExternalFunctionCfgs `hcl:"external-function,block"`
	Version           string               `hcl:"version,optional"`
}

type RequestTiming struct {
	// DNS resolve time
	Dns time.Duration
	// Time to establish connection
	Conn time.Duration
	// TLS handshake time
	Tls time.Duration
	// Time to first byte
	FirstByte time.Duration
	// Total request duration
	Total time.Duration
}

type Response struct {
	Method     string
	Url        string
	StatusCode int
	Proto      string
	Body       []byte
	Headers    http.Header
	Timing     RequestTiming
	Request    *Request
}

type Request struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
	EvalContext
	PrecedingRequests []*Response
	*Response
	*ExecutionContext
	RoundTripper http.RoundTripper
}

type EvalContext struct {
	Functions     *map[string]function.Function
	Variables     *map[string]cty.Value
	Environment   *cty.Value
	RequestAsVars RequestAsVars
	RawRequests   RequestCfgs
	RawDynamics		VariableCfgs
}

type RequestAsVars map[string]cty.Value
type Requests map[string]map[string]cty.Value

type ExecutionContext struct {
	Version string
	Debug   bool
}
