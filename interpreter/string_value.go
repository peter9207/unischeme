package interpreter

type StringValue struct {
	Value string `json:"value"`
}

func (i StringValue) Node() interface{} {
	return i
}

func (i StringValue) Type() string {
	return "stringValue"
}

func (i StringValue) Children() []ASTNode {
	return []ASTNode{}
}

func (i StringValue) String() string {
	return i.Value
}

func (i StringValue) Resolve(_ map[string]Expression, _ map[string]FunctionDeclaration) (Value, error) {
	return i, nil
}
