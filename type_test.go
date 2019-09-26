package arg

import (
	"testing"
)

func TestArgTypeDashed(t *testing.T) {
	tests := []string{
		"-e",
		"-1",
	}

	for _, arg := range tests {
		if T_DASHED != argType(arg) {
			t.Errorf("failed to parse %s as dashed arg", arg)
		}
	}
}

func TestArgTypeGroup(t *testing.T) {
	tests := []string{
		"-djgd",
		"-t=3",
		"-dgef=3dw-d",
	}

	for _, arg := range tests {
		if T_GROUP != argType(arg) {
			t.Errorf("failed to parse %s as a dash group", arg)
		}
	}
}

func TestArgTypeNamed(t *testing.T) {
	tests := []string{
		"--e",
		"--1",
		"--djgd",
		"--thing=other",
		"--thing-1",
	}

	for _, arg := range tests {
		if T_NAMED != argType(arg) {
			t.Errorf("failed to parse %s as named arg", arg)
		}
	}
}

func TestArgPositional(t *testing.T) {
	tests := []string{
		"thing-1",
		"thing=1",
		"353",
		"-",
		"--",
		"---",
	}

	for _, arg := range tests {
		if T_POSITIONAL != argType(arg) {
			t.Errorf("failed to parse %s as positional arg", arg)
		}
	}
}
