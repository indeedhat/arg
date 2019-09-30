package arg

import (
	"fmt"
	"reflect"
	"testing"
)

var t_testData []string = []string{"--thing1=thing1", "-t=thing2", "position", "-tgroup", "val", "val=other", "--val=dthin=2"}

func TestParserTracksRawArgs(t *testing.T) {
	p := NewParser()
	if nil != p.Parse(t_testData) {
		t.Error("Parse returned an error")
	}

	if !reflect.DeepEqual(t_testData, p.RawArgs) {
		t.Error("Did not keep track of raw args")
	}
}

func TestParserExtras(t *testing.T) {
	expected := []Argument{
		{T_NAMED, "thing1", "thing1", 0},
		{T_DASHED, "true", "t", 0},
		{T_POSITIONAL, "thing2", "", 0},
		{T_POSITIONAL, "position", "", 1},
		{T_DASHED, "true", "t", 0},
		{T_DASHED, "true", "g", 0},
		{T_DASHED, "true", "r", 0},
		{T_DASHED, "true", "o", 0},
		{T_DASHED, "true", "u", 0},
		{T_DASHED, "true", "p", 0},
		{T_POSITIONAL, "val", "", 2},
		{T_POSITIONAL, "val=other", "", 3},
		{T_NAMED, "dthin=2", "val", 0},
	}

	p := NewParser()
	if nil != p.Parse(t_testData) {
		t.Error("Parse returned an error")
	}

	if !reflect.DeepEqual(expected, p.extras) {
		t.Error("Did not return the extra args it should have")
		fmt.Printf("%+v\n%+v", expected, p.extras)
	}
}
