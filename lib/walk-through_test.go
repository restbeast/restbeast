package lib

import (
	"encoding/json"
	"github.com/zclconf/go-cty/cty"
	"reflect"
	"testing"
)

func TestWalkThrough(t *testing.T) {
	set := []struct {
		Json []byte
		Cty  cty.Value
	}{
		{[]byte(`"only-string"`), cty.StringVal("only-string")},
		{[]byte(`["test1", "test2"]`), cty.TupleVal([]cty.Value{cty.StringVal("test1"), cty.StringVal("test2")})},
		{
			[]byte(`{ "key1": "value1", "key2": "value2" }`), cty.ObjectVal(
				map[string]cty.Value{
					"key1": cty.StringVal("value1"), "key2": cty.StringVal("value2"),
				},
			),
		},
		{
			[]byte(`{ "key1": ["test1", "test2"] }`), cty.ObjectVal(
				map[string]cty.Value{
					"key1": cty.TupleVal([]cty.Value{cty.StringVal("test1"), cty.StringVal("test2")}),
				},
			),
		},
		{[]byte(`true`), cty.BoolVal(true)},
	}

	for _, run := range set {
		var decoded interface{}
		json.Unmarshal(run.Json, &decoded)

		result := walkThrough(reflect.ValueOf(decoded))
		if !reflect.DeepEqual(result, run.Cty) {
			t.Errorf("error parsing json to cty,\n got %+v,\n want %+v", result, run.Cty)
		}
	}
}
