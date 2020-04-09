package lib

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

func funcGenerateUuid() function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{},
		Type:   function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			return cty.StringVal(uuid.NewV4().String()), nil
		},
	})
}
