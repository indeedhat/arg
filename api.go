package arg

import (
	"reflect"
)

func (p *Parser) Parse(args []string) error {
	return p.parse(args)
}

func (p *Parser) String(arg *string, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Strings(arg *[]string, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Int8(arg *int8, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Int8s(arg *[]int8, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Int16(arg *int16, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Int16s(arg *[]int16, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Int32(arg *int32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Int32s(arg *[]int32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Int64(arg *int64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Int64s(arg *[]int64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Uuint8(arg *uint8, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Uuint8s(arg *[]uint8, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Uuint16(arg *uint16, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Uuint16s(arg *[]uint16, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Uuint32(arg *uint32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Uuint32s(arg *[]uint32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Uuint64(arg *uint64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Uuint64s(arg *[]uint64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Float32(arg *float32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Float32s(arg *[]float32, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Float64(arg *float64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Float64s(arg *[]float64, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Bool(arg *bool, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}

func (p *Parser) Bools(arg *[]bool, keys string, config *ArgConfig) {
	expected := p.addExpectedArg(keys, config)
	expected.ScalarType = reflect.ValueOf(arg)
}
