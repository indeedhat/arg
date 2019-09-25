package arg

type Parser struct {
	Version string
	Title   string
	Usage   string

	RawArgs []string

	expected []expectedArg
	extras   []Argument
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) parse(args []string) error {
	return nil
}
