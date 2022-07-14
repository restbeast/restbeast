package lib

import (
	"reflect"

	"github.com/zclconf/go-cty/cty"
)

func walkThrough(v reflect.Value) cty.Value {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		length := v.Len()

		// Return empty value for empty slice
		if length == 0 {
			return cty.Value{}
		}

		newSlice := make([]cty.Value, length)

		for i := 0; i < length; i++ {
			newSlice[i] = walkThrough(v.Index(i))
		}

		return cty.TupleVal(newSlice)
	case reflect.Map:
		newMap := make(map[string]cty.Value)

		for _, k := range v.MapKeys() {
			newMap[k.String()] = walkThrough(v.MapIndex(k))
		}

		return cty.ObjectVal(newMap)
	case reflect.Bool:
		return cty.BoolVal(v.Bool())
	case reflect.Int:
		return cty.NumberIntVal(v.Int())
	case reflect.Float64:
		return cty.NumberFloatVal(v.Float())
	case reflect.Float32:
		return cty.NumberFloatVal(v.Float())
	case reflect.String:
	default:
		return cty.StringVal(v.String())
	}

	return cty.StringVal(v.String())
}
