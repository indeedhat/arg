package arg

import (
	"fmt"
	"strings"
)

func generateHelpHeader(p *Parser) string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n\n%s\n",
		wrapText(p.Title, C_WHITE),
		wrapText(strings.Repeat("-", len(p.Title)), C_LIGHT_GRAY),
		wrapText(fmt.Sprintf("Version: %s", p.Version), C_LIGHT_GRAY),
		wrapText(p.Usage, C_LIGHT_GRAY),
	)
}

func generateHelpArgument(e *expectedArg) string {
	return ""
}
