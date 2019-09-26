package arg

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestParserTracksRawArgs(t *testing.T) {
	test := []string{"--thing1=thing1", "-t=thing2", "position", "-tgroup", "val", "val=other", "--val=dthin=2"}

	p := NewParser()
	if nil != p.Parse(test) {
		t.Error("Parse returned an error")
	}

	if !reflect.DeepEqual(test, p.RawArgs) {
		t.Error("Did not keep track of raw args")
	}

	b, err := json.Marshal(p.extras)
	if nil != err {
		t.Error(err)
	}

	fmt.Println(string(b))
}
