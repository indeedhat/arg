package arg

import (
	"fmt"
	"reflect"
	"strings"
)

type Parser struct {
	// Data bout the application
	Version string
	Title   string
	Usage   string

	// args to be parsed
	RawArgs []string
	itr     stringIterator

	// parsed args
	positional []expectedArg
	named      []expectedArg
	extras     []Argument

	// cursor related gubbins
	position int
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) parse(args []string) error {
	// setup the parser state
	p.RawArgs = args

	p.itr = newStringIterator(normalise(args))

	// do the parsing
	for p.itr.Next() {
		switch argType(p.itr.Value()) {
		case T_POSITIONAL:
			if err := p.parsePositional(); nil != err {
				return err
			}

		case T_NAMED:
			if err := p.parseNamed(); nil != err {
				return err
			}

		case T_DASHED:
			if err := p.parseDashed(); nil != err {
				return err
			}

		case T_GROUP:
			if err := p.parseGroup(); nil != err {
				return err
			}
		}
	}

	// finish up
	p.fillDefaults()

	return p.inflateExpectedArgs()
}

func (p *Parser) parsePositional() error {
	arg := p.findPositional()
	if nil == arg {
		p.addUnexpectedArg(T_POSITIONAL)
		return nil
	}

	arg.Values = append(arg.Values, p.itr.Value())
	return nil
}

func (p *Parser) parseNamed() error {
	key := p.itr.Value()[2:]
	arg := p.findNamed(key)

	if nil == arg {
		p.addUnexpectedArg(T_NAMED)
		return nil
	}

	arg.Values = append(arg.Values, p.tryFindValue(reflect.Bool == kind(arg.ScalarType)))
	return nil
}

func (p *Parser) parseDashed() error {
	key := p.itr.Value()[1:]
	arg := p.findNamed(key)

	if nil == arg {
		p.addUnexpectedArg(T_DASHED)
		return nil
	}

	arg.Values = append(arg.Values, p.tryFindValue(reflect.Bool == kind(arg.ScalarType)))
	return nil
}

func (p *Parser) parseGroup() error {
	// get the iterator to a standard state before the group bullshittery
	p.itr.Split(1)
	p.itr.Next()

	group := p.itr.Value()

	for i := 0; i < len(group); i++ {
		key := string(group[i])
		p.itr.Split(1)

		arg := p.findNamed(key)
		if nil == arg {
			p.addUnexpectedArg(T_DASHED)

			if strings.EqualFold(p.itr.Value(), "true") || strings.EqualFold(p.itr.Value(), "false") {
				break
			}
		} else {
			arg.Values = append(arg.Values, p.tryFindValue(reflect.Bool == kind(arg.ScalarType)))

			if p.itr.Value() == arg.Values[len(arg.Values)-1] {
				break
			}
		}
	}

	return nil
}

func (p *Parser) addUnexpectedArg(atype ArgType) {
	extra := Argument{
		Type: atype,
	}

	if T_POSITIONAL == atype {
		extra.Positon = p.position
		extra.Value = p.itr.Value()
	} else if T_DASHED == atype {
		extra.Key = p.itr.Value()
		extra.Value = p.tryFindValue(true)
	} else {
		extra.Key = p.itr.Value()
		extra.Value = p.tryFindValue(false)
	}

	p.extras = append(p.extras, extra)
}

func (p *Parser) tryFindValue(isBool bool) string {
	next, ok := p.itr.Peek()
	if !ok {
		return "true"
	}

	if T_POSITIONAL == argType(next) {
		if !isBool || strings.EqualFold(next, "true") || strings.EqualFold(next, "false") {
			p.itr.Next()
			return next
		}
	}

	return "true"
}

func (p *Parser) addExpectedArg(keys string, config *ArgConfig) *expectedArg {
	expected := expectedArg{
		Keys:     strings.Split(keys, ","),
		Default:  config.Default,
		Override: config.Override,
		Required: config.Required,
		Usage:    config.Usage,
		Validate: config.Validate,
	}

	if "" == keys {
		p.positional = append(p.positional, expected)
	} else {
		p.named = append(p.named, expected)
	}

	return &expected
}

func (p *Parser) findPositional() *expectedArg {
	if p.position < len(p.positional) {
		p.position++
		return &p.positional[p.position-1]
	}

	return nil
}

func (p *Parser) findNamed(key string) *expectedArg {
	for _, arg := range p.named {
		for _, k := range arg.Keys {
			if k == key {
				return &arg
			}
		}
	}

	return nil
}

func (p *Parser) fillDefaults() {
	for _, arg := range p.positional {
		if 0 != len(arg.Default) && 0 == len(arg.Values) {
			arg.Values = arg.Default
		}
	}

	for _, arg := range p.named {
		if 0 != len(arg.Default) && 0 == len(arg.Values) {
			arg.Values = arg.Default
		}
	}
}

func (p *Parser) inflateExpectedArgs() error {
	for i, arg := range p.positional {
		if arg.Required && 0 == len(arg.Values) {
			return fmt.Errorf("Required positional argument missing at position %d", i)
		}

		if nil != arg.Validate {
			for _, val := range arg.Values {
				if err := arg.Validate(val); nil != err {
					return err
				}
			}
		}

		if err := arg.inflate(); nil != err {
			return err
		}

		if arg.Override && 0 < len(arg.Values) {
			return nil
		}
	}

	for _, arg := range p.named {
		if arg.Required && 0 == len(arg.Values) {
			return fmt.Errorf("Required named argument '%s' missing", arg.Keys[0])
		}

		if nil != arg.Validate {
			for _, val := range arg.Values {
				if err := arg.Validate(val); nil != err {
					return err
				}
			}
		}

		if err := arg.inflate(); nil != err {
			return err
		}

		if arg.Override && 0 < len(arg.Values) {
			return nil
		}
	}

	return nil
}
