package arg

import (
	"testing"
)

func TestHeader(t *testing.T) {
	parser := NewParser()

	parser.Usage = "Some usage text"
	parser.Title = "Test App"
	parser.Version = "0.1.0a"

	expected := `\033[1;37mTest App\033[0m
\033[0;37m--------\033[0m
\033[0;37mVersion: 0.1.0a\033[0m

\033[0;37mSome usage text\033[0m
`

	if expected != generateHelpHeader(parser) {
		t.Errorf("Did not producte a valid header string\n%s\n\n%s", expected, generateHelpHeader(parser))
	}
}
