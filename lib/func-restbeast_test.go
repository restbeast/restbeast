package lib

import (
	"os"
	"testing"
	"time"

	"github.com/zclconf/go-cty/cty"
)

func Test_restbeastReadfileImpl(t *testing.T) {
	argList := []cty.Value{
		cty.StringVal("filename"),
	}

	t.Run("success", func(t *testing.T) {
		got, _ := restbeastReadfileImpl(argList, cty.String)

		if got.AsString() != "###READFILE=filename###" {
			t.Errorf("restbeastReadfileImpl() invalid output")
		}
	})

	t.Run("error", func(t *testing.T) {
		_, err := restbeastReadfileImpl([]cty.Value{}, cty.String)

		if err == nil {
			t.Errorf("restbeastReadfileImpl() should throw an error")
		}
	})
}

func Test_restbeastFillNullImpl(t *testing.T) {

}

func Test_restbeastEnvVarImpl(t *testing.T) {
	_ = os.Setenv("restbeast_var_xxx", "hello")

	t.Run("success", func(t *testing.T) {
		got, _ := restbeastEnvVarImpl([]cty.Value{cty.StringVal("xxx")}, cty.String)

		if got.AsString() != "hello" {
			t.Errorf("restbeastEnvVarImpl() invalid output")
		}
	})
}

func Test_restbeastUnixTimestampImpl(t *testing.T) {
	now := time.Now()

	t.Run("success", func(t *testing.T) {
		_, err := restbeastUnixTimestampImpl([]cty.Value{cty.StringVal("lalala")}, cty.Number)

		if err == nil {
			t.Errorf("restbeastUnixTimestampImpl() invalid output")
		}
	})

	t.Run("error", func(t *testing.T) {
		got, _ := restbeastUnixTimestampImpl([]cty.Value{cty.StringVal(now.Format(time.RFC3339))}, cty.Number)

		asInt, _ := got.AsBigFloat().Int64()

		if asInt != now.Unix() {
			t.Errorf("restbeastUnixTimestampImpl() should throw an error")
		}
	})
}

func Test_restbeastNowImpl(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, _ := restbeastNowImpl([]cty.Value{}, cty.String)

		_, err := time.Parse(time.RFC3339, got.AsString())
		if err != nil {
			t.Errorf("restbeastNowImpl() invalid output")
		}
	})
}
