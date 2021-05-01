package interpreter

type FunctionDeclaration struct {
	Name       string
	Params     []string
	Definition ASTNode
}

func (f FunctionDeclaration) Node() interface{} {
	return f
}
func (f FunctionDeclaration) Type() string {
	return "functionDeclaration"
}
func (f FunctionDeclaration) Children() []ASTNode {
	return []ASTNode{}
}
