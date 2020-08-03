package lib

import (
	"fmt"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"os"
	"os/exec"
)

func parseExternalFunctions(internalFunctions map[string]function.Function, externalFunctions []*ExternalFunctionCfg) (allFunctions map[string]function.Function) {
	for _, exFn := range externalFunctions {

		if _, chk := internalFunctions[exFn.Name]; chk {
			fmt.Printf("Error: overwriting an internal function isn't allowed, %s\n", exFn.Name)
			os.Exit(1)
		}

		params := make([]function.Parameter, len(exFn.Args))

		for i, arg := range exFn.Args {
			switch arg {
			case "string":
				params[i] = function.Parameter{
					Name: fmt.Sprintf("arg%s", i),
					Type: cty.String,
				}
			case "list":
				params[i] = function.Parameter{
					Name: fmt.Sprintf("arg%s", i),
					Type: cty.List(cty.DynamicPseudoType),
				}
			case "map":
				params[i] = function.Parameter{
					Name: fmt.Sprintf("arg%s", i),
					Type: cty.Map(cty.DynamicPseudoType),
				}
			}
		}

		internalFunctions[exFn.Name] = function.New(&function.Spec{
			Params: params,
			Type: function.StaticReturnType(cty.String),
			Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
				execArgs := []string{exFn.Script}
				for u, arg := range args {
					switch exFn.Args[u] {
					case "map":
						jsonVal, err := ctyjson.Marshal(arg, cty.Map(cty.DynamicPseudoType))

						if err != nil {
							fmt.Printf("Error: Unable to convert variable to json as map, %s", err)
							os.Exit(1)
						}

						execArgs = append(execArgs, string(jsonVal))
					case "list":
						jsonVal, err := ctyjson.Marshal(arg, cty.List(cty.DynamicPseudoType))

						if err != nil {
							fmt.Printf("Error: Unable to convert variable to json as list, %s", err)
							os.Exit(1)
						}

						execArgs = append(execArgs, string(jsonVal))
					case "string":
						execArgs = append(execArgs, arg.AsString())
					case "number":
						execArgs = append(execArgs, arg.AsString())
					default:
						fmt.Printf("Error: Unknown variable type, %s", exFn.Args[u])
						os.Exit(1)
					}
				}

				data, err := exec.Command(exFn.Interpreter, execArgs...).Output()
				if err != nil {
					return cty.StringVal(""), err
				}

				return cty.StringVal(string(data)), nil
			},
		})
	}

	return internalFunctions
}
