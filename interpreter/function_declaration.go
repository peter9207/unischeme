package interpreter

type FunctionDeclaration struct {
	Name       string     `json:"name"`
	Params     []string   `json:"params"`
	Definition Expression `json:"definition"`
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
