package lib

import (
	"github.com/zclconf/go-cty/cty"
	"reflect"
	"testing"
)

// @TODO Add more test cases
func Test_parseEnv(t *testing.T) {
	type args struct {
		env             string
		rawEnvironments EnvironmentCfgs
	}

	execCtx := ExecutionContext{
		Version: "test",
		Debug:   false,
	}

	set1 := EnvironmentCfgs{
		&EnvironmentCfg{
			Name:      "",
			Default:   false,
			Secrets:   nil,
			Variables: nil,
		},
	}
	args1 := args{
		env:             "",
		rawEnvironments: set1,
	}

	tests := []struct {
		name    string
		args    args
		want    *cty.Value
		wantErr bool
	}{
		{"set1", args1, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEnv(tt.args.env, tt.args.rawEnvironments, &execCtx)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if !reflect.DeepEqual(got, *tt.want) {
					t.Errorf("parseEnv() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
