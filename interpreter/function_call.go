package interpreter

type FunctionCall struct {
	Name   string
	Params []Expression
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

func (fn FunctionCall) Resolve(scope map[string]ASTNode) (value Value, err error) {
	name := fn.Name
	declared, ok := scope[name]
	if !ok {
		err = ErrUndefinedIdentifier
		return
	}

	for _, param := range fn.Params {
		resolved, err := param.Resolve(scope)
	}

	return
}
