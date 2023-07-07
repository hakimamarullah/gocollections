package utils

import "reflect"

func IsSameKind[T any](a, b T) (bool, reflect.Type) {
	return reflect.TypeOf(a) == reflect.TypeOf(b), reflect.TypeOf(a)
}
