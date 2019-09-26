package arg

import (
	"errors"
)

type stringIterator struct {
	data    []string
	current int
}

func newStringIterator(data []string) stringIterator {
	return stringIterator{
		data:    data,
		current: -1,
	}
}

func (itr *stringIterator) Next() bool {
	if itr.current == len(itr.data)-1 {
		return false
	}

	itr.current++
	return true
}

func (itr *stringIterator) Value() string {
	if -1 == itr.current {
		return ""
	}

	return itr.data[itr.current]
}

func (itr *stringIterator) Peek() (string, bool) {
	if itr.current == len(itr.data)-1 {
		return "", false
	}

	return itr.data[itr.current+1], true
}

func (itr *stringIterator) Split(offset int) error {
	if -1 == itr.current {
		return errors.New("cannot split before the first call to Next")
	}

	if 1 > offset {
		return errors.New("offset cannot be less than 1")
	}

	currentVal := itr.Value()
	if len(currentVal)-1 <= offset {
		return errors.New("offset cannot be greater than len(value) -2")
	}

	// TODO: do this in a less stupid way when im less tired
	tmp := itr.data[:itr.current]
	tmp = append(tmp, currentVal[:offset])
	tmp = append(tmp, currentVal[offset:])
	tmp = append(tmp, itr.data[itr.current+1:]...)

	itr.data = tmp

	return nil
}
