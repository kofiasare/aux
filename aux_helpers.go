package aux

import (
	"reflect"
)

func isIterable(i interface{}) bool {
	t := argKind(i)
	return t == reflect.Array || t == reflect.Slice || t == reflect.Map
}

func isFunctionWithSingleArg(i interface{}) bool {
	return argKind(i) == reflect.Func &&
		argType(i).NumIn() == 1 &&
		argType(i).In(0).Kind() == reflect.Int
}

func isFunction(i interface{}) bool {
	return argKind(i) == reflect.Func
}

func isInteger(i interface{}) bool {
	return argKind(i) == reflect.Int
}

func isString(i interface{}) bool {
	return argKind(i) == reflect.String
}

func argKind(i interface{}) reflect.Kind {
	return argType(i).Kind()
}

func argVal(i interface{}) reflect.Value {
	return reflect.ValueOf(i)
}

func argType(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}
