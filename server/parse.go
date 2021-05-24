package server

import (
	"errors"
	"fmt"
	"github.com/peter9207/unischeme/interpreter"
)

var ErrNoType = errors.New("cannot deserialize request, no type value")

func parseExpression(exp map[string]interface{}) (e interpreter.Expression, err error) {

	t, ok := exp["type"]
	if !ok {
		err = ErrNoType
		return
	}

	switch t {
	case "identifier":
		nameField, ok := exp["name"]
		if !ok {
			errors.New("identifier must have name field")
			return
		}

		name, ok := nameField.(string)
		if !ok {
			errors.New("identifier name field must be string")
			return

		}

		e = interpreter.Identifier{
			Name: name,
		}
		return

	case "functionCall":
		f := interpreter.FunctionCall{}

		n, ok := exp["name"]
		if !ok {
			errors.New("functionCall must have name field")
			return
		}
		f.Name, ok = n.(string)
		if !ok {
			errors.New("functionCall name must be string")
			return
		}

		p, ok := exp["params"]
		if !ok {
			errors.New("functionCall must have params field")
			return
		}

		p1, ok := p.([]map[string]interface{})
		if !ok {
			errors.New("functionCall failed to parse params")
			return
		}

		for _, v := range p1 {
			var e1 interpreter.Expression
			e1, err = parseExpression(v)
			if err != nil {
				err = errors.New("functionCall failed to parse params")
				return
			}
			f.Params = append(f.Params, e1)
		}
		e = f
		return

	case "int":
		v, ok := exp["value"]
		if !ok {
			errors.New("int value must have value field")
			return
		}

		value, ok := v.(int)
		if !ok {
			errors.New("int value must be int")
		}

		e = interpreter.IntValue{
			Value: value,
		}
		return

	case "string":
		v, ok := exp["value"]
		if !ok {
			errors.New("string value must have value field")
			return
		}
		value, ok := v.(string)
		if !ok {
			errors.New("string value must be string")
			return
		}

		e = interpreter.StringValue{
			Value: value,
		}

		return
	default:
		err = fmt.Errorf("unknown type %s", t)
		return
	}
}
