package interpreter

import (
	"errors"
	"fmt"
	"github.com/peter9207/unischeme/lexer"
)

type ASTNode interface {
	Type() string
}

type Expression interface {
	Resolve(scope map[string]Expression, functionScope map[string]FunctionDeclaration) (Value, error)
	Type() string
}

type Statement interface {
	Perform(scope map[string]ASTNode) error
	Type() string
}

type Value interface {
	String() string
	Type() string
	Resolve(map[string]Expression, map[string]FunctionDeclaration) (Value, error)
	MarshalJSON() ([]byte, error)
}

func Exec(program lexer.Program) (result []string, err error) {
	ast, err := ToAST(program.Expressions)
	if err != nil {
		return
	}

	result, err = Eval(ast)
	return
}

var ErrCannotParseValue = errors.New("cannot parse value")
var ErrUndefinedIdentifier = errors.New("undefined identifier")

func Eval(ast []ASTNode) (results []string, err error) {

	scope := map[string]Expression{}
	fnScope := map[string]FunctionDeclaration{}

	for _, t := range ast {
		var r []string
		r, err = eval(t, scope, fnScope)
		if err != nil {
			return
		}
		results = append(results, r...)
	}

	return
}

func eval(t ASTNode, scope map[string]Expression, functionScope map[string]FunctionDeclaration) (results []string, err error) {
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
		v, err = fn.Resolve(scope, functionScope)
		if err != nil {
			return
		}

		results = append(results, v.String())

	default:
		fmt.Printf("unknown ast syntax %T, skipping...\n", t)
	}
	return
}
