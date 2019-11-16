package arg

import (
	"reflect"
)

func (p *Parser) Parse(args []string) error {
	return p.parse(args)
}

func (p *Parser) Command(key string, config *CommandConfig) {
	p.addCommand(key, config)
}

func (p *Parser) String(arg *string, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Strings(arg *[]string, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Int(arg *int, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Ints(arg *[]int, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Int8(arg *int8, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Int8s(arg *[]int8, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Int16(arg *int16, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Int16s(arg *[]int16, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Int32(arg *int32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Int32s(arg *[]int32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Int64(arg *int64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Int64s(arg *[]int64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Uint(arg *uint, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Uints(arg *[]uint, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Uint8(arg *uint8, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Uint8s(arg *[]uint8, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Uint16(arg *uint16, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Uint16s(arg *[]uint16, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Uint32(arg *uint32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Uint32s(arg *[]uint32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Uint64(arg *uint64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Uint64s(arg *[]uint64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Float32(arg *float32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Float32s(arg *[]float32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Float64(arg *float64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Float64s(arg *[]float64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Bool(arg *bool, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}

func (p *Parser) Bools(arg *[]bool, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg).Elem()
}
