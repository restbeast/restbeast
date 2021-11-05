package lib

import (
	"reflect"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

func Test_parseParamsBlock(t *testing.T) {
	parser := hclparse.NewParser()
	testStringBody1 := `
		ohmy = {
      var1 = "value1"
    }
	`
	file, _ := parser.ParseHCL([]byte(testStringBody1), "test.hcl")

	parser2 := hclparse.NewParser()
	testStringBody2 := `
		var1 = "value1"
		var2 = "value2"
	`
	file2, _ := parser2.ParseHCL([]byte(testStringBody2), "test.hcl")

	type args struct {
		request     *Request
		paramsBlock *ParamsBlockCfg
		ctx         hcl.EvalContext
	}
	tests := []struct {
		name       string
		args       args
		wantErr    hcl.Diagnostics
		wantParams *map[string]string
	}{
		{
			"no block",
			args{
				request:     nil,
				paramsBlock: nil,
				ctx:         hcl.EvalContext{},
			},
			nil,
			nil,
		},
		{
			"invalid block",
			args{
				request:     nil,
				paramsBlock: &ParamsBlockCfg{Body: file.Body},
				ctx:         hcl.EvalContext{},
			},
			hcl.Diagnostics{
				&hcl.Diagnostic{
					Severity: 0,
					Summary:  "Unsuitable value type",
					Detail:   "Unsuitable value: string required",
					Subject: &hcl.Range{
						Filename: "test.hcl",
						Start:    hcl.Pos{Line: 2, Column: 10},
						End:      hcl.Pos{Line: 2, Column: 11},
					},
					Context:     nil,
					Expression:  nil,
					EvalContext: nil,
				},
			},
			nil,
		},
		{
			"valid block",
			args{
				request: &Request{
					Method:            "",
					Url:               "",
					Headers:           Headers{},
					Body:              nil,
					Params:            nil,
					EvalContext:       EvalContext{},
					PrecedingRequests: nil,
					Response:          nil,
					ExecutionContext:  nil,
					RoundTripper:      nil,
				},
				paramsBlock: &ParamsBlockCfg{Body: file2.Body},
				ctx:         hcl.EvalContext{},
			},
			nil,
			&map[string]string{
				"var1": "value1",
				"var2": "value2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := parseParamsBlock(tt.args.request, tt.args.paramsBlock, tt.args.ctx)
				if tt.wantErr != nil && !reflect.DeepEqual(got.Error(), tt.wantErr.Error()) {
					t.Errorf("parseParamsBlock() = %v, want %v", got.Error(), tt.wantErr.Error())
				}

				if tt.wantParams != nil && !reflect.DeepEqual(tt.args.request.Params, *tt.wantParams) {
					t.Errorf("parseParamsBlock() = %v, want %v", tt.args.request.Params, tt.wantParams)
				}
			},
		)
	}
}
