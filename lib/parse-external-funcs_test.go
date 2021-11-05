package lib

import (
	"reflect"
	"testing"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

func Test_prepParams(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want []function.Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := prepParams(tt.args.args); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("prepParams() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_prepParams1(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want []function.Parameter
	}{
		{
			"test1", args{[]string{"string", "string"}}, []function.Parameter{
			{Name: "arg0", Type: cty.String}, {Name: "arg1", Type: cty.String},
			},
		},
		{
			"test2", args{[]string{"list", "string"}}, []function.Parameter{
			{Name: "arg0", Type: cty.List(cty.DynamicPseudoType)},
			{Name: "arg1", Type: cty.String},
			},
		},
		{
			"test3", args{[]string{"map", "string"}}, []function.Parameter{
			{Name: "arg0", Type: cty.Map(cty.DynamicPseudoType)},
			{Name: "arg1", Type: cty.String},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := prepParams(tt.args.args); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("prepParams() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_prepArgs(t *testing.T) {
	type args struct {
		exFn *ExternalFunctionCfg
		args []cty.Value
	}
	tests := []struct {
		name         string
		args         args
		wantExecArgs []string
		wantErr      bool
	}{
		{
			"test1", args{&ExternalFunctionCfg{"", "", "", []string{"string"}}, []cty.Value{cty.StringVal("test")}},
			[]string{"", "test"}, false,
		},
		{
			"test2", args{&ExternalFunctionCfg{"", "", "", []string{"number"}}, []cty.Value{cty.NumberIntVal(1)}},
			[]string{"", "1"}, false,
		},
		{
			"test3", args{
				&ExternalFunctionCfg{"", "", "", []string{"list"}},
				[]cty.Value{cty.ListVal([]cty.Value{cty.StringVal("val1"), cty.StringVal("val2")})},
			}, []string{"", `["val1","val2"]`}, false,
		},
		{
			"test4", args{
				&ExternalFunctionCfg{"", "", "", []string{"map"}},
				[]cty.Value{cty.MapVal(map[string]cty.Value{"key1": cty.StringVal("hello")})},
			}, []string{"", `{"key1":"hello"}`}, false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotExecArgs, err := prepArgs(tt.args.exFn, tt.args.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("prepArgs() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotExecArgs, tt.wantExecArgs) {
					t.Errorf("prepArgs() gotExecArgs = %v, want %v", gotExecArgs, tt.wantExecArgs)
				}
			},
		)
	}
}
