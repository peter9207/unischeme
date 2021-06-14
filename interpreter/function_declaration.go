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

func (f *FunctionDeclaration) Resolve(varScope map[string]Expression) (g Value, err error) {
	varScope[f.Name] = f

	g = &Function{
		Name:       f.Name,
		Params:     f.Params,
		Definition: f.Definition,
	}
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
