package lib

import (
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

type VariableCfg struct {
	Name  string   `hcl:"name,label"`
	Value hcl.Body `hcl:"value,remain"`
}

type VariableCfgs []*VariableCfg

type BasicAuthCfg struct {
	Body hcl.Body `hcl:",remain"`
}

type BearerAuthCfg struct {
	Body hcl.Body `hcl:",remain"`
}

type AuthCfg struct {
	BasicAuth  *BasicAuthCfg  `hcl:"basic,block"`
	BearerAuth *BearerAuthCfg `hcl:"bearer,block"`
}

type ParamsBlockCfg struct {
	Body hcl.Body `hcl:",remain"`
}

type RequestCfg struct {
	Name      string          `hcl:"name,label"`
	DependsOn []string        `hcl:"depends_on,optional"`
	Remain    hcl.Body        `hcl:",remain"`
	Auth      *AuthCfg        `hcl:"auth,block"`
	Params    *ParamsBlockCfg `hcl:"params,block"`
	Repeat    int             `hcl:"repeat,optional"`
}

type RequestCfgs []*RequestCfg

type ListObject struct {
	Name string
	Type string
}

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

type TestCfg struct {
	Name string   `hcl:"name,label"`
	Body hcl.Body `hcl:",remain"`
}

type TestCfgs []*TestCfg

type RootCfg struct {
	Requests          RequestCfgs          `hcl:"request,block"`
	Variables         VariableCfgs         `hcl:"variable,block"`
	Dynamics          VariableCfgs         `hcl:"dynamic,block"`
	Environments      EnvironmentCfgs      `hcl:"env,block"`
	ExternalFunctions ExternalFunctionCfgs `hcl:"external-function,block"`
	Tests             TestCfgs             `hcl:"test,block"`
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
	Method        string
	Url           string
	StatusCode    int
	Proto         string
	Body          []byte
	Headers       *Headers
	Timing        RequestTiming
	Request       *Request
	BytesSend     uint64
	BytesReceived uint64
}

type Request struct {
	Method  string
	Url     string
	FullUrl string
	Headers Headers
	Body    io.Reader
	Params  map[string]string
	EvalContext
	PrecedingRequests []*Response
	*Response
	*ExecutionContext
	RoundTripper http.RoundTripper
}

type Assertion struct {
	Name string
	Pass bool
	Text string
}

type Assertions []Assertion

type Test struct {
	Name       string
	Assertions Assertions
}

type Tests []*Test

type EvalContext struct {
	Functions     *map[string]function.Function
	Variables     *map[string]cty.Value
	Environment   *cty.Value
	RequestAsVars *RequestAsVars
	RawRequests   RequestCfgs
	RawDynamics   VariableCfgs
	RawTests      TestCfgs
}

type RequestAsVars struct {
	sync.Map
}

func (rv *RequestAsVars) AsCtyMap() map[string]cty.Value {
	out := make(map[string]cty.Value)
	rv.Range(
		func(key interface{}, value interface{}) bool {
			out[key.(string)] = value.(cty.Value)
			return true
		},
	)

	return out
}

type Requests map[string]map[string]cty.Value

type ExecutionContext struct {
	Version string
	Debug   bool
}
