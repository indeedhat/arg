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
