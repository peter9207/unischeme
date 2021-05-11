package interpreter

import (
	"errors"
	"fmt"
)

type FunctionCall struct {
	Name   string
	Params []Expression
}

func (i FunctionCall) Node() interface{} {
	return i
}

func (i FunctionCall) Type() string {
	return "functionCall"
}

func (i FunctionCall) Children() []ASTNode {
	return []ASTNode{}
}

func (fn FunctionCall) Resolve(scope map[string]Expression, functionScope map[string]FunctionDeclaration) (value Value, err error) {

	switch fn.Name {
	case "plus":
		return fn.resolvePlus(scope, functionScope)
	case "subtract":
		return fn.resolveMinus(scope, functionScope)
	default:
		return fn.resolve(scope, functionScope)
	}
	return
}

func (fn FunctionCall) resolvePlus(scope map[string]Expression, functionScope map[string]FunctionDeclaration) (value Value, err error) {

	if len(fn.Params) != 2 {
		err = fmt.Errorf("wrong number of arguments for plus %d", len(fn.Params))
		return
	}

	v1, err := fn.Params[0].Resolve(scope, functionScope)
	if err != nil {
		return
	}
	n1, ok := v1.(IntValue)
	if !ok {
		err = fmt.Errorf("wrong argument type for plus %T", v1)
		return
	}

	v2, err := fn.Params[1].Resolve(scope, functionScope)
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

func (fn FunctionCall) resolveMinus(scope map[string]Expression, functionScope map[string]FunctionDeclaration) (value Value, err error) {

	if len(fn.Params) != 2 {
		err = fmt.Errorf("wrong number of arguments for plus %d", len(fn.Params))
		return
	}

	v1 := fn.Params[0]
	n1, ok := v1.(IntValue)
	if !ok {
		err = fmt.Errorf("wrong argument type for plus %T", v1)
		return
	}

	v2 := fn.Params[1]
	n2, ok := v2.(IntValue)
	if !ok {
		err = fmt.Errorf("wrong argument type for plus %T", v1)
		return
	}
	value = IntValue{
		Value: n1.Value - n2.Value,
	}
	return
}

func (fn FunctionCall) resolve(scope map[string]Expression, functionScope map[string]FunctionDeclaration) (value Value, err error) {
	name := fn.Name
	declared, ok := functionScope[name]
	if !ok {
		err = ErrUndefinedIdentifier
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
		v, err = fn.Params[i].Resolve(scope, functionScope)
		if err != nil {
			return
		}

		nestedScope[id] = v
	}

	value, err = declared.Definition.Resolve(nestedScope, functionScope)

	return
}
