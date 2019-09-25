package arg

import (
	"reflect"
)

func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// check the reflection value to see if a type is an array or slice
func isSlice(kind reflect.Value) bool {
	return reflect.Slice == kind.Kind() ||
		reflect.Array == kind.Kind()
}
