package lib

import (
	"reflect"
	"testing"
)

func TestCookies_Add(t *testing.T) {
	type fields struct {
		kv map[string]string
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Cookies
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cookies := &Cookies{
				kv: tt.fields.kv,
			}
			if got := cookies.Add(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
