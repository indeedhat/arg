package arg

import (
	"reflect"
)

func kind(field reflect.Value) reflect.Kind {
	if isSlice(field) {
		return field.Type().Elem().Kind()
	}

	return field.Type().Kind()
}
