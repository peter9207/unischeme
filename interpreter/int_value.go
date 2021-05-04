package interpreter

import (
	"fmt"
)

type IntValue struct {
	Value int
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

func (i IntValue) Resolve() string {
	return i.String()
}
