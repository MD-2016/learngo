package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkVal := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkVal(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkVal(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkVal(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkVal(v)
		}
	case reflect.Func:
		valRes := val.Call(nil)
		for _, res := range valRes {
			walkVal(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
