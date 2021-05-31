package interpreter

import (
	"encoding/json"
)

type FunctionDeclaration struct {
	Name       string     `json:"name"`
	Params     []string   `json:"params"`
	Definition Expression `json:"definition"`
}

func (f *FunctionDeclaration) Node() interface{} {
	return f
}
func (f *FunctionDeclaration) Type() string {
	return "functionDeclaration"
}
func (f *FunctionDeclaration) Children() []ASTNode {
	return []ASTNode{}
}

func (f *FunctionDeclaration) Perform(varScope map[string]Expression, fnScope map[string]FunctionDeclaration) (err error) {
	fnScope[f.Name] = *f
	return
}

func (f *FunctionDeclaration) MarshalJSON() (data []byte, err error) {

	i := map[string]interface{}{
		"name":       f.Name,
		"params":     f.Params,
		"definition": f.Definition,
		"type":       f.Type(),
	}

	data, err = json.Marshal(i)

	return
}
