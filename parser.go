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
	command    []*commandArg
	positional []*expectedArg
	named      []*expectedArg
	extras     []*Argument

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
			p.parsePositional()

		case T_NAMED:
			p.parseNamed()

		case T_DASHED:
			p.parseDashed()

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

func (p *Parser) parsePositional() {
	arg := p.findPositional()
	if nil == arg {
		p.addUnexpectedArg(T_POSITIONAL)
		return
	}

	arg.Values = append(arg.Values, p.itr.Value())
}

func (p *Parser) parseNamed() {
	key := p.itr.Value()[2:]
	arg := p.findNamed(key)

	if nil == arg {
		p.addUnexpectedArg(T_NAMED)
		return
	}

	arg.Values = append(arg.Values, p.tryFindValue(reflect.Bool == kind(arg.ScalarType)))
}

func (p *Parser) parseDashed() {
	key := p.itr.Value()[1:]
	arg := p.findNamed(key)

	if nil == arg {
		p.addUnexpectedArg(T_DASHED)
		return
	}

	arg.Values = append(arg.Values, p.tryFindValue(reflect.Bool == kind(arg.ScalarType)))
}

func (p *Parser) parseGroup() error {
	// get the iterator to a standard state before the group bullshittery
	p.itr.Split(1)

	group, ok := p.itr.Peek()
	if !ok {
		return fmt.Errorf("Something went wrong with group parsing")
	}

	for i := 0; i < len(group); i++ {
		p.itr.Next()
		p.itr.Split(1)

		key := string(group[i])

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

func (p *Parser) addCommand(key string, meta *CommandConfig) {
	cmd := commandArg{
		Key:   key,
		Usage: meta.Usage,
		Title: meta.Title,
	}

	p.command = append(p.command, &cmd)
}

func (p *Parser) addUnexpectedArg(atype ArgType) {
	extra := &Argument{
		Type: atype,
	}

	if T_POSITIONAL == atype {
		extra.Positon = p.position
		extra.Value = p.itr.Value()
		p.position++
	} else if T_DASHED == atype {
		extra.Key = trimKey(p.itr.Value())
		extra.Value = p.tryFindValue(true)
	} else {
		extra.Key = trimKey(p.itr.Value())
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
	expected := &expectedArg{
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

	return expected
}

func (p *Parser) findPositional() *expectedArg {
	if p.position < len(p.positional) {
		if !isSlice(p.positional[p.position].ScalarType) {
			p.position++
			return p.positional[p.position-1]
		}

		return p.positional[p.position]
	}

	return nil
}

func (p *Parser) findNamed(key string) *expectedArg {
	for _, arg := range p.named {
		for _, k := range arg.Keys {
			if k == key {
				return arg
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

func trimKey(key string) string {
	return strings.TrimPrefix(strings.TrimPrefix(key, "--"), "-")
}
