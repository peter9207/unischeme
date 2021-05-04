package interpreter

type FunctionDeclaration struct {
	Name       string
	Params     []string
	Definition Expression
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

func (f FunctionDeclaration) Perform(scope map[string]ASTNode) error {
	scope[f.Name] = f
	return nil
}
