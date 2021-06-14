package interpreter

import (
	"errors"
	"fmt"
	"github.com/peter9207/unischeme/lexer"
)

type Program struct {
	Scope map[string]Expression
	Main  *FunctionCall
	Args  []string
}

type Expression interface {
	Resolve(scope map[string]Expression) (Value, error)
	Type() string
}

type Value interface {
	String() string
	Type() string
	Resolve(map[string]Expression) (Value, error)
	MarshalJSON() ([]byte, error)
}

func Exec(program lexer.Program) (result string, err error) {
	p, err := ToProgram(program.Expressions)
	if err != nil {
		return
	}

	v, err := Eval(p)
	if err != nil {
		return
	}
	result = v.String()
	return
}

var ErrCannotParseValue = errors.New("cannot parse value")
var ErrUndefinedIdentifier = errors.New("undefined identifier")

func Eval(program Program) (v Value, err error) {

	scope := map[string]Expression{}
	fnScope := map[string]FunctionDeclaration{}

	// for _, t := range ast {
	// 	var r []string
	// 	r, err = eval(t, scope, fnScope)
	// 	if err != nil {
	// 		return
	// 	}
	// 	results = append(results, r...)
	// }

	return
}

func eval(t Expression, scope map[string]Expression, functionScope map[string]FunctionDeclaration) (results []string, err error) {
	switch t.(type) {
	case IntValue, StringValue:
		result, ok := t.(Value)
		if !ok {
			err = ErrCannotParseValue
			return
		}
		results = append(results, result.String())

	case *FunctionDeclaration:
		fn := t.(*FunctionDeclaration)
		functionScope[fn.Name] = *fn

	case FunctionCall:
		fn := t.(FunctionCall)
		var v Value
		v, err = fn.Resolve(scope)
		if err != nil {
			return
		}

		results = append(results, v.String())

	default:
		fmt.Printf("unknown ast syntax %T, skipping...\n", t)
	}
	return
}
