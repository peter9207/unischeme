package interpreter

import (
	"encoding/json"
	"errors"
	"fmt"
)

type FunctionCall struct {
	Name   string       `json:"name"`
	Params []Expression `json:"params"`
}

var ErrInvalidJSON = errors.New("unable to parse json object")

func (i FunctionCall) Node() interface{} {
	return i
}

func (i FunctionCall) Type() string {
	return "functionCall"
}

func (f *FunctionCall) MarshalJSON() (data []byte, err error) {

	i := map[string]interface{}{
		"name":   f.Name,
		"params": f.Params,
		"type":   f.Type(),
	}

	data, err = json.Marshal(i)
	return
}

func (fn FunctionCall) Resolve(scope map[string]Expression) (value Value, err error) {

	switch fn.Name {
	case "plus":
		return fn.resolvePlus(scope)
	case "subtract":
		return fn.resolveMinus(scope)
	default:
		return fn.resolve(scope)
	}
	return
}

func (fn FunctionCall) resolvePlus(scope map[string]Expression) (value Value, err error) {

	if len(fn.Params) != 2 {
		err = fmt.Errorf("wrong number of arguments for plus %d", len(fn.Params))
		return
	}

	v1, err := fn.Params[0].Resolve(scope)
	if err != nil {
		return
	}
	n1, ok := v1.(IntValue)
	if !ok {
		err = fmt.Errorf("wrong argument type for plus %T", v1)
		return
	}

	v2, err := fn.Params[1].Resolve(scope)
	if err != nil {
		return
	}
	n2, ok := v2.(IntValue)
	if !ok {
		err = fmt.Errorf("wrong argument type for plus %T", v1)
		return
	}
	value = IntValue{
		Value: n2.Value + n1.Value,
	}
	return
}

func (fn FunctionCall) resolveMinus(scope map[string]Expression) (value Value, err error) {

	if len(fn.Params) != 2 {
		err = fmt.Errorf("wrong number of arguments for subtract %d", len(fn.Params))
		return
	}

	v1, err := fn.Params[0].Resolve(scope)
	n1, ok := v1.(IntValue)
	if !ok {
		err = fmt.Errorf("wrong argument type for subtract %T", v1)
		return
	}

	v2, err := fn.Params[1].Resolve(scope)
	n2, ok := v2.(IntValue)
	if !ok {
		err = fmt.Errorf("wrong argument type for subtract %T", v1)
		return
	}
	value = IntValue{
		Value: n1.Value - n2.Value,
	}
	return
}

func (fn FunctionCall) resolve(scope map[string]Expression) (value Value, err error) {
	name := fn.Name
	exist, ok := scope[name]
	if !ok {
		err = fmt.Errorf("unindentifier identifier %s", name)
		return
	}

	declared, ok := exist.(*Function)
	if !ok {
		err = fmt.Errorf("value is not a function  %s", name)
		return
	}

	nestedScope := map[string]Expression{}
	for k, v := range scope {
		nestedScope[k] = v
	}

	if len(declared.Params) != len(fn.Params) {
		err = errors.New("function call wrong number of params")
		return
	}

	for i := range declared.Params {
		var v Value
		id := declared.Params[i]
		v, err = fn.Params[i].Resolve(scope)
		if err != nil {
			return
		}

		nestedScope[id] = v
	}

	value, err = declared.Definition.Resolve(nestedScope)

	return
}
