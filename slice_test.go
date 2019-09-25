package arg

import (
	"reflect"
	"testing"
)

func TestSliceEqualsValid(t *testing.T) {
	a1 := []string{"one", "two", "three"}
	a2 := []string{"one", "two", "three"}

	if !sliceEqual(a1, a2) {
		t.Error("Reported matching slices as not equal")
	}
}

func TestSliceEqualsInvalid(t *testing.T) {
	a1 := []string{"one", "two", "three"}
	a2 := []string{"one", "four", "three"}

	if sliceEqual(a1, a2) {
		t.Error("Reported different slices as equal")
	}
}

func TestIsSlice(t *testing.T) {
	var i int
	val := reflect.ValueOf(i)
	if isSlice(val) {
		t.Error("Detected int type as slice")
	}

	is := []int{1}
	val = reflect.ValueOf(is)
	if !isSlice(val) {
		t.Error("Did not detect array as slice (it should)")
	}

	val = reflect.ValueOf(is[:])
	if !isSlice(val) {
		t.Error("Did not detect slice")
	}
}
