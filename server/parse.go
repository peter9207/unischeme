package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/peter9207/unischeme/interpreter"
)

var ErrNoType = errors.New("cannot deserialize request, no type value")

func parseExpression(exp map[string]interface{}) (e expression, err error) {

	t, ok := exp["type"]
	if !ok {
		err = ErrNoType
		return
	}

	switch t {
	case "identifier":
		name, ok := exp["name"]
		if !ok {
			errors.New("identifier must have name field")
			return
		}

		e = interpreter.Identifier{
			Name: name,
		}
		return

	case "functionCall":
		f := interpreter.FunctionCall{}

		f.Name, ok = exp["name"]
		if !ok {
			errors.New("functionCall must have name field")
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
			e1, err := parseExpression(v)
			if err != nil {
				errors.New("functionCall failed to parse params")
				return
			}
			f.Params = append(f.Params, e1)
		}
		e = f
		return

	case "int":
		v, ok := exp["value"]
		if !ok {
			errors.New("identifier must have name field")
			return
		}

		e = interpreter.IntValue{
			Value: v,
		}
		return

	case "string":
		v, ok := exp["value"]
		if !ok {
			errors.New("identifier must have name field")
			return
		}

		e = interpreter.StringValue{
			Value: v,
		}

		return
	default:
		err = fmt.Errorf("unknown type %s", t)
		return
	}
}
