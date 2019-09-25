package arg

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Argument struct {
	Type  ArgType
	Key   string
	Value string
}

type expectedArg struct {
	Type       ArgType
	ScalarType reflect.Value

	Offset int
	Keys   []string
	Values []string

	Usage string

	Default  []string
	Required bool
	Override bool
	Validate ValidatorFunc
}

func (e *expectedArg) inflate() error {
	var err error
	switch kind(e.ScalarType) {
	case reflect.Bool:
		err = inflateBool(e)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		err = inflateInt(e)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		err = inflateUint(e)
	case reflect.Float32, reflect.Float64:
		err = inflateFloat(e)
	case reflect.String:
		err = inflateString(e)
	default:
		// do nothing
	}

	return err
}

func newExpectedArg(keys string, config *ArgConfig) expectedArg {
	return expectedArg{
		Keys:     strings.Split(keys, ","),
		Default:  config.Default,
		Override: config.Override,
		Required: config.Required,
		Usage:    config.Usage,
		Validate: config.Validate,
	}
}

func inflateString(arg *expectedArg) error {
	if isSlice(arg.ScalarType) {
		arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(arg.Values)))
	} else {
		arg.ScalarType.SetString(arg.Values[len(arg.Values)-1])
	}

	return nil
}

func inflateInt(arg *expectedArg) error {
	var err error
	var i int64

	if isSlice(arg.ScalarType) {
		switch kind(arg.ScalarType) {
		case reflect.Int, reflect.Int32:
			err = inflateInt32slice(arg)
		case reflect.Int64:
			err = inflateInt64slice(arg)
		case reflect.Int16:
			err = inflateInt16slice(arg)
		case reflect.Int8:
			err = inflateInt8slice(arg)
		}
	} else {
		i, err = strconv.ParseInt(arg.Values[len(arg.Values)-1], 10, 64)
		if err == nil {
			arg.ScalarType.SetInt(i)
		}
	}

	return generateError(err, arg, "int")
}

func inflateInt64slice(arg *expectedArg) error {
	var err error
	var i int64
	var newvals []int64

	for _, v := range arg.Values {
		if i, err = strconv.ParseInt(v, 10, 64); nil != err {
			break
		}

		newvals = append(newvals, int64(i))
	}

	arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	return err
}

func inflateInt32slice(arg *expectedArg) error {
	var err error
	var i int64
	var newvals []int32

	for _, v := range arg.Values {
		if i, err = strconv.ParseInt(v, 10, 32); nil != err {
			break
		}

		newvals = append(newvals, int32(i))
	}

	arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	return err
}

func inflateInt16slice(arg *expectedArg) error {
	var err error
	var i int64
	var newvals []int16

	for _, v := range arg.Values {
		if i, err = strconv.ParseInt(v, 10, 16); nil != err {
			break
		}

		newvals = append(newvals, int16(i))
	}

	arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	return err
}

func inflateInt8slice(arg *expectedArg) error {
	var err error
	var i int64
	var newvals []int8

	for _, v := range arg.Values {
		if i, err = strconv.ParseInt(v, 10, 8); nil != err {
			break
		}

		newvals = append(newvals, int8(i))
	}

	arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	return err
}

func inflateUint(arg *expectedArg) error {
	var i uint64
	var err error

	if isSlice(arg.ScalarType) {
		switch kind(arg.ScalarType) {
		case reflect.Uint, reflect.Uint32:
			err = inflateUint32slice(arg)
		case reflect.Uint64:
			err = inflateUint64slice(arg)
		case reflect.Uint16:
			err = inflateUint16slice(arg)
		case reflect.Uint8:
			err = inflateUint8slice(arg)
		}
	} else {
		i, err = strconv.ParseUint(arg.Values[len(arg.Values)-1], 10, 64)
		if err == nil {
			arg.ScalarType.SetUint(i)
		}
	}

	return generateError(err, arg, "uint")
}

func inflateUint64slice(arg *expectedArg) error {
	var err error
	var i uint64
	var newvals []uint64

	for _, v := range arg.Values {
		if i, err = strconv.ParseUint(v, 10, 64); nil != err {
			break
		}

		newvals = append(newvals, uint64(i))
	}

	arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	return err
}

func inflateUint32slice(arg *expectedArg) error {
	var err error
	var i uint64
	var newvals []uint32

	for _, v := range arg.Values {
		if i, err = strconv.ParseUint(v, 10, 32); nil != err {
			break
		}

		newvals = append(newvals, uint32(i))
	}

	arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	return err
}

func inflateUint16slice(arg *expectedArg) error {
	var err error
	var i uint64
	var newvals []uint16

	for _, v := range arg.Values {
		if i, err = strconv.ParseUint(v, 10, 16); nil != err {
			break
		}

		newvals = append(newvals, uint16(i))
	}

	arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	return err
}

func inflateUint8slice(arg *expectedArg) error {
	var err error
	var i uint64
	var newvals []uint8

	for _, v := range arg.Values {
		if i, err = strconv.ParseUint(v, 10, 8); nil != err {
			break
		}

		newvals = append(newvals, uint8(i))
	}

	arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	return err
}

func inflateFloat(arg *expectedArg) error {
	var i float64
	var err error

	if isSlice(arg.ScalarType) {
		switch kind(arg.ScalarType) {
		case reflect.Float32:
			err = inflateFloat32slice(arg)
		case reflect.Float64:
			err = inflateFloat64slice(arg)
		}
	} else {
		i, err = strconv.ParseFloat(arg.Values[len(arg.Values)-1], 64)
		if err == nil {
			arg.ScalarType.SetFloat(i)
		}
	}

	return generateError(err, arg, "float")
}

func inflateFloat32slice(arg *expectedArg) error {
	var err error
	var i float64
	var newvals []float32

	for _, v := range arg.Values {
		if i, err = strconv.ParseFloat(v, 32); nil != err {
			break
		}

		newvals = append(newvals, float32(i))
	}

	arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	return err
}

func inflateFloat64slice(arg *expectedArg) error {
	var err error
	var i float64
	var newvals []float64

	for _, v := range arg.Values {
		if i, err = strconv.ParseFloat(v, 64); nil != err {
			break
		}

		newvals = append(newvals, float64(i))
	}

	arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	return err
}

func inflateBool(arg *expectedArg) error {
	var i bool
	var err error

	if isSlice(arg.ScalarType) {
		var newvals []bool
		for _, v := range arg.Values {
			if i, err = strconv.ParseBool(v); nil == err {
				newvals = append(newvals, i)
			}
		}
		arg.ScalarType.Set(reflect.AppendSlice(arg.ScalarType, reflect.ValueOf(newvals)))
	} else {
		i, err = strconv.ParseBool(arg.Values[len(arg.Values)-1])
		if err == nil {
			arg.ScalarType.SetBool(i)
		}
	}

	return generateError(err, arg, "boolean")
}

func generateError(err error, arg *expectedArg, typ string) error {
	if nil == err {
		return nil
	}

	if T_POSITIONAL == arg.Type {
		return errors.New(fmt.Sprintf("Arg given at position %d was not of type %s", arg.Offset, typ))
	} else {
		return errors.New(fmt.Sprintf("Arg given for %s was notof type %s", arg.Keys[0], typ))
	}
}
