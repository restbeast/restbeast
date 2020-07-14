package lib

import (
  "github.com/zclconf/go-cty/cty"
  "reflect"
)

func walkThrough(v reflect.Value) cty.Value {
  for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
    v = v.Elem()
  }

  switch v.Kind() {
    case reflect.Array, reflect.Slice:
      length := v.Len()

      newSlice := make([]cty.Value, length)

      for i := 0; i < length; i++ {
      newSlice[i] = walkThrough(v.Index(i))
    }

      return cty.ListVal(newSlice)
    case reflect.Map:
      newMap := make(map[string]cty.Value)

      for _, k := range v.MapKeys() {
      newMap[k.String()] = walkThrough(v.MapIndex(k))
    }

      return cty.ObjectVal(newMap)
    case reflect.Bool:
      return cty.BoolVal(v.Bool())
    default:
      return cty.StringVal(v.String())
  }
}
