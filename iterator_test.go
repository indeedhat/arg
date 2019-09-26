package arg

import (
	"strconv"
	"testing"
)

func TestStringIterator(t *testing.T) {
	itr := newStringIterator([]string{
		"0",
		"1",
		"2",
		"3",
		"4",
	})

	if "" != itr.Value() {
		t.Error("iterator is returning a value before the first call to next")
	}

	i := 0
	for itr.Next() {
		if i < 4 {
			if val, ok := itr.Peek(); !ok {
				t.Errorf("Peek is failing while there are stil future values")
			} else if strconv.Itoa(i+1) != val {
				t.Errorf("Peek is returning an incorrect value")
			}

			if strconv.Itoa(i) != itr.Value() {
				t.Errorf("Value is returning an incorrect value")
			}
		} else {
			if _, ok := itr.Peek(); ok {
				t.Error("Peek is returning an ok status when there is nothing left to Peek")
			}

			if strconv.Itoa(i) != itr.Value() {
				t.Error("Value is failing to return the final entry")
			}
		}

		i++
	}

	if 5 != i {
		t.Error("Did not iterate the correct number of times")
	}
}

func TestStringIteratorSplit(t *testing.T) {
	itr := newStringIterator([]string{"1234567890"})

	if err := itr.Split(3); nil == err {
		t.Error("Split should return an error if called before Next")
	}

	itr.Next()
	if err := itr.Split(0); nil == err {
		t.Error("Split should return an error if the offset is < 1")
	}

	if err := itr.Split(9); nil == err {
		t.Error("Split should return an error if the offset is > len(data) -2")
	}

	if err := itr.Split(3); nil != err {
		t.Errorf("Split returned error: '%s' on a valid offset", err)
	}

	if 2 != len(itr.data) {
		t.Error("invalid data entry count after split")
	}

	if 3 != len(itr.data[0]) || 7 != len(itr.data[1]) {
		t.Errorf("Split did not happen at the correct offset %v", itr.data)
	}
}
