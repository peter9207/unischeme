package interpreter

type FunctionCall struct {
	Name   string
	Params []ASTNode
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
