package lib

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"testing"
)

func Test_parseBasicAuth(t *testing.T) {
	type args struct {
		request   *Request
		basicAuth BasicAuthCfg
		ctx       hcl.EvalContext
	}

	parser1 := hclparse.NewParser()
	test1StringBody := `
		username = "username-test-1"
		password = "password-test-1"
	`
	file1, _ := parser1.ParseHCL([]byte(test1StringBody), "test.hcl")
	cfg1 := BasicAuthCfg{file1.Body}
	request1 := Request{}
	args1 := args{
		request:   &request1,
		basicAuth: cfg1,
		ctx:       hcl.EvalContext{},
	}

	parser2 := hclparse.NewParser()
	test2StringBody := `
		username = "username-test-2"
	`
	file2, _ := parser2.ParseHCL([]byte(test2StringBody), "test.hcl")
	cfg2 := BasicAuthCfg{file2.Body}
	args2 := args{
		request:   &Request{},
		basicAuth: cfg2,
		ctx:       hcl.EvalContext{},
	}

	tests := []struct {
		name   string
		args   args
		err    bool
		header string
	}{
		{"success case", args1, false, b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", "username-test-1", "password-test-1")))},
		{"error case", args2, true, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := parseBasicAuth(tt.args.request, tt.args.basicAuth, tt.args.ctx)
			if tt.err && diags == nil {
				t.Errorf("parseBasicAuth() = %v, want %v", diags, tt.err)
			}

			if !tt.err && fmt.Sprintf("Basic %s", tt.header) != tt.args.request.Headers["Authorization"] {
				t.Errorf("parseBasicAuth() = %v, want %v", tt.args.request.Headers["Authorization"], fmt.Sprintf("Basic %s", tt.header))
			}
		})
	}
}

func Test_parseAuthBlock(t *testing.T) {
	type args struct {
		request   *Request
		authBlock *AuthCfg
		ctx       hcl.EvalContext
	}

	cfg1 := AuthCfg{}
	args1 := args{
		&Request{},
		&cfg1,
		hcl.EvalContext{},
	}

	parser1 := hclparse.NewParser()
	test1StringBody := `
		username = "username-test-1"
		password = "password-test-1"
	`
	file1, _ := parser1.ParseHCL([]byte(test1StringBody), "test.hcl")
	cfg2 := BasicAuthCfg{file1.Body}
	args2 := args{
		&Request{},
		&AuthCfg{BasicAuth: &cfg2},
		hcl.EvalContext{},
	}

	parser3 := hclparse.NewParser()
	test3StringBody := `
		username = "username-test-2"
	`
	file3, _ := parser3.ParseHCL([]byte(test3StringBody), "test.hcl")
	cfg3 := BasicAuthCfg{file3.Body}
	args3 := args{
		&Request{},
		&AuthCfg{BasicAuth: &cfg3},
		hcl.EvalContext{},
	}

	parser4 := hclparse.NewParser()
	test4StringBody := `
		token = "a-token"
	`
	file4, _ := parser4.ParseHCL([]byte(test4StringBody), "test.hcl")
	cfg4 := BearerAuthCfg{file4.Body}
	args4 := args{
		&Request{},
		&AuthCfg{BearerAuth: &cfg4},
		hcl.EvalContext{},
	}

	parser5 := hclparse.NewParser()
	test5StringBody := `
		no-token = "a-token"
	`
	file5, _ := parser5.ParseHCL([]byte(test5StringBody), "test.hcl")
	cfg5 := BearerAuthCfg{file5.Body}
	args5 := args{
		&Request{},
		&AuthCfg{BearerAuth: &cfg5},
		hcl.EvalContext{},
	}

	tests := []struct {
		name string
		args args
		err  bool
	}{
		{"nil authBlock", args{nil, nil, hcl.EvalContext{}}, false},
		{"empty authBlock", args1, false},
		{"with basic auth block", args2, false},
		{"with basic auth block with diags", args3, true},
		{"with bearer auth block", args4, false},
		{"with bearer auth block with diags", args5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseAuthBlock(tt.args.request, tt.args.authBlock, tt.args.ctx); !tt.err && got != nil {
				t.Errorf("parseAuthBlock() = %v", got)
			}
		})
	}
}

func Test_parseBearerAuth(t *testing.T) {
	type args struct {
		request   *Request
		basicAuth BearerAuthCfg
		ctx       hcl.EvalContext
	}

	parser1 := hclparse.NewParser()
	test1StringBody := `
		token = "token-1"
	`
	file1, _ := parser1.ParseHCL([]byte(test1StringBody), "test.hcl")
	cfg1 := BearerAuthCfg{file1.Body}
	request1 := Request{}
	args1 := args{
		request:   &request1,
		basicAuth: cfg1,
		ctx:       hcl.EvalContext{},
	}

	parser2 := hclparse.NewParser()
	test2StringBody := `
		username = "username-test-2"
	`
	file2, _ := parser2.ParseHCL([]byte(test2StringBody), "test.hcl")
	cfg2 := BearerAuthCfg{file2.Body}
	args2 := args{
		request:   &Request{},
		basicAuth: cfg2,
		ctx:       hcl.EvalContext{},
	}

	tests := []struct {
		name   string
		args   args
		err    bool
		header string
	}{
		{"success case", args1, false, "token-1"},
		{"error case", args2, true, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := parseBearerAuth(tt.args.request, tt.args.basicAuth, tt.args.ctx)

			if tt.err && diags == nil {
				t.Errorf("parseBasicAuth() = %v, want %v", diags, tt.err)
			}

			if !tt.err && fmt.Sprintf("Bearer %s", tt.header) != tt.args.request.Headers["Authorization"] {
				t.Errorf("parseBearerAuth() = %v, want %v", tt.args.request.Headers["Authorization"], fmt.Sprintf("Bearer %s", tt.header))
			}
		})
	}
}
