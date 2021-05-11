package interpreter

import (
	"errors"
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
