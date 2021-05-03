package interpreter

import (
	"github.com/peter9207/unischeme/lexer"
)

var globalScope map[string]FunctionDeclaration

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

func Exec(program lexer.Program) (result []ASTNode, err error) {
	result, err = ToAST(program.Expressions)
	if err != nil {
		return
	}

	return
}
