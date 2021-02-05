package lib

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"net/http"
	"reflect"
	"testing"
)

func Test_lowercaseHeaders(t *testing.T) {
	val1 := []string{
		"header-value-1",
	}
	headers := http.Header{}
	headers["Key1"] = val1

	wantHeaders := headers
	wantHeaders["key1"] = val1

	type args struct {
		headers http.Header
	}

	tests := []struct {
		name string
		args args
		want http.Header
	}{
		{"adds lower case headers", args{headers}, wantHeaders},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowercaseHeaders(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowercaseHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRequestObjSpec(t *testing.T) {
	tests := []struct {
		name string
		want hcldec.ObjectSpec
	}{
		{"success", hcldec.ObjectSpec{
			"method": &hcldec.AttrSpec{
				Name:     "method",
				Required: true,
				Type:     cty.String,
			},
			"url": &hcldec.AttrSpec{
				Name:     "url",
				Required: true,
				Type:     cty.String,
			},
			"headers": &hcldec.AttrSpec{
				Name:     "headers",
				Required: false,
				Type:     cty.Map(cty.String),
			},
			"body": &hcldec.AttrSpec{
				Name:     "body",
				Required: false,
				Type:     cty.DynamicPseudoType,
			},
			"depends_on": &hcldec.AttrSpec{
				Name:     "depends_on",
				Required: false,
				Type:     cty.List(cty.String),
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRequestObjSpec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getObjSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUniqueDependencies(t *testing.T) {
	type args struct {
		intSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test1", args{[]string{"a", "b"}}, []string{"a", "b"}},
		{"test2", args{[]string{"a", "b", "c", "b"}}, []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUniqueDependencies(tt.args.intSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUniqueDependencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowercaseHeaders1(t *testing.T) {
	testGot1 := map[string][]string{
		"content-type": []string{"application/json"},
	}

	testGot2 := map[string][]string{
		"Content-Type": []string{"application/json"},
	}

	testWant2 := map[string][]string{
		"Content-Type": []string{"application/json"},
		"content-type": []string{"application/json"},
	}

	type args struct {
		headers http.Header
	}
	tests := []struct {
		name string
		args args
		want http.Header
	}{
		{"test1", args{testGot1}, testGot1},
		{"test1", args{testGot2}, testWant2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowercaseHeaders(tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowercaseHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPossibleDependencies(t *testing.T) {
	diag1 := hcl.Diagnostics{
		&hcl.Diagnostic{
			Severity: 0,
			Summary:  "",
			Detail:   "",
		},
	}

	var wantDep1 []string
	var wantRes1 []string

	diag2 := hcl.Diagnostics{
		&hcl.Diagnostic{
			Severity: 0,
			Summary:  "Unsupported attribute",
			Detail:   `This object does not have an attribute named "dep1"`,
		},
	}

	wantDep2 := []string{"dep1"}
	var wantRes2 []string

	diag3 := hcl.Diagnostics{
		&hcl.Diagnostic{
			Severity:    0,
			Summary:     "Unsupported attribute",
			Detail:      `some other message`,
			Subject:     &hcl.Range{},
			Context:     &hcl.Range{},
			Expression:  nil,
			EvalContext: nil,
		},
	}

	var wantDep3 []string
	wantRes3 := []string{fmt.Sprint(diag3[0])}

	type args struct {
		diags hcl.Diagnostics
	}
	tests := []struct {
		name             string
		args             args
		wantDependencies []string
		wantRestDiagMsgs []string
	}{
		{"empty", args{diag1}, wantDep1, wantRes1},
		{"with dep", args{diag2}, wantDep2, wantRes2},
		{"with res", args{diag3}, wantDep3, wantRes3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDependencies, gotRestDiagMsgs := getPossibleDependencies(tt.args.diags)
			if !reflect.DeepEqual(gotDependencies, tt.wantDependencies) {
				t.Errorf("getPossibleDependencies() gotDependencies = %v, want %v", gotDependencies, tt.wantDependencies)
			}

			if !reflect.DeepEqual(gotRestDiagMsgs, tt.wantRestDiagMsgs) {
				t.Errorf("getPossibleDependencies() gotRestDiagMsgs = %v, want %v", gotRestDiagMsgs, tt.wantRestDiagMsgs)
			}
		})
	}
}

func Test_findRequest(t *testing.T) {
	cfg1 := RequestCfg{
		Name:      "x",
		DependsOn: nil,
		Body:      nil,
	}

	type args struct {
		name        string
		rawRequests RequestCfgs
	}
	tests := []struct {
		name    string
		args    args
		want    *RequestCfg
		wantErr bool
	}{
		{"got 1", args{"x", RequestCfgs{&cfg1}}, &cfg1, false},
		{"got err", args{"y", RequestCfgs{&cfg1}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findRequest(tt.args.name, tt.args.rawRequests)
			if (err != nil) != tt.wantErr {
				t.Errorf("findRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRequest(t *testing.T) {
	type args struct {
		cfg        cty.Value
		requestCfg RequestCfg
		execCtx    *ExecutionContext
		evCtx      EvalContext
	}

	args1 := args{
		cty.ObjectVal(map[string]cty.Value{
			"body": cty.StringVal(`{ x: "y" }`),
			"headers": cty.MapVal(map[string]cty.Value{
				"content-type": cty.StringVal("application/json"),
			}),
			"method": cty.StringVal("GET"),
			"url":    cty.StringVal("localhost"),
		}),
		RequestCfg{
			Name:      "",
			DependsOn: nil,
			Body:      nil,
			Auth:      nil,
		},
		&ExecutionContext{
			Version: "",
			Debug:   false,
		},
		EvalContext{
			Functions:     nil,
			Variables:     nil,
			Environment:   nil,
			RequestAsVars: nil,
			RawRequests:   nil,
		},
	}

	args2 := args{
		cty.ObjectVal(map[string]cty.Value{
			"method": cty.StringVal("GET"),
			"url":    cty.StringVal("localhost"),
		}),
		RequestCfg{
			Name:      "",
			DependsOn: nil,
			Body:      nil,
			Auth:      nil,
		},
		&ExecutionContext{
			Version: "",
			Debug:   false,
		},
		EvalContext{
			Functions:     nil,
			Variables:     nil,
			Environment:   nil,
			RequestAsVars: nil,
			RawRequests:   nil,
		},
	}

	args3 := args{
		cty.ObjectVal(map[string]cty.Value{
			"method":   cty.StringVal("GET"),
			"url":      cty.StringVal("localhost"),
			"body":     cty.ObjectVal(map[string]cty.Value{"hey": cty.StringVal("ho")}),
			"boundary": cty.StringVal("test"),
		}),
		RequestCfg{
			Name:      "",
			DependsOn: nil,
			Body:      nil,
			Auth:      nil,
		},
		&ExecutionContext{
			Version: "",
			Debug:   false,
		},
		EvalContext{
			Functions:     nil,
			Variables:     nil,
			Environment:   nil,
			RequestAsVars: nil,
			RawRequests:   nil,
		},
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args1, false},
		{"test2", args2, false},
		{"test3", args3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err, _ := getRequest(tt.args.cfg, tt.args.requestCfg, tt.args.evCtx, tt.args.execCtx)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_getCtxEvalContext(t *testing.T) {
	var emptyMap map[string]cty.Value
	var emptyFMap map[string]function.Function
	vars := map[string]cty.Value{
		"var":     cty.ObjectVal(emptyMap),
		"request": cty.ObjectVal(emptyMap),
		"env":     cty.Value{},
	}

	type args struct {
		evCtx EvalContext
	}
	tests := []struct {
		name string
		args args
		want hcl.EvalContext
	}{
		{
			"success", args{EvalContext{
				Functions:     &emptyFMap,
				Variables:     &emptyMap,
				Environment:   &cty.Value{},
				RequestAsVars: emptyMap,
				RawRequests:   nil,
			}}, hcl.EvalContext{
				Variables: vars,
				Functions: emptyFMap,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCtxEvalContext(tt.args.evCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCtxEvalContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_retryWithDependency(t *testing.T) {
	type args struct {
		requestCfg *RequestCfg
		cfg        cty.Value
		diags      hcl.Diagnostics
		evCtx      EvalContext
		execCtx    *ExecutionContext
		responses  []*Response
	}

	args1 := args{}
	args2 := args{
		diags: hcl.Diagnostics{&hcl.Diagnostic{
			Severity: 0,
			Summary:  "Unsupported attribute",
			Detail:   "detail",
		}},
	}

	args3 := args{
		diags: hcl.Diagnostics{&hcl.Diagnostic{
			Severity: 0,
			Summary:  "Unsupported attribute",
			Detail:   "This object does not have an attribute named \"deprequest\"",
		}},
	}

	args4 := args{
		diags: hcl.Diagnostics{&hcl.Diagnostic{
			Severity: 0,
			Summary:  "Unsupported attribute",
			Detail:   "This object does not have an attribute named \"deprequest\"",
		}},
		evCtx: EvalContext{RequestAsVars: RequestAsVars{"deprequest": cty.Value{}}},
		requestCfg: &RequestCfg{
			Name:      "a-request",
			DependsOn: nil,
			Body:      hcl.EmptyBody(),
			Auth:      nil,
		},
	}

	tests := []struct {
		name    string
		args    args
		want1   []*Response
		wantErr bool
	}{
		{"no dep, no diag", args1, nil, false},
		{"no dep, rest diag", args2, nil, true},
		{"not found dep", args3, nil, true},
		{"dep", args4, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1, err := retryWithDependency(tt.args.requestCfg, tt.args.cfg, tt.args.diags, tt.args.evCtx, tt.args.execCtx, tt.args.responses)
			if (err != nil) != tt.wantErr {
				t.Errorf("retryWithDependency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("retryWithDependency() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
