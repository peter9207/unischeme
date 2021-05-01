package interpreter

type IntValue struct {
	Value int
}

// func (i IntValue) Value() interface{} {
// 	return i.Value
// }

func (i IntValue) Node() interface{} {
	return i
}

func (i IntValue) Type() string {
	return "intValue"
}

func (i IntValue) Children() []ASTNode {
	return []ASTNode{}
}
