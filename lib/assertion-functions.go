package lib

import (
	"encoding/json"
	. "fmt"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"regexp"
)

type AssertionFunc struct {
	Params []function.Parameter
	Type   function.TypeFunc
	Impl   function.ImplFunc
}

var assertionFunctionList = map[string]AssertionFunc{
	"assertEmail": {
		Params: []function.Parameter{
			function.Parameter{
				Name: "email",
				Type: cty.String,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

			rVal := "PASS"

			if !emailRegex.MatchString(args[0].AsString()) {
				rVal = Sprintf("want: assertEmail()\ngot: %s", args[0].AsString())
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertEqual": {
		Params: []function.Parameter{
			function.Parameter{
				Name:             "a",
				Type:             cty.DynamicPseudoType,
				AllowNull:        true,
				AllowDynamicType: true,
			},
			function.Parameter{
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

				rVal = Sprintf("want: %s\ngot: %s", string(argsAasjson), string(argsBasjson))
			}

			return cty.StringVal(rVal), nil
		},
	},
	"assertNotEqual": {
		Params: []function.Parameter{
			function.Parameter{
				Name:             "a",
				Type:             cty.DynamicPseudoType,
				AllowNull:        true,
				AllowDynamicType: true,
			},
			function.Parameter{
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

				rVal = Sprintf("not want: %s\ngot: %s", string(argsAasjson), string(argsBasjson))
			}

			return cty.StringVal(rVal), nil
		},
	},
}
