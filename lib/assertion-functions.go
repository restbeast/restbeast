package lib

import (
	"encoding/json"
	. "fmt"
	"regexp"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

type AssertionFunc struct {
	Params []function.Parameter
	Type   function.TypeFunc
	Impl   function.ImplFunc
}

func formatRegexAssertionError(format string, arg cty.Value) string {
	return Sprintf("expected %s to be a valid %s", arg.AsString(), format)
}

var assertionFunctionList = map[string]AssertionFunc{
	"assertEqual": {
		Params: []function.Parameter{
			{
				Name:             "a",
				Type:             cty.DynamicPseudoType,
				AllowNull:        true,
				AllowDynamicType: true,
						},
			{
				Name:             "b",
				Type:             cty.DynamicPseudoType,
				AllowNull:        true,
				AllowDynamicType: true,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			rVal := "PASS"

			if args[0].Equals(args[1]).False() {
				argsAasjson, jsonErr := json.MarshalIndent(ctyjson.SimpleJSONValue{args[0]}, "", "  ")
				if jsonErr != nil {
					return cty.StringVal(""), Errorf("Error: failed to variable, \n%s\n", jsonErr)
				}

				argsBasjson, jsonErr := json.MarshalIndent(ctyjson.SimpleJSONValue{args[1]}, "", "  ")
				if jsonErr != nil {
					return cty.StringVal(""), Errorf("Error: failed to variable, \n%s\n", jsonErr)
				}

				rVal = Sprintf("expected: %s\ngot: %s", string(argsAasjson), string(argsBasjson))
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertNotEqual": {
		Params: []function.Parameter{
			{
				Name:             "a",
				Type:             cty.DynamicPseudoType,
				AllowNull:        true,
				AllowDynamicType: true,
						},
			{
				Name:             "b",
				Type:             cty.DynamicPseudoType,
				AllowNull:        true,
				AllowDynamicType: true,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			rVal := "PASS"

			if args[0].NotEqual(args[1]).False() {
				argsAasjson, jsonErr := json.MarshalIndent(ctyjson.SimpleJSONValue{args[0]}, "", "  ")
				if jsonErr != nil {
					return cty.StringVal(""), Errorf("Error: failed to variable, \n%s\n", jsonErr)
				}

				argsBasjson, jsonErr := json.MarshalIndent(ctyjson.SimpleJSONValue{args[1]}, "", "  ")
				if jsonErr != nil {
					return cty.StringVal(""), Errorf("Error: failed to variable, \n%s\n", jsonErr)
				}

				rVal = Sprintf("not expected: %s\ngot: %s", string(argsAasjson), string(argsBasjson))
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertGreaterThan": {
		Params: []function.Parameter{
			{
				Name: "a",
				Type: cty.Number,
						},
			{
				Name: "b",
				Type: cty.Number,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			rVal := "PASS"

			if args[0].GreaterThan(args[1]).False() {
				rVal = Sprintf(
					"expected %s to be greater than %s", args[0].AsBigFloat().String(), args[1].AsBigFloat().String(),
				)
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertGreaterThanOrEqualTo": {
		Params: []function.Parameter{
			{
				Name: "a",
				Type: cty.Number,
						},
			{
				Name: "b",
				Type: cty.Number,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			rVal := "PASS"

			if args[0].GreaterThanOrEqualTo(args[1]).False() {
				rVal = Sprintf(
					"expected %s to be greater than or equal to %s", args[0].AsBigFloat().String(), args[1].AsBigFloat().String(),
				)
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertLessThan": {
		Params: []function.Parameter{
			{
				Name: "a",
				Type: cty.Number,
						},
			{
				Name: "b",
				Type: cty.Number,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			rVal := "PASS"

			if args[0].LessThan(args[1]).False() {
				rVal = Sprintf("expected %s to be less than %s", args[0].AsBigFloat().String(), args[1].AsBigFloat().String())
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertLessThanOrEqualTo": {
		Params: []function.Parameter{
			{
				Name: "a",
				Type: cty.Number,
						},
			{
				Name: "b",
				Type: cty.Number,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			rVal := "PASS"

			if args[0].LessThanOrEqualTo(args[1]).False() {
				rVal = Sprintf(
					"expected %s to be less than or equal to %s", args[0].AsBigFloat().String(), args[1].AsBigFloat().String(),
				)
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertTrue": {
		Params: []function.Parameter{
			{
				Name: "input",
				Type: cty.Bool,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			rVal := "PASS"

			if args[0].False() {
				rVal = Sprintf("expected to be true")
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertFalse": {
		Params: []function.Parameter{
			{
				Name: "input",
				Type: cty.Bool,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			rVal := "PASS"

			if args[0].True() {
				rVal = Sprintf("expected to be false")
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertNonEmptyString": {
		Params: []function.Parameter{
			{
				Name: "value",
				Type: cty.DynamicPseudoType,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			rVal := "PASS"

			if args[0].IsNull() ||
				!args[0].Type().IsPrimitiveType() ||
				args[0].Type().FriendlyName() != "string" ||
				len(args[0].AsString()) == 0 {
				rVal = Sprintf("expected value to be a non empty string")
			}

			return cty.StringVal(rVal), nil
		},
	},
	// Regex based assertions
	"assertEmail": {
		Params: []function.Parameter{
			{
				Name: "email",
				Type: cty.String,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

			rVal := "PASS"

			if !emailRegex.MatchString(args[0].AsString()) {
				rVal = formatRegexAssertionError("email", args[0])
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertUUIDv4": {
		Params: []function.Parameter{
			{
				Name: "UUIDv4",
				Type: cty.String,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var regex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

			rVal := "PASS"

			if !regex.MatchString(args[0].AsString()) {
				rVal = formatRegexAssertionError("UUIDv4", args[0])
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertIpv4": {
		Params: []function.Parameter{
			{
				Name: "ipv4",
				Type: cty.String,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var regex = regexp.MustCompile("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$")

			rVal := "PASS"

			if !regex.MatchString(args[0].AsString()) {
				rVal = formatRegexAssertionError("ipv4", args[0])
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertRegex": {
		Params: []function.Parameter{
			{
				Name: "regex",
				Type: cty.String,
						},
			{
				Name: "input",
				Type: cty.String,
						},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var regex, err = regexp.Compile(args[0].AsString())
			if err != nil {
				return cty.StringVal(""), err
			}

			rVal := "PASS"

			if !regex.MatchString(args[1].AsString()) {
				rVal = formatRegexAssertionError(args[0].AsString(), args[1])
			}

			return cty.StringVal(rVal), nil
		},
	},
}
