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

func (i Identifier) Resolve(scope map[string]Expression, _ map[string]FunctionDeclaration) (value Value, err error) {

	name := i.Name
	declared, ok := scope[name]
	if !ok {
		err = ErrUndefinedIdentifier
		return
	}

	switch declared.(type) {
	case StringValue:
		value = declared.(StringValue)
		return
	case IntValue:
		value = declared.(IntValue)
		return
	default:
	}

	return
}
