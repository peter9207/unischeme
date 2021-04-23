package interpreter

type IntValue struct {
	Value int
}

func (i IntValue) Node() interface{} {
	return i
}

func (i IntValue) Type() string {
	return "int"
}

func (i IntValue) Children() []ASTNode {
	return []ASTNode{}
}
