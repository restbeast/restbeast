package lib

import (
	"reflect"
	"testing"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

func Test_findTest(t *testing.T) {
	cfg1 := TestCfg{
		Name: "x",
		Body: nil,
	}

	type args struct {
		name     string
		rawTests TestCfgs
	}
	tests := []struct {
		name    string
		args    args
		want    *TestCfg
		wantErr bool
	}{
		{"got 1", args{"x", TestCfgs{&cfg1}}, &cfg1, false},
		{"got err", args{"y", TestCfgs{&cfg1}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := findTest(tt.args.name, tt.args.rawTests)
				if (err != nil) != tt.wantErr {
					t.Errorf("findTest() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("findTest() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_updateEvalContextWithTestFns(t *testing.T) {
	fnList := map[string]function.Function{}
	evCtx := EvalContext{Functions: &fnList}

	t.Run(
		"Add all assertion functions to evaluation context", func(t *testing.T) {
			updateEvalContextWithTestFns(&evCtx)

			if len(fnList) != len(assertionFunctionList) {
				t.Errorf("updateEvalContextWithTestFns, failed to update functions list")
			}
		},
	)
}

func Test_prepareResults(t *testing.T) {
	assertions := map[string]cty.Value{
		"should-pass": cty.StringVal("PASS"),
		"should-fail": cty.StringVal("an other text"),
	}

	results := prepareResults("test name", assertions)
	if len(results.Assertions) != 2 {
		t.Errorf("prepareResults(), , failed to generate test result")
	}
}
