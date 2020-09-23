package lib

import (
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
	"os"
	"reflect"
	"testing"
)

func Test_parseEnv(t *testing.T) {
	parser := hclparse.NewParser()
	testStringBody := `
		variables = {
      var1 = "value1"
    }
	`
	file, _ := parser.ParseHCL([]byte(testStringBody), "test.hcl")

	parser2 := hclparse.NewParser()
	testStringBody2 := `
		ohmy = {
      var1 = "value1"
    }
	`
	file2, _ := parser2.ParseHCL([]byte(testStringBody2), "test.hcl")

	type args struct {
		env             string
		rawEnvironments EnvironmentCfgs
	}

	execCtx := ExecutionContext{
		Version: "test",
		Debug:   false,
	}

	tests := []struct {
		name    string
		args    args
		want    *cty.Value
		wantErr bool
	}{
		{
			"test1",
			args{
				env: "",
				rawEnvironments: EnvironmentCfgs{
					&EnvironmentCfg{
						Name:      "",
						Default:   false,
						Secrets:   nil,
						Variables: nil,
					},
				},
			},
			nil,
			false,
		},
		{
			"test2",
			args{
				env: "",
				rawEnvironments: EnvironmentCfgs{
					&EnvironmentCfg{
						Name:      "env1",
						Default:   true,
						Secrets:   nil,
						Variables: file.Body,
					},
				},
			},
			nil,
			false,
		},
		{
			"test3",
			args{
				env: "",
				rawEnvironments: EnvironmentCfgs{
					&EnvironmentCfg{
						Name:      "env1",
						Default:   true,
						Secrets:   nil,
						Variables: file2.Body,
					},
				},
			},
			nil,
			true,
		},
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

func Test_parseSecrets(t *testing.T) {
	_ = os.Setenv("restbeast_var_TESTENVAR", "testEnvValue")

	tests := []struct {
		name       string
		secretCfgs SecretCfgs
		want       AllParsedSecrets
	}{
		{"test1", SecretCfgs{&SecretCfg{Name: "", Type: "", Paths: nil}}, AllParsedSecrets{}},
		{"test2", SecretCfgs{&SecretCfg{Name: "from_shell_env", Type: "env-var", Paths: map[string]string{"test": "TESTENVAR"}}}, AllParsedSecrets{"from_shell_env": map[string]string{"test": "testEnvValue"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSecrets(tt.secretCfgs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSecrets() = %v, want %v", got, tt.want)
			}
		})
	}
}
