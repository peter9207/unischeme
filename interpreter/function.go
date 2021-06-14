package interpreter

import (
	"encoding/json"
	"fmt"
)

type Function struct {
	Name       string     `json:"name"`
	Params     []string   `json:"params"`
	Definition Expression `json:"definition"`
}

func (f *Function) Node() interface{} {
	return f
}
func (f *Function) Type() string {
	return "function"
}

func (f *Function) Resolve(varScope map[string]Expression) (g Value, err error) {
	g = f
	return
}

func (f *Function) String() string {
	data, err := json.Marshal(f)
	if err != nil {
		fmt.Println("failed to stringify funciton", err)
	}

	return string(data)
}

func (f *Function) MarshalJSON() (data []byte, err error) {

	i := map[string]interface{}{
		"name":       f.Name,
		"params":     f.Params,
		"definition": f.Definition,
		"type":       f.Type(),
	}

	data, err = json.Marshal(i)

	return
}
