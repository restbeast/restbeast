package lib

import (
	"os"
	"reflect"
	"testing"
)

func Test_secretEngineEnvVar(t *testing.T) {
	//os.Getenv("restbeast_var_" + systemKey)
	_ = os.Setenv("restbeast_var_KEY1", "value1")
	paths := map[string]string{
		"path1": "KEY1",
	}

	parsedSecret := ParsedSecret{
		"path1": "value1",
	}

	type args struct {
		paths map[string]string
	}
	tests := []struct {
		name        string
		args        args
		wantSecrets ParsedSecret
	}{
		{"success", args{paths}, parsedSecret},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if gotSecrets := secretEngineEnvVar(tt.args.paths); !reflect.DeepEqual(gotSecrets, tt.wantSecrets) {
					t.Errorf("secretEngineEnvVar() = %v, want %v", gotSecrets, tt.wantSecrets)
				}
			},
		)
	}
}
