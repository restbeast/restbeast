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
				t.Errorf("assertEmail() = %v, want %v", got.AsString(), tt.want.AsString())
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
				t.Errorf("assertUuidv4() = %v, want %v", got.AsString(), tt.want.AsString())
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
				t.Errorf("assertEqual() = %v, want %v", got.AsString(), tt.want.AsString())
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
				t.Errorf("assertNotEqual() = %v, want %v", got.AsString(), tt.want.AsString())
			}
		})
	}
}

func Test_assertGreaterThan(t *testing.T) {
	tests := []struct {
		name string
		args []cty.Value
		want cty.Value
	}{
		{"greater value", []cty.Value{cty.NumberIntVal(int64(2)), cty.NumberIntVal(int64(1))}, cty.StringVal("PASS")},
		{"not greater value", []cty.Value{cty.NumberIntVal(int64(1)), cty.NumberIntVal(int64(2))}, cty.StringVal(`expected 1 to be greater than 2`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := assertionFunctionList["assertGreaterThan"].Impl(tt.args, cty.String); tt.want.NotEqual(got).True() {
				t.Errorf("assertGreaterThan() = %v, want %v", got.AsString(), tt.want.AsString())
			}
		})
	}
}

func Test_assertGreaterThanOrEqualTo(t *testing.T) {
	tests := []struct {
		name string
		args []cty.Value
		want cty.Value
	}{
		{"greater value", []cty.Value{cty.NumberIntVal(int64(2)), cty.NumberIntVal(int64(1))}, cty.StringVal("PASS")},
		{"equal value", []cty.Value{cty.NumberIntVal(int64(1)), cty.NumberIntVal(int64(1))}, cty.StringVal("PASS")},
		{"not greater value", []cty.Value{cty.NumberIntVal(int64(1)), cty.NumberIntVal(int64(2))}, cty.StringVal(`expected 1 to be greater than or equal to 2`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := assertionFunctionList["assertGreaterThanOrEqualTo"].Impl(tt.args, cty.String); tt.want.NotEqual(got).True() {
				t.Errorf("assertGreaterThanOrEqualTo() = %v, want %v", got.AsString(), tt.want.AsString())
			}
		})
	}
}

func Test_assertLessThan(t *testing.T) {
	tests := []struct {
		name string
		args []cty.Value
		want cty.Value
	}{
		{"greater value", []cty.Value{cty.NumberIntVal(int64(1)), cty.NumberIntVal(int64(2))}, cty.StringVal("PASS")},
		{"not greater value", []cty.Value{cty.NumberIntVal(int64(2)), cty.NumberIntVal(int64(1))}, cty.StringVal(`expected 2 to be less than 1`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := assertionFunctionList["assertLessThan"].Impl(tt.args, cty.String); tt.want.NotEqual(got).True() {
				t.Errorf("assertLessThan() = %v, want %v", got.AsString(), tt.want.AsString())
			}
		})
	}
}

func Test_assertLessThanOrEqualTo(t *testing.T) {
	tests := []struct {
		name string
		args []cty.Value
		want cty.Value
	}{
		{"greater value", []cty.Value{cty.NumberIntVal(int64(1)), cty.NumberIntVal(int64(2))}, cty.StringVal("PASS")},
		{"equal value", []cty.Value{cty.NumberIntVal(int64(1)), cty.NumberIntVal(int64(1))}, cty.StringVal("PASS")},
		{"not greater value", []cty.Value{cty.NumberIntVal(int64(2)), cty.NumberIntVal(int64(1))}, cty.StringVal(`expected 2 to be less than or equal to 1`)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := assertionFunctionList["assertLessThanOrEqualTo"].Impl(tt.args, cty.String); tt.want.NotEqual(got).True() {
				t.Errorf("assertLessThanOrEqualTo() = %v, want %v", got.AsString(), tt.want.AsString())
			}
		})
	}
}
