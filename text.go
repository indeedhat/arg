package arg

import (
	"fmt"
)

// normal
const (
	C_NONE = "0"

	C_BLACK        = "0;30"
	C_DARK_GRAY    = "1;30"
	C_RED          = "0;31"
	C_LIGHT_RED    = "1;31"
	C_GREEN        = "0;32"
	C_LIGHT_GREEN  = "1;32"
	C_BROWN_ORANGE = "0;33"
	C_YELLOW       = "1;33"
	C_BLUE         = "0;34"
	C_LIGHT_BLUE   = "1;34"
	C_PURPLE       = "0;35"
	C_PINK         = "1;35"
	C_CYAN         = "0;36"
	C_LIGHT_CYAN   = "1;36"
	C_LIGHT_GRAY   = "0;37"
	C_WHITE        = "1;37"
	C_DARK_RED     = "2;31"
	C_DARK_GREEN   = "2;32"
	C_DARK_YELLOW  = "2;33"
	C_DARK_BLUE    = "2;34"
	C_DARK_PURPLE  = "2;35"
	C_DARK_CYAN    = "2;36"
)

// wrap the given text in the color for shell output
func wrapText(text, color string) string {
	return fmt.Sprintf(`\033[%sm%s\033[%sm`, color, text, C_NONE)
}
