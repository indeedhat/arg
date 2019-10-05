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
		t.Fatalf("failed to set indexed string")
	}

	if 43938 != named {
		t.Fatalf("failed to set indexed string")
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
		t.Fatalf("failed to set indexed []string: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []string: %v", named)
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
		t.Fatalf("failed to set indexed string")
	}

	if 43938 != named {
		t.Fatalf("failed to set indexed string")
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
		t.Fatalf("failed to set indexed []string: %v", indexed)
	}

	if !reflect.DeepEqual(nExpected, named) {
		t.Fatalf("failed to set indexed []string: %v", named)
	}
}
