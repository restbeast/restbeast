package lib

import (
	"reflect"
	"testing"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

func Test_parseVariables(t *testing.T) {
	emptyVarCfg := []*VariableCfg{}

	parser := hclparse.NewParser()
	test1StringBody := `
		value = "testvalue"
	`
	file1, _ := parser.ParseHCL([]byte(test1StringBody), "test.hcl")
	varCfg1 := []*VariableCfg{
		{"varName", file1.Body},
	}

	parser2 := hclparse.NewParser()
	test2StringBody := `
		oh-my {}
	`
	file2, _ := parser2.ParseHCL([]byte(test2StringBody), "test.hcl")
	varCfg2 := []*VariableCfg{
		{"varName", file2.Body},
	}

	type args struct {
		rawVars   []*VariableCfg
		functions map[string]function.Function
	}
	tests := []struct {
		name    string
		args    args
		want    *map[string]cty.Value
		wantErr bool
	}{
		{"test1", args{emptyVarCfg, map[string]function.Function{}}, &map[string]cty.Value{}, false},
		{
			"test2", args{varCfg1, map[string]function.Function{}},
			&map[string]cty.Value{"varName": cty.StringVal("testvalue")}, false,
		},
		{"test2", args{varCfg2, map[string]function.Function{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := parseVariables(tt.args.rawVars, tt.args.functions)
				if (err != nil) != tt.wantErr {
					t.Errorf("parseVariables() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("parseVariables() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
