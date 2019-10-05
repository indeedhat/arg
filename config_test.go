package arg

import (
	"errors"
	"strconv"
	"testing"
)

func TestValidatorPass(t *testing.T) {
	args := []string{
		"100",
	}

	var sub int

	parser := NewParser()
	parser.Int(&sub, "", &ArgConfig{
		Validate: func(val string) error {
			intVal, err := strconv.Atoi(val)
			if nil != err {
				return err
			}

			if 30 > intVal || 120 < intVal {
				return errors.New("The number was not in the range of 30-120")
			}

			return nil
		},
	})

	if err := parser.Parse(args); nil != err {
		t.Fatalf("Arg faled to validate with message: %s", err)
	}
}

func TestValidatorFail(t *testing.T) {
	args := []string{
		"130",
	}

	var sub int

	parser := NewParser()
	parser.Int(&sub, "", &ArgConfig{
		Validate: func(val string) error {
			intVal, err := strconv.Atoi(val)
			if nil != err {
				return err
			}

			if 30 > intVal || 120 < intVal {
				return errors.New("The number was not in the range of 30-120")
			}

			return nil
		},
	})

	if err := parser.Parse(args); nil == err {
		t.Fatalf("Faled to validate with message: %s", err)
	}
}
