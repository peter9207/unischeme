package server

import (
	"errors"
	"github.com/peter9207/unischeme/interpreter"
)

type Value struct {
	Type        string `json:"type"`
	IntValue    int    `json:"int_value"`
	StringValue string `json:"string_value"`
}

type InterpretRequest struct {
	URL      string            `json:"url"`
	VarScope map[string]Value  `json:"var_scope"`
	FnScope  map[string]string `json:"fn_scope"`
	Name     string            `json:"name"`
	Params   []Value           `json:"values"`
}

var ErrInvalidParamValue = errors.New("inavlid value in parameter")

func parseValues(v Value) (e interpreter.Expression, err error) {
	switch v.Type {
	case "int":
		e = interpreter.IntValue{
			Value: v.IntValue,
		}
	case "string":
		e = interpreter.StringValue{
			Value: v.StringValue,
		}
	default:
		err = ErrInvalidParamValue
	}
	return
}

func exec(req InterpetRequest) (v Value, err error) {

	scope := map[string]Expression{}

	for k, v := range req.VarScope {
		e, err := parseValues(v)
		if err != nil {
			return
		}
		scope[k] = e
	}

	params := []Expression{}
	for _, v := range req.Params {
		e, err := parseValues(v)
		if err != nil {
			return
		}
		params = append(params, e)
	}

	fn := interpreter.FunctionCall{
		Name:   req.Name,
		Params: params,
	}

	fn.Resolve()

}
