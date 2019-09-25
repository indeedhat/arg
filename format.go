package arg

import (
	"strings"
)

func normalise(args []string) []string {
	var normal []string

	for _, arg := range args {
		if T_POSITIONAL == argType(arg) {
			normal = append(normal, arg)
			continue
		}

		if !strings.Contains(arg, "=") {
			normal = append(normal, arg)
			continue
		}

		normal = append(normal, strings.SplitN(arg, "=", 2)...)
	}

	return normal
}
