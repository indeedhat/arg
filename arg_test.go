package arg

import (
	"reflect"
	"testing"
)

func t_setupExpectedArg(v *interface{}, vals []string) *expectedArg {
	return &expectedArg{
		Values:     vals,
		ScalarType: reflect.ValueOf(v),
	}
}

func TestInflateString(t *testing.T) {
	s := ""
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"stringval"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateString generated error: %s", err)
	}

	if s != "stringval" {
		t.Error("Failed to inflate string")
	}
}

func TestInflateStringSlice(t *testing.T) {
	var s []string
	expected := []string{"string1", "string2"}
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     expected,
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateString generated error: %s", err)
	}

	if !reflect.DeepEqual(expected, s) {
		t.Error("Failed to inflate string slice")
	}
}

func TestInflateInt(t *testing.T) {
	var s int
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if s != int(1) {
		t.Error("Failed to inflate int")
	}
}

func TestInflateIntSlice(t *testing.T) {
	var s []int16
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1", "2", "3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if !reflect.DeepEqual([]int16{1, 2, 3}, s) {
		t.Error("Failed to inflate int slice")
	}
}

func TestInflateUint(t *testing.T) {
	var s uint
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if s != uint(1) {
		t.Error("Failed to inflate uint")
	}
}

func TestInflateUintSlice(t *testing.T) {
	var s []uint8
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1", "2", "3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if !reflect.DeepEqual([]uint8{1, 2, 3}, s) {
		t.Error("Failed to inflate uint8 slice")
	}
}

func TestInflateFloat(t *testing.T) {
	var s float32
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1.1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateFloat generated error: %s", err)
	}

	if s != float32(1.1) {
		t.Error("Failed to inflate float")
	}
}

func TestInflateFloatSlice(t *testing.T) {
	var s []float64
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1.1", "2.2", "3.3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateFloat generated error: %s", err)
	}

	if !reflect.DeepEqual([]float64{1.1, 2.2, 3.3}, s) {
		t.Error("Failed to inflate float slice")
	}
}

func TestInflateBool(t *testing.T) {
	var s bool
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"true"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateBool generated error: %s", err)
	}

	if s != true {
		t.Error("Failed to inflate bool")
	}
}

func TestInflateBoolSloce(t *testing.T) {
	var s []bool
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"true", "false", "true"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateBool generated error: %s", err)
	}

	if !reflect.DeepEqual([]bool{true, false, true}, s) {
		t.Error("Failed to inflate bool slice")
	}
}
