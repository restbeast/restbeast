package lib

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-errors/errors"
	"github.com/zclconf/go-cty/cty"
)

func restbeastReadfileImpl(args []cty.Value, retType cty.Type) (cty.Value, error) {
	if len(args) < 1 {
		return cty.StringVal(""), errors.New("Invalid argument count")
	}

	return cty.StringVal(fmt.Sprintf("###READFILE=%s###", args[0].AsString())), nil
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
	return cty.StringVal(value), nil
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