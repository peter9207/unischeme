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

// func exec(req InterpretRequest) (v Value, err error) {

// 	scope := map[string]interpreter.Expression{}

// 	for k, v1 := range req.VarScope {
// 		parseExpression()
// 		var e interpreter.Expression
// 		e, err = parseValues(v1)
// 		if err != nil {
// 			return
// 		}
// 		scope[k] = e
// 	}

// 	params := []interpreter.Expression{}
// 	for _, v := range req.Params {
// 		var e interpreter.Expression
// 		e, err := parseValues(v)
// 		if err != nil {
// 			return
// 		}
// 		params = append(params, e)
// 	}

// 	fn := interpreter.FunctionCall{
// 		Name:   req.Name,
// 		Params: params,
// 	}

// 	fn.Resolve()

// }
