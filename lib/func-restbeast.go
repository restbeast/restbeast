package lib

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-errors/errors"
	"github.com/zclconf/go-cty/cty"
)

func restbeastReadFileImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	if len(args) < 1 {
		return cty.StringVal(""), errors.New("Invalid argument count")
	}

	return cty.StringVal(fmt.Sprintf("###READFILE=%s###", args[0].AsString())), nil
}
func restbeastFilePartImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	if len(args) < 3 {
		return cty.StringVal(""), errors.New("Invalid argument count")
	}

	if cty.Number != args[1].Type() {
		return cty.StringVal(""), errors.New("Invalid argument type, offset expected to be a number")
	}

	if cty.Number != args[2].Type() {
		return cty.StringVal(""), errors.New("Invalid argument type, length expected to be a number")
	}

	offset, _ := args[1].AsBigFloat().Int64()
	length, _ := args[2].AsBigFloat().Int64()

	return cty.StringVal(fmt.Sprintf("###READFILE=%s:%d:%d###", args[0].AsString(), offset, length)), nil
}

func restbeastFillNullImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	r := rand.Float32()
	p, _ := args[0].AsBigFloat().Float32()

	if r > p/100 {
		return args[1], nil
	} else {
		return cty.NilVal, nil
	}
}

func restbeastEnvVarImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	varKey := args[0].AsString()
	value := os.Getenv(fmt.Sprintf("restbeast_var_%s", varKey))
	if value != "" {
		return cty.StringVal(value), nil
	}

	value = os.Getenv(varKey)
	if value != "" {
		return cty.StringVal(value), nil
	}

	return cty.StringVal(""), errors.New(fmt.Sprintf("Environment variable %s not found", varKey))
}

func restbeastEnvVarWithDefaultImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	varKey := args[0].AsString()
	value := os.Getenv(fmt.Sprintf("restbeast_var_%s", varKey))
	if value != "" {
		return cty.StringVal(value), nil
	}

	value = os.Getenv(varKey)
	if value != "" {
		return cty.StringVal(value), nil
	}

	return args[1], nil
}

func restbeastUnixTimestampImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	varDate := args[0].AsString()
	t, err := time.Parse(time.RFC3339, varDate)

	if err != nil {
		return cty.NilVal, err
	}

	return cty.NumberIntVal(t.Unix()), nil
}

func restbeastNowImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	return cty.StringVal(time.Now().UTC().Format(time.RFC3339)), nil
}
