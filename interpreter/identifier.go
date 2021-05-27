package interpreter

import (
	"encoding/json"
)

type Identifier struct {
	Name string `json:"name"`
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

func (f *Identifier) MarshalJSON() (data []byte, err error) {

	i := map[string]interface{}{
		"name": f.Name,
		"type": f.Type(),
	}

	data, err = json.Marshal(i)

	return
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
