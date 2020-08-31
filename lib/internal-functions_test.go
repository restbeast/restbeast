package lib

import (
	"fmt"
	"testing"
)

func Test_getCtyFunctions(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got := getCtyFunctions()
		wantLen := len(defaultFunctions) + len(gofakeitFunctionList)

		fmt.Println(len(defaultFunctions), len(gofakeitFunctionList))
		fmt.Println(len(got))

		if len(got) != wantLen {
			t.Error("function builder failed")
		}
	})
}
