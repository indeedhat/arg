package arg

type Parser struct {
	Version string
	Title   string
	Usage   string

	RawArgs []string
}

func NewParser() *Parser {
	return &Parser{}
}
