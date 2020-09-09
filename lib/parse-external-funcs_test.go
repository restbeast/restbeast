package lib

import (
	"github.com/zclconf/go-cty/cty/function"
	"reflect"
	"testing"
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
		t.Run(tt.name, func(t *testing.T) {
			if got := prepParams(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
