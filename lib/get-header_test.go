package lib

import (
	"reflect"
	"testing"
)

func Test_getHeader(t *testing.T) {
	keyNotFoundH := map[string]string{"existent-key": "heyho-1"}

	exactKeyFoudH := map[string]string{"existent-key": "heyho-2"}
	want2 := "heyho-2"

	keyFound := map[string]string{"existent-key": "heyho-3"}
	want3 := "heyho-3"

	type args struct {
		key     string
		headers *map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"key not found", args{"non-existent-key", &keyNotFoundH}, ""},
		{"exact key found", args{"existent-key", &exactKeyFoudH}, want2},
		{"key found", args{"Existent-Key", &keyFound}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			getHeader(tt.args.key, tt.args.headers, &got)
			if got != "" && tt.want != "" {
				t.Errorf("getHeader() = %v, want %v", got, tt.want)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
