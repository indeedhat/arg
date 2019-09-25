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

func TestInflateInt8(t *testing.T) {
	var s int8
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if s != int8(1) {
		t.Error("Failed to inflate int8")
	}
}

func TestInflateInt8Slice(t *testing.T) {
	var s []int8
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1", "2", "3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if !reflect.DeepEqual([]int8{1, 2, 3}, s) {
		t.Error("Failed to inflate int8 slice")
	}
}

func TestInflateInt16(t *testing.T) {
	var s int16
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if s != int16(1) {
		t.Error("Failed to inflate int16")
	}
}

func TestInflateInt16Slice(t *testing.T) {
	var s []int16
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1", "2", "3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if !reflect.DeepEqual([]int16{1, 2, 3}, s) {
		t.Error("Failed to inflate int16 slice")
	}
}

func TestInflateInt32(t *testing.T) {
	var s int32
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if s != int32(1) {
		t.Error("Failed to inflate int32")
	}
}

func TestInflateInt32Slice(t *testing.T) {
	var s []int32
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1", "2", "3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if !reflect.DeepEqual([]int32{1, 2, 3}, s) {
		t.Error("Failed to inflate int32 slice")
	}
}

func TestInflateInt64(t *testing.T) {
	var s int64
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if s != int64(1) {
		t.Error("Failed to inflate int64")
	}
}

func TestInflateInt64Slice(t *testing.T) {
	var s []int64
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1", "2", "3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateInt generated error: %s", err)
	}

	if !reflect.DeepEqual([]int64{1, 2, 3}, s) {
		t.Error("Failed to inflate int64 slice")
	}
}

func TestInflateUint8(t *testing.T) {
	var s uint8
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateUint generated error: %s", err)
	}

	if s != uint8(1) {
		t.Error("Failed to inflate uint8")
	}
}

func TestInflateUint8Slice(t *testing.T) {
	var s []uint8
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1", "2", "3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateUint generated error: %s", err)
	}

	if !reflect.DeepEqual([]uint8{1, 2, 3}, s) {
		t.Error("Failed to inflate uint8 slice")
	}
}

func TestInflateUint16(t *testing.T) {
	var s uint16
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateUint generated error: %s", err)
	}

	if s != uint16(1) {
		t.Error("Failed to inflate uint16")
	}
}

func TestInflateUint16Slice(t *testing.T) {
	var s []uint16
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1", "2", "3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateUint generated error: %s", err)
	}

	if !reflect.DeepEqual([]uint16{1, 2, 3}, s) {
		t.Error("Failed to inflate uint16 slice")
	}
}

func TestInflateUint32(t *testing.T) {
	var s uint32
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateUint generated error: %s", err)
	}

	if s != uint32(1) {
		t.Error("Failed to inflate uint32")
	}
}

func TestInflateUint32Slice(t *testing.T) {
	var s []uint32
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1", "2", "3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateUint generated error: %s", err)
	}

	if !reflect.DeepEqual([]uint32{1, 2, 3}, s) {
		t.Error("Failed to inflate uint32 slice")
	}
}

func TestInflateUint64(t *testing.T) {
	var s uint64
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateUint generated error: %s", err)
	}

	if s != uint64(1) {
		t.Error("Failed to inflate uint64")
	}
}

func TestInflateUint64Slice(t *testing.T) {
	var s []uint64
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1", "2", "3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateUint generated error: %s", err)
	}

	if !reflect.DeepEqual([]uint64{1, 2, 3}, s) {
		t.Error("Failed to inflate uint64 slice")
	}
}

func TestInflateFloat32(t *testing.T) {
	var s float32
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1.1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateFloat generated error: %s", err)
	}

	if s != float32(1.1) {
		t.Error("Failed to inflate float32")
	}
}

func TestInflateFloat32Slice(t *testing.T) {
	var s []float32
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1.1", "2.2", "3.3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateFloat generated error: %s", err)
	}

	if !reflect.DeepEqual([]float32{1.1, 2.2, 3.3}, s) {
		t.Error("Failed to inflate float32 slice")
	}
}

func TestInflateFloat64(t *testing.T) {
	var s float64
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1.1"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateFloat generated error: %s", err)
	}

	if s != float64(1.1) {
		t.Error("Failed to inflate float64")
	}
}

func TestInflateFloat64Slice(t *testing.T) {
	var s []float64
	ex := &expectedArg{
		ScalarType: reflect.ValueOf(&s).Elem(),
		Values:     []string{"1.1", "2.2", "3.3"},
	}

	if err := ex.inflate(); nil != err {
		t.Errorf("inflateFloat generated error: %s", err)
	}

	if !reflect.DeepEqual([]float64{1.1, 2.2, 3.3}, s) {
		t.Error("Failed to inflate float64 slice")
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
