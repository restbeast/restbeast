package lib

import (
	"encoding/json"
	. "fmt"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"io"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

func prepParams(args []string) []function.Parameter {
	params := make([]function.Parameter, len(args))

	for i, arg := range args {
		switch arg {
		case "string":
			params[i] = function.Parameter{
				Name: Sprintf("arg%d", i),
				Type: cty.String,
			}
		case "list":
			params[i] = function.Parameter{
				Name: Sprintf("arg%d", i),
				Type: cty.List(cty.DynamicPseudoType),
			}
		case "map":
			params[i] = function.Parameter{
				Name: Sprintf("arg%d", i),
				Type: cty.Map(cty.DynamicPseudoType),
			}
		}
	}

	return params
}

func prepArgs(exFn *ExternalFunctionCfg, args []cty.Value) (execArgs []string, err error) {
	execArgs = append(execArgs, exFn.Script)

	for u, arg := range args {
		switch exFn.Args[u] {
		case "map":
			newMap := make(map[string]string)
			for k, v := range arg.AsValueMap() {
				newMap[k] = v.AsString()
			}

			marshal, err := json.Marshal(newMap)
			if err != nil {
				return execArgs, Errorf("unable to convert variable to json as map, %s", err)
			}

			execArgs = append(execArgs, string(marshal))
		case "list":
			var list []string
			for i := 0; i < arg.LengthInt(); i++ {
				list = append(list, arg.Index(cty.NumberIntVal(int64(i))).AsString())
			}

			marshal, err := json.Marshal(list)
			if err != nil {
				return execArgs, Errorf("unable to convert variable to json as list, %s", err)
			}

			execArgs = append(execArgs, string(marshal))
		case "string":
			execArgs = append(execArgs, arg.AsString())
		case "number":
			execArgs = append(execArgs, arg.AsBigFloat().String())
		default:
			return execArgs, Errorf("unknown variable type, %s", exFn.Args[u])
		}
	}

	return execArgs, nil
}

func prepImpl(exFn *ExternalFunctionCfg, execCtx *ExecutionContext) function.ImplFunc {
	ctx := *execCtx

	return func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		execArgs, argErr := prepArgs(exFn, args)

		if argErr != nil {
			return cty.Value{}, argErr
		}

		if ctx.Debug {
			log.Printf("executing interpreter %s with file %s", exFn.Interpreter, exFn.Script)
			log.Printf("function %s execution arguments %v", exFn.Name, execArgs[1:])
		}
		cmd := exec.Command(exFn.Interpreter, execArgs...)

		stdout, stdoutErr := cmd.StdoutPipe()
		stderr, stderrErr := cmd.StderrPipe()

		if stdoutErr != nil {
			return cty.Value{}, Errorf("couldn't get stdout %s, %s", exFn.Name, stdoutErr)
		}

		if stderrErr != nil {
			return cty.Value{}, Errorf("couldn't get stderr %s, %s", exFn.Name, stdoutErr)
		}

		if err := cmd.Start(); err != nil {
			return cty.Value{}, Errorf("couldn't start command %s, %s", exFn.Name, err)
		}

		stdOutBuffer := new(strings.Builder)
		if _, ioErr := io.Copy(stdOutBuffer, stdout); ioErr != nil {
			return cty.Value{}, Errorf("io %s, %s", exFn.Name, ioErr)
		}

		stdErrBuffer := new(strings.Builder)
		if _, ioErr := io.Copy(stdErrBuffer, stderr); ioErr != nil {
			return cty.Value{}, Errorf("io %s, %s", exFn.Name, ioErr)
		}

		if err := cmd.Wait(); err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					if ctx.Debug {
						log.Printf("external function stderr %s", stdErrBuffer.String())
					}

					return cty.Value{}, Errorf("external function %s exited with status: %d", exFn.Name, status.ExitStatus())
				}
			} else {
				return cty.Value{}, Errorf("external function %s error %s, %s", exFn.Name, err, exiterr)
			}
		}

		functionOutput := stdOutBuffer.String()
		if ctx.Debug {
			log.Printf("function %s output %s", exFn.Name, functionOutput)
		}

		return cty.StringVal(functionOutput), nil
	}
}

func parseExternalFunctions(internalFunctions map[string]function.Function, externalFunctions []*ExternalFunctionCfg, execCtx *ExecutionContext) (*map[string]function.Function, error) {
	for _, exFn := range externalFunctions {
		if _, chk := internalFunctions[exFn.Name]; chk {
			return nil, Errorf("Error: overwriting an internal function isn't allowed, %s\n", exFn.Name)
		}

		internalFunctions[exFn.Name] = function.New(&function.Spec{
			Params: prepParams(exFn.Args),
			Type:   function.StaticReturnType(cty.String),
			Impl:   prepImpl(exFn, execCtx),
		})
	}

	return &internalFunctions, nil
}
