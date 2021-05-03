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

type AST struct {
	Children []AST
	Type     string
	Value    string
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

func Eval(ast []ASTNode) (results []string, err error) {

	// var scope map[string]FunctionDeclaration

	for _, t := range ast {

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

		default:
			fmt.Printf("unknown ast syntax %T, skipping...\n", t)
		}
	}

	return
}
