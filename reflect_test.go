package arg

import (
	"reflect"
	"testing"
)

func TestKindOnScalar(t *testing.T) {
	var i int
	val := reflect.ValueOf(i)
	if reflect.Int != kind(val) {
		t.Error("Cannot detect int scalar type")
	}

	var s string
	val = reflect.ValueOf(s)
	if reflect.String != kind(val) {
		t.Error("Cannot detect string scalar type")
	}

	var f float32
	val = reflect.ValueOf(f)
	if reflect.Float32 != kind(val) {
		t.Error("Cannot detect float32 scalar type")
	}

	var b bool
	val = reflect.ValueOf(b)
	if reflect.Bool != kind(val) {
		t.Error("Cannot detect bool scalar type")
	}

}

func TestKindOnSlice(t *testing.T) {
	i := [1]int{1}
	val := reflect.ValueOf(i[:])
	if reflect.Int != kind(val) {
		t.Error("Cannot detect int slice")
	}

	s := [1]string{"1"}
	val = reflect.ValueOf(s[:])
	if reflect.String != kind(val) {
		t.Error("Cannot detect string slice")
	}

	f := [1]float32{1.1}
	val = reflect.ValueOf(f[:])
	if reflect.Float32 != kind(val) {
		t.Error("Cannot detect float32 slice")
	}

	b := [1]bool{true}
	val = reflect.ValueOf(b[:])
	if reflect.Bool != kind(val) {
		t.Error("Cannot detect bool slice")
	}

}

func TestKindOnArray(t *testing.T) {
	i := [1]int{1}
	val := reflect.ValueOf(i)
	if reflect.Int != kind(val) {
		t.Error("Cannot detect int array")
	}

	s := [1]string{"1"}
	val = reflect.ValueOf(s)
	if reflect.String != kind(val) {
		t.Error("Cannot detect string array")
	}

	f := [1]float32{1.1}
	val = reflect.ValueOf(f)
	if reflect.Float32 != kind(val) {
		t.Error("Cannot detect float32 array")
	}

	b := [1]bool{true}
	val = reflect.ValueOf(b)
	if reflect.Bool != kind(val) {
		t.Error("Cannot detect bool array")
	}

}
