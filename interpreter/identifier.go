package interpreter

type Identifier struct {
	Name string
}

func (i Identifier) Node() interface{} {
	return i
}
func (i Identifier) Type() string {
	return "identifier"
}
func (i Identifier) Children() []ASTNode {
	return []ASTNode{}
}
