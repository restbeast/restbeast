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
		{"invalid email", []cty.Value{cty.StringVal("test.com")}, cty.StringVal(`expected test.com to be a valid email`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := assertionFunctionList["assertEmail"].Impl(tt.args, cty.String); tt.want.NotEqual(got).True() {
				t.Errorf("prepParams() = %v, want %v", got.AsString(), tt.want.AsString())
			}
		})
	}
}

func Test_assertUuidv4(t *testing.T) {
	tests := []struct {
		name string
		args []cty.Value
		want cty.Value
	}{
		{"valid uuidv4", []cty.Value{cty.StringVal("3df3f84d-0142-4d33-80a0-7e23b5b0eba6")}, cty.StringVal("PASS")},
		{"invalid uuidv4", []cty.Value{cty.StringVal("i-am-not-a-valid-uuid")}, cty.StringVal(`expected i-am-not-a-valid-uuid to be a valid uuidv4`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := assertionFunctionList["assertUuidv4"].Impl(tt.args, cty.String); tt.want.NotEqual(got).True() {
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
		{"not equal values", []cty.Value{cty.StringVal("valueA"), cty.StringVal("valueB")}, cty.StringVal(`expected: "valueA"
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
		{"equal values", []cty.Value{cty.StringVal("valueA"), cty.StringVal("valueA")}, cty.StringVal(`not expected: "valueA"
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
