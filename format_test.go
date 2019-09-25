package arg

import (
	"testing"
)

func TestNormalise(t *testing.T) {
	input := []string{
		"-sd",
		"--something",
		"else",
		"--with=val",
		"without=val",
	}

	expected := []string{
		"-sd",
		"--something",
		"else",
		"--with",
		"val",
		"without=val",
	}

	if !sliceEqual(expected, normalise(input)) {
		t.Error("Failed to normalise input")
	}
}
