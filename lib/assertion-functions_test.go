package lib

import (
	"github.com/zclconf/go-cty/cty"
	"testing"
)

func Test_assertEmail(t *testing.T) {
	tests := []struct {
		name string
		args []cty.Value
		want cty.Value
	}{
		{"valid email", []cty.Value{cty.StringVal("test@test.com")}, cty.StringVal("PASS")},
		{"invalid email", []cty.Value{cty.StringVal("test.com")}, cty.StringVal(`want: assertEmail()
got: test.com`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := assertionFunctionList["assertEmail"].Impl(tt.args, cty.String); tt.want.NotEqual(got).True() {
				t.Errorf("prepParams() = %v, want %v", got.AsString(), tt.want.AsString())
			}
		})
	}
}

func Test_assertEqual(t *testing.T) {
	tests := []struct {
		name string
		args []cty.Value
		want cty.Value
	}{
		{"equal values", []cty.Value{cty.StringVal("valueA"), cty.StringVal("valueA")}, cty.StringVal("PASS")},
		{"not equal values", []cty.Value{cty.StringVal("valueA"), cty.StringVal("valueB")}, cty.StringVal(`want: "valueA"
got: "valueB"`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := assertionFunctionList["assertEqual"].Impl(tt.args, cty.String); tt.want.NotEqual(got).True() {
				t.Errorf("prepParams() = %v, want %v", got.AsString(), tt.want.AsString())
			}
		})
	}
}

func Test_assertNotEqual(t *testing.T) {
	tests := []struct {
		name string
		args []cty.Value
		want cty.Value
	}{
		{"equal values", []cty.Value{cty.StringVal("valueA"), cty.StringVal("valueA")}, cty.StringVal(`not want: "valueA"
got: "valueA"`)},
		{"not equal values", []cty.Value{cty.StringVal("valueA"), cty.StringVal("valueB")}, cty.StringVal("PASS")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := assertionFunctionList["assertNotEqual"].Impl(tt.args, cty.String); tt.want.NotEqual(got).True() {
				t.Errorf("prepParams() = %v, want %v", got.AsString(), tt.want.AsString())
			}
		})
	}
}
