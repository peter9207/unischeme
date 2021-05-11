package interpreter

import (
	"fmt"
)

type IntValue struct {
	Value int `json:"value"`
}

func (i IntValue) Node() interface{} {
	return i
}

func (i IntValue) Type() string {
	return "intValue"
}

func (i IntValue) Children() []ASTNode {
	return []ASTNode{}
}

func (i IntValue) String() string {
	return fmt.Sprintf("%v", i.Value)
}

func (i IntValue) Resolve(_ map[string]Expression, _ map[string]FunctionDeclaration) (Value, error) {
	return i, nil
}
