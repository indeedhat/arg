package arg

import (
	"strings"
)

type ArgType int

const (
	T_POSITIONAL ArgType = iota
	T_DASHED
	T_NAMED
	T_GROUP
)

func argType(arg string) ArgType {
	if strings.HasPrefix(arg, "--") && 2 < len(arg) && '-' != arg[2] {
		return T_NAMED
	}

	if strings.HasPrefix(arg, "-") && 1 < len(arg) && '-' != arg[1] {
		if 2 < len(arg) {
			return T_GROUP
		}

		return T_DASHED
	}

	return T_POSITIONAL
}
