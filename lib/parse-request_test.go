package lib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

func Test_getRequestObjSpec(t *testing.T) {
	tests := []struct {
		name string
		want hcldec.ObjectSpec
	}{
		{
			"success", hcldec.ObjectSpec{
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
				"cookies": &hcldec.AttrSpec{
					Name:     "cookies",
					Required: false,
					Type:     cty.Map(cty.String),
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := getRequestObjSpec(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("getObjSpec() = %v, want %v", got, tt.want)
				}
			},
		)
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
		t.Run(
			tt.name, func(t *testing.T) {
				if got := getUniqueDependencies(tt.args.intSlice); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("getUniqueDependencies() = %v, want %v", got, tt.want)
				}
			},
		)
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
		t.Run(
			tt.name, func(t *testing.T) {
				gotDependencies, gotRestDiagMsgs := getPossibleDependencies(tt.args.diags)
				if !reflect.DeepEqual(gotDependencies, tt.wantDependencies) {
					t.Errorf("getPossibleDependencies() gotDependencies = %v, want %v", gotDependencies, tt.wantDependencies)
				}

				if !reflect.DeepEqual(gotRestDiagMsgs, tt.wantRestDiagMsgs) {
					t.Errorf("getPossibleDependencies() gotRestDiagMsgs = %v, want %v", gotRestDiagMsgs, tt.wantRestDiagMsgs)
				}
			},
		)
	}
}

func Test_findRequest(t *testing.T) {
	cfg1 := RequestCfg{
		Name:      "x",
		DependsOn: nil,
		Remain:    nil,
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
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := findRequest(tt.args.name, tt.args.rawRequests)
				if (err != nil) != tt.wantErr {
					t.Errorf("findRequest() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("findRequest() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_getRequest(t *testing.T) {
	type args struct {
		cfg        cty.Value
		requestCfg RequestCfg
		execCtx    *ExecutionContext
		evCtx      *EvalContext
	}

	args1 := args{
		cty.ObjectVal(
			map[string]cty.Value{
				"body": cty.StringVal(`{ x: "y" }`),
				"headers": cty.MapVal(
					map[string]cty.Value{
						"content-type": cty.StringVal("application/json"),
						"Cookie":       cty.StringVal("name=value; another-name=another-value"),
					},
				),
				"method": cty.StringVal("GET"),
				"url":    cty.StringVal("localhost"),
			},
		),
		RequestCfg{
			Name:      "",
			DependsOn: nil,
			Remain:    nil,
			Auth:      nil,
		},
		&ExecutionContext{
			Version: "",
			Debug:   false,
		},
		&EvalContext{
			Functions:     nil,
			Variables:     nil,
			Environment:   nil,
			RequestAsVars: &RequestAsVars{},
			RawRequests:   nil,
		},
	}

	args2 := args{
		cty.ObjectVal(
			map[string]cty.Value{
				"method": cty.StringVal("GET"),
				"url":    cty.StringVal("localhost"),
			},
		),
		RequestCfg{
			Name:      "",
			DependsOn: nil,
			Remain:    nil,
			Auth:      nil,
		},
		&ExecutionContext{
			Version: "",
			Debug:   false,
		},
		&EvalContext{
			Functions:     nil,
			Variables:     nil,
			Environment:   nil,
			RequestAsVars: &RequestAsVars{},
			RawRequests:   nil,
		},
	}

	args3 := args{
		cty.ObjectVal(
			map[string]cty.Value{
				"method":   cty.StringVal("GET"),
				"url":      cty.StringVal("localhost"),
				"body":     cty.ObjectVal(map[string]cty.Value{"hey": cty.StringVal("ho")}),
				"boundary": cty.StringVal("test"),
			},
		),
		RequestCfg{
			Name:      "",
			DependsOn: nil,
			Remain:    nil,
			Auth:      nil,
		},
		&ExecutionContext{
			Version: "",
			Debug:   false,
		},
		&EvalContext{
			Functions:     nil,
			Variables:     nil,
			Environment:   nil,
			RequestAsVars: &RequestAsVars{},
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
		t.Run(
			tt.name, func(t *testing.T) {
				_, err, _ := getRequest(tt.args.cfg, tt.args.requestCfg, tt.args.evCtx, tt.args.execCtx)
				if (err != nil) != tt.wantErr {
					t.Errorf("getRequest() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

func Test_getCtxEvalContext(t *testing.T) {
	var emptyMap map[string]cty.Value
	var emptyFMap map[string]function.Function
	vars := map[string]cty.Value{
		"var":     cty.ObjectVal(emptyMap),
		"request": cty.ObjectVal(emptyMap),
		"env":     {},
	}

	type args struct {
		evCtx *EvalContext
	}
	tests := []struct {
		name string
		args args
		want hcl.EvalContext
	}{
		{
			"success", args{
				&EvalContext{
					Functions:     &emptyFMap,
					Variables:     &emptyMap,
					Environment:   &cty.Value{},
					RequestAsVars: &RequestAsVars{},
					RawRequests:   nil,
				},
			}, hcl.EvalContext{
				Variables: vars,
				Functions: emptyFMap,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := getCtxEvalContext(tt.args.evCtx); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("getCtxEvalContext() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_retryWithDependency(t *testing.T) {
	type args struct {
		requestCfg *RequestCfg
		cfg        cty.Value
		diags      hcl.Diagnostics
		evCtx      *EvalContext
		execCtx    *ExecutionContext
		responses  []*Response
	}

	args1 := args{}
	args2 := args{
		diags: hcl.Diagnostics{
			&hcl.Diagnostic{
				Severity: 0,
				Summary:  "Unsupported attribute",
				Detail:   "detail",
			},
		},
	}

	args3 := args{
		diags: hcl.Diagnostics{
			&hcl.Diagnostic{
				Severity: 0,
				Summary:  "Unsupported attribute",
				Detail:   "This object does not have an attribute named \"deprequest\"",
			},
		},
		evCtx: &EvalContext{
			RequestAsVars: &RequestAsVars{},
		},
	}

	rav := RequestAsVars{}
	rav.Store("deprequest", cty.Value{})

	args4 := args{
		diags: hcl.Diagnostics{
			&hcl.Diagnostic{
				Severity: 0,
				Summary:  "Unsupported attribute",
				Detail:   "This object does not have an attribute named \"deprequest\"",
			},
		},
		evCtx: &EvalContext{RequestAsVars: &rav},
		requestCfg: &RequestCfg{
			Name:      "a-request",
			DependsOn: nil,
			Remain:    hcl.EmptyBody(),
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
		t.Run(
			tt.name, func(t *testing.T) {
				fmt.Println(tt.name)
				_, got1, err := retryWithDependency(
					tt.args.requestCfg, tt.args.cfg, tt.args.diags, tt.args.evCtx, tt.args.execCtx, tt.args.responses,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("retryWithDependency() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got1, tt.want1) {
					t.Errorf("retryWithDependency() got1 = %v, want %v", got1, tt.want1)
				}
			},
		)
	}
}

func Test_getHeadersAsMap(t *testing.T) {
	type args struct {
		cfg cty.Value
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			"success",
			args{cty.ObjectVal(map[string]cty.Value{"headers": cty.MapVal(map[string]cty.Value{"hey": cty.StringVal("ho")})})},
			map[string]string{"hey": "ho"}, false,
		},
		{"no headers", args{cty.ObjectVal(map[string]cty.Value{})}, map[string]string{}, false},
		{"error", args{cty.ObjectVal(map[string]cty.Value{"headers": cty.StringVal("o.O")})}, map[string]string{}, true},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := getHeadersAsMap(tt.args.cfg)
				if (err != nil) != tt.wantErr {
					t.Errorf("getHeadersAsMap() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("getHeadersAsMap() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_parseRequest(t *testing.T) {
	type args struct {
		name    string
		evCtx   *EvalContext
		execCtx *ExecutionContext
	}

	args1 := args{
		name:    "not-found-request",
		evCtx:   &EvalContext{},
		execCtx: nil,
	}

	parser := hclparse.NewParser()
	testStringBody1 := `
		url = "test"
		method = "get"
	`
	file, _ := parser.ParseHCL([]byte(testStringBody1), "test.hcl")
	args2 := args{
		name: "a-request",
		evCtx: &EvalContext{
			RawRequests: RequestCfgs{
				&RequestCfg{
					Name:      "a-request",
					DependsOn: []string{},
					Remain:    file.Body,
					Auth:      nil,
					Params:    nil,
				},
			},
			RequestAsVars: &RequestAsVars{},
			RawDynamics:   VariableCfgs{},
			RawTests:      TestCfgs{},
		},
		execCtx: nil,
	}

	parse3 := hclparse.NewParser()
	testStringBody3 := `
		url = "test+${request.parent.body.hey}"
		method = "get"
	`
	file3, _ := parse3.ParseHCL([]byte(testStringBody3), "test.hcl")
	args3 := args{
		name: "a-request",
		evCtx: &EvalContext{
			RawRequests: RequestCfgs{
				&RequestCfg{
					Name:      "a-request",
					DependsOn: []string{},
					Remain:    file3.Body,
					Auth:      nil,
					Params:    nil,
				},
			},
			RequestAsVars: &RequestAsVars{},
			RawDynamics:   VariableCfgs{},
			RawTests:      TestCfgs{},
		},
		execCtx: nil,
	}

	parse4 := hclparse.NewParser()
	testStringBody4 := `
		url = "test+${request.parent.body.x}"
		method = "get"
	`
	file4, _ := parse4.ParseHCL([]byte(testStringBody4), "test.hcl")
	parentDep := RequestAsVars{}
	parentDep.Store(
		"parent", cty.MapVal(map[string]cty.Value{"body": cty.ObjectVal(map[string]cty.Value{"x": cty.StringVal("y")})}),
	)

	args4 := args{
		name: "a-request",
		evCtx: &EvalContext{
			RawRequests: RequestCfgs{
				&RequestCfg{
					Name:      "a-request",
					DependsOn: []string{},
					Remain:    file4.Body,
					Auth:      nil,
					Params:    nil,
				},
			},
			RequestAsVars: &parentDep,
			RawDynamics:   VariableCfgs{},
			RawTests:      TestCfgs{},
		},
		execCtx: nil,
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"error, request not found", args1, true},
		{"a-request", args2, false},
		{"error with-deps", args3, true},
		{"request-with-resolved-deps", args4, false},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				_, err := parseRequest(tt.args.name, tt.args.evCtx, tt.args.execCtx)
				if (err != nil) != tt.wantErr {
					t.Errorf("parseRequest() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

func Test_getCookiesAsMap(t *testing.T) {
	type args struct {
		cfg cty.Value
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			"success",
			args{cty.ObjectVal(map[string]cty.Value{"cookies": cty.MapVal(map[string]cty.Value{"hey": cty.StringVal("ho")})})},
			map[string]string{"hey": "ho"}, false,
		},
		{"no headers", args{cty.ObjectVal(map[string]cty.Value{})}, map[string]string{}, false},
		{"error", args{cty.ObjectVal(map[string]cty.Value{"cookies": cty.StringVal("o.O")})}, map[string]string{}, true},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := getCookiesAsMap(tt.args.cfg)
				if (err != nil) != tt.wantErr {
					t.Errorf("getCookiesAsMap() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("getCookiesAsMap() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_processResponseBody(t *testing.T) {
	jsonCt := "application/json; utf-8"
	textCt := "text/plain; utf-8"

	type args struct {
		contentType *string
		body        []byte
	}
	tests := []struct {
		name    string
		args    args
		want    cty.Value
		wantErr bool
	}{
		{"empty body", args{&jsonCt, []byte("")}, cty.Value{}, false},
		{"text body", args{nil, []byte("text")}, cty.StringVal("text"), false},
		{"text body", args{&textCt, []byte("text")}, cty.StringVal("text"), false},
		{"json body", args{&jsonCt, []byte(`["text"]`)}, cty.TupleVal([]cty.Value{cty.StringVal("text")}), false},
		{"error json body", args{&jsonCt, []byte(`=::`)}, cty.Value{}, true},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := processResponseBody(tt.args.contentType, tt.args.body)
				if (err != nil) != tt.wantErr {
					t.Errorf("processResponseBody() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("processResponseBody() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
