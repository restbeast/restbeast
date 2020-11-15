package lib

import (
	"reflect"
	"testing"
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := findTest(tt.args.name, tt.args.rawTests)
			if (err != nil) != tt.wantErr {
				t.Errorf("findTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
