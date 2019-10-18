package arg

import (
	"reflect"
	"testing"
)

func TestApiString(t *testing.T) {
	args := []string{
		"somestring",
		"--key=some-string-val",
	}

	var indexed string
	var named string

	parser := NewParser()

	parser.String(&indexed, "", &ArgConfig{})
	parser.String(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if "somestring" != indexed {
		t.Fatalf("failed to set indexed string")
	}

	if "some-string-val" != named {
		t.Fatalf("failed to set indexed string")
	}
}

func TestApiStringSlice(t *testing.T) {
	args := []string{
		"somestring",
		"--key=some-string-val",
		"someotherstring",
		"--key=some-string-v2",
	}

	iExpected := []string{
		"somestring",
		"someotherstring",
	}

	nExpected := []string{
		"some-string-val",
		"some-string-v2",
	}

	var indexed []string
	var named []string

	parser := NewParser()

	parser.Strings(&indexed, "", &ArgConfig{})
	parser.Strings(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []string: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []string: %v", named)
	}
}

func TestApiInt(t *testing.T) {
	args := []string{
		"1004",
		"--key=43938",
	}

	var indexed int
	var named int

	parser := NewParser()

	parser.Int(&indexed, "", &ArgConfig{})
	parser.Int(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 1004 != indexed {
		t.Fatalf("failed to set indexed int")
	}

	if 43938 != named {
		t.Fatalf("failed to set indexed int")
	}
}

func TestApiIntSlice(t *testing.T) {
	args := []string{
		"1004",
		"--key=43938",
		"1804",
		"--key=83938",
	}

	iExpected := []int{
		1004,
		1804,
	}

	nExpected := []int{
		43938,
		83938,
	}

	var indexed []int
	var named []int

	parser := NewParser()

	parser.Ints(&indexed, "", &ArgConfig{})
	parser.Ints(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []int: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []int: %v", named)
	}
}

func TestApiInt8(t *testing.T) {
	args := []string{
		"120",
		"--key=19",
	}

	var indexed int8
	var named int8

	parser := NewParser()

	parser.Int8(&indexed, "", &ArgConfig{})
	parser.Int8(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 120 != indexed {
		t.Fatalf("failed to set indexed int8")
	}

	if 19 != named {
		t.Fatalf("failed to set indexed int8")
	}
}

func TestApiInt8Slice(t *testing.T) {
	args := []string{
		"1",
		"--key=127",
		"12",
		"--key=94",
	}

	iExpected := []int8{
		1,
		12,
	}

	nExpected := []int8{
		127,
		94,
	}

	var indexed []int8
	var named []int8

	parser := NewParser()

	parser.Int8s(&indexed, "", &ArgConfig{})
	parser.Int8s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []int8: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []int8: %v", named)
	}
}

func TestApiBadInt8Slice(t *testing.T) {
	args := []string{
		"1",
		"--key=666",
		"12",
		"--key=94",
	}

	iExpected := []int8{
		1,
		12,
	}

	nExpected := []int8{
		127,
		94,
	}

	var indexed []int8
	var named []int8

	parser := NewParser()

	parser.Int8s(&indexed, "", &ArgConfig{})
	parser.Int8s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []int8: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []int8: %v", named)
	}
}

func TestApiInt16(t *testing.T) {
	args := []string{
		"244",
		"--key=19",
	}

	var indexed int16
	var named int16

	parser := NewParser()

	parser.Int16(&indexed, "", &ArgConfig{})
	parser.Int16(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 244 != indexed {
		t.Fatalf("failed to set indexed int16")
	}

	if 19 != named {
		t.Fatalf("failed to set indexed int16")
	}
}

func TestApiInt16Slice(t *testing.T) {
	args := []string{
		"1",
		"--key=255",
		"244",
		"--key=94",
	}

	iExpected := []int16{
		1,
		244,
	}

	nExpected := []int16{
		255,
		94,
	}

	var indexed []int16
	var named []int16

	parser := NewParser()

	parser.Int16s(&indexed, "", &ArgConfig{})
	parser.Int16s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []int16: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []int16: %v", named)
	}
}

func TestApiInt32(t *testing.T) {
	args := []string{
		"244",
		"--key=19",
	}

	var indexed int32
	var named int32

	parser := NewParser()

	parser.Int32(&indexed, "", &ArgConfig{})
	parser.Int32(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 244 != indexed {
		t.Fatalf("failed to set indexed int32")
	}

	if 19 != named {
		t.Fatalf("failed to set indexed int32")
	}
}

func TestApiInt32Slice(t *testing.T) {
	args := []string{
		"1",
		"--key=255",
		"244",
		"--key=94",
	}

	iExpected := []int32{
		1,
		244,
	}

	nExpected := []int32{
		255,
		94,
	}

	var indexed []int32
	var named []int32

	parser := NewParser()

	parser.Int32s(&indexed, "", &ArgConfig{})
	parser.Int32s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []int32: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []int32: %v", named)
	}
}

func TestApiInt64(t *testing.T) {
	args := []string{
		"244",
		"--key=19",
	}

	var indexed int64
	var named int64

	parser := NewParser()

	parser.Int64(&indexed, "", &ArgConfig{})
	parser.Int64(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 244 != indexed {
		t.Fatalf("failed to set indexed int32")
	}

	if 19 != named {
		t.Fatalf("failed to set indexed int32")
	}
}

func TestApiInt64Slice(t *testing.T) {
	args := []string{
		"1",
		"--key=255",
		"244",
		"--key=94",
	}

	iExpected := []int64{
		1,
		244,
	}

	nExpected := []int64{
		255,
		94,
	}

	var indexed []int64
	var named []int64

	parser := NewParser()

	parser.Int64s(&indexed, "", &ArgConfig{})
	parser.Int64s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []int32: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []int32: %v", named)
	}
}

func TestApiUint(t *testing.T) {
	args := []string{
		"1004",
		"--key=43938",
	}

	var indexed uint
	var named uint

	parser := NewParser()

	parser.Uint(&indexed, "", &ArgConfig{})
	parser.Uint(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 1004 != indexed {
		t.Fatalf("failed to set indexed uint")
	}

	if 43938 != named {
		t.Fatalf("failed to set indexed uint")
	}
}

func TestApiUintSlice(t *testing.T) {
	args := []string{
		"1004",
		"--key=43938",
		"1804",
		"--key=83938",
	}

	iExpected := []uint{
		1004,
		1804,
	}

	nExpected := []uint{
		43938,
		83938,
	}

	var indexed []uint
	var named []uint

	parser := NewParser()

	parser.Uints(&indexed, "", &ArgConfig{})
	parser.Uints(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []uint: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []uint: %v", named)
	}
}

func TestApiUint8(t *testing.T) {
	args := []string{
		"120",
		"--key=19",
	}

	var indexed uint8
	var named uint8

	parser := NewParser()

	parser.Uint8(&indexed, "", &ArgConfig{})
	parser.Uint8(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 120 != indexed {
		t.Fatalf("failed to set indexed uint8 %d", indexed)
	}

	if 19 != named {
		t.Fatalf("failed to set indexed uint8")
	}
}

func TestApiUint8Slice(t *testing.T) {
	args := []string{
		"1",
		"--key=127",
		"12",
		"--key=94",
	}

	iExpected := []uint8{
		1,
		12,
	}

	nExpected := []uint8{
		127,
		94,
	}

	var indexed []uint8
	var named []uint8

	parser := NewParser()

	parser.Uint8s(&indexed, "", &ArgConfig{})
	parser.Uint8s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []uint8: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []uint8: %v", named)
	}
}

func TestApiUint16(t *testing.T) {
	args := []string{
		"244",
		"--key=19",
	}

	var indexed uint16
	var named uint16

	parser := NewParser()

	parser.Uint16(&indexed, "", &ArgConfig{})
	parser.Uint16(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 244 != indexed {
		t.Fatalf("failed to set indexed uint16")
	}

	if 19 != named {
		t.Fatalf("failed to set indexed uint16")
	}
}

func TestApiUint16Slice(t *testing.T) {
	args := []string{
		"1",
		"--key=255",
		"244",
		"--key=94",
	}

	iExpected := []uint16{
		1,
		244,
	}

	nExpected := []uint16{
		255,
		94,
	}

	var indexed []uint16
	var named []uint16

	parser := NewParser()

	parser.Uint16s(&indexed, "", &ArgConfig{})
	parser.Uint16s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []uint16: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []uint16: %v", named)
	}
}

func TestApiUint32(t *testing.T) {
	args := []string{
		"244",
		"--key=19",
	}

	var indexed uint32
	var named uint32

	parser := NewParser()

	parser.Uint32(&indexed, "", &ArgConfig{})
	parser.Uint32(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 244 != indexed {
		t.Fatalf("failed to set indexed uint32")
	}

	if 19 != named {
		t.Fatalf("failed to set indexed uint32")
	}
}

func TestApiUint32Slice(t *testing.T) {
	args := []string{
		"1",
		"--key=255",
		"244",
		"--key=94",
	}

	iExpected := []uint32{
		1,
		244,
	}

	nExpected := []uint32{
		255,
		94,
	}

	var indexed []uint32
	var named []uint32

	parser := NewParser()

	parser.Uint32s(&indexed, "", &ArgConfig{})
	parser.Uint32s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []uint32: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []uint32: %v", named)
	}
}

func TestApiUint64(t *testing.T) {
	args := []string{
		"244",
		"--key=19",
	}

	var indexed uint64
	var named uint64

	parser := NewParser()

	parser.Uint64(&indexed, "", &ArgConfig{})
	parser.Uint64(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 244 != indexed {
		t.Fatalf("failed to set indexed uint32")
	}

	if 19 != named {
		t.Fatalf("failed to set indexed uint32")
	}
}

func TestApiUint64Slice(t *testing.T) {
	args := []string{
		"1",
		"--key=255",
		"244",
		"--key=94",
	}

	iExpected := []uint64{
		1,
		244,
	}

	nExpected := []uint64{
		255,
		94,
	}

	var indexed []uint64
	var named []uint64

	parser := NewParser()

	parser.Uint64s(&indexed, "", &ArgConfig{})
	parser.Uint64s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []uint32: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []uint32: %v", named)
	}
}

func TestApiFloat32(t *testing.T) {
	args := []string{
		"1004",
		"--key=43938",
	}

	var indexed float32
	var named float32

	parser := NewParser()

	parser.Float32(&indexed, "", &ArgConfig{})
	parser.Float32(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 1004 != indexed {
		t.Fatalf("failed to set indexed float32")
	}

	if 43938 != named {
		t.Fatalf("failed to set indexed float32")
	}
}

func TestApiFloat32Slice(t *testing.T) {
	args := []string{
		"1004",
		"--key=43938",
		"1804",
		"--key=83938",
	}

	iExpected := []float32{
		1004,
		1804,
	}

	nExpected := []float32{
		43938,
		83938,
	}

	var indexed []float32
	var named []float32

	parser := NewParser()

	parser.Float32s(&indexed, "", &ArgConfig{})
	parser.Float32s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []float32: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []float32: %v", named)
	}
}

func TestApiFloat64(t *testing.T) {
	args := []string{
		"1004",
		"--key=43938",
	}

	var indexed float64
	var named float64

	parser := NewParser()

	parser.Float64(&indexed, "", &ArgConfig{})
	parser.Float64(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if 1004 != indexed {
		t.Fatalf("failed to set indexed float64")
	}

	if 43938 != named {
		t.Fatalf("failed to set indexed float64")
	}
}

func TestApiFloat64Slice(t *testing.T) {
	args := []string{
		"1004",
		"--key=43938",
		"1804",
		"--key=83938",
	}

	iExpected := []float64{
		1004,
		1804,
	}

	nExpected := []float64{
		43938,
		83938,
	}

	var indexed []float64
	var named []float64

	parser := NewParser()

	parser.Float64s(&indexed, "", &ArgConfig{})
	parser.Float64s(&named, "key", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(iExpected, indexed) {
		t.Fatalf("failed to set indexed []float64: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []float64: %v", named)
	}
}

func TestApiBool(t *testing.T) {
	args := []string{
		"true",
		"--key",
		"-k",
		"-j",
		"false",
		"--key2=false",
	}

	var indexed bool
	var key bool
	var k bool
	var j bool
	var key2 bool

	parser := NewParser()

	parser.Bool(&indexed, "", &ArgConfig{})
	parser.Bool(&key, "key", &ArgConfig{})
	parser.Bool(&k, "k", &ArgConfig{})
	parser.Bool(&j, "j", &ArgConfig{})
	parser.Bool(&key2, "key2", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !indexed {
		t.Fatalf("failed to set indexed bool")
	}

	if !key {
		t.Fatalf("failed to set named bool")
	}

	if !k {
		t.Fatalf("failed to set flagged bool")
	}

	if j {
		t.Fatalf("failed to set flagged bool to false")
	}

	if key2 {
		t.Fatalf("failed to set named bool to false")
	}
}

func TestApiBoolSlice(t *testing.T) {
	args := []string{
		"true",
		"--key",
		"-k",
		"true",
		"false",
		"-k=false",
		"--key=false",
	}

	expected := []bool{
		true,
		false,
	}
	var indexed []bool
	var key []bool
	var k []bool

	parser := NewParser()

	parser.Bools(&indexed, "", &ArgConfig{})
	parser.Bools(&key, "key", &ArgConfig{})
	parser.Bools(&k, "k", &ArgConfig{})

	if err := parser.Parse(args); nil != err {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, indexed) {
		t.Fatalf("failed to set indexed []bool")
	}

	if !reflect.DeepEqual(expected, key) {
		t.Fatalf("failed to set named []bool")
	}

	if !reflect.DeepEqual(expected, k) {
		t.Fatalf("failed to set flagged []bool")
	}
}
