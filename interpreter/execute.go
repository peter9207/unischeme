package interpreter

import (
	"errors"
	"fmt"
	"github.com/peter9207/unischeme/lexer"
)

type ASTNode interface {
	Node() interface{}
	Type() string
	Children() []ASTNode
}

type Expression interface {
	Resolve(scope map[string]ASTNode) (Value, error)
}

type Statement interface {
	Perform(scope map[string]ASTNode) error
}

type Value interface {
	String() string
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

func evalFunctionDeclaration(fn FunctionDeclaration, scope map[string]ASTNode) (err error) {
	scope[fn.Name] = fn
	return
}

func evalFunctionCall(fn FunctionCall, scope map[string]ASTNode) (value string, err error) {
	name := fn.Name
	declared, ok := scope[name]
	if !ok {
		err = ErrUndefinedIdentifier
		return
	}

	return
}

func eval(t ASTNode, scope map[string]ASTNode) (results []string, err error) {
	switch t.(type) {
	case IntValue:
		result, ok := t.(Value)
		if !ok {
			err = ErrCannotParseValue
			return
		}
		results = append(results, result.String())
	case StringValue:
		result, ok := t.(Value)
		if !ok {
			err = ErrCannotParseValue
			return
		}
		results = append(results, result.String())

	case FunctionDeclaration:
		fn := t.(FunctionDeclaration)
		scope[fn.Name] = fn

	case FunctionCall:
		fn := t.(FunctionCall)
		name := fn.Name
		declared, ok := scope[name]
		if !ok {
			err = ErrUndefinedIdentifier
			return
		}

	default:
		fmt.Printf("unknown ast syntax %T, skipping...\n", t)
	}
	return
}

func Eval(ast []ASTNode) (results []string, err error) {

	var scope map[string]FunctionDeclaration

	for _, t := range ast {
		r, err := eval(t, scope)
		results = append(results, r)
	}

	return
}
