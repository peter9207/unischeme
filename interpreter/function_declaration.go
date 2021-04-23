package interpreter

type FunctionDeclaration struct {
	Name       string
	params     []string
	Definition AST
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
