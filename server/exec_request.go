package server

import (
	"encoding/json"
	"errors"
	"github.com/peter9207/unischeme/interpreter"
)

func MakeExecRequest(url string, varScope map[string]interpreter.Value, fnScope map[string]interpreter.FunctionDeclaration, name string, params []interpreter.Value) (req ExecRequest, err error) {

	req.URL = url
	req.Name = name

	for _, v := range params {
		value := Value{}
		switch v.(type) {
		case interpreter.IntValue:
			s := v.(interpreter.IntValue)
			value.Type = "int"
			value.IntValue = s.Value
		case interpreter.StringValue:
			s := v.(interpreter.StringValue)
			value.Type = "string"
			value.StringValue = s.Value
		default:
			err = errors.New("unknonw interpreterValue type")
			return
		}
		req.Params = append(req.Params, value)
	}

	req.FnScope = fnScope

	return
}

type ExecRequest struct {
	URL      string                                     `json:"url"`
	VarScope map[string]Value                           `json:"var_scope"`
	FnScope  map[string]interpreter.FunctionDeclaration `json:"fn_scope"`
	Name     string                                     `json:"name"`
	Params   []Value                                    `json:"values"`
}

type intermediateRequest struct {
	URL      string           `json:"url"`
	VarScope map[string]Value `json:"var_scope"`
	Name     string           `json:"name"`
	Params   []Value          `json:"values"`
}

func (req *ExecRequest) UnmarshalJSON(b []byte) (err error) {

	ir := intermediateRequest{}
	err = json.Unmarshal(b, &ir)
	if err != nil {
		return
	}

	req.URL = ir.URL
	req.VarScope = ir.VarScope
	req.Name = ir.Name
	req.Params = ir.Params

	data := map[string]interface{}{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return
	}

	req.FnScope = map[string]interpreter.FunctionDeclaration{}
	fnScopeData, ok := data["fn_scope"]
	if !ok {
		return
	}

	fnScope, ok := fnScopeData.(map[string]interface{})
	if !ok {
		err = errors.New("failed to parse fn scope")
		return
	}

	for k, v := range fnScope {

		f, ok := v.(map[string]interface{})
		if !ok {
			err = errors.New("failed to parse fn scope")
			return
		}
		req.FnScope[k], err = parseFunctionDeclaration(f)
		if err != nil {

			return
		}
	}

	return
}

func parseFunctionDeclaration(f map[string]interface{}) (decl interpreter.FunctionDeclaration, err error) {
	n, ok := f["name"]
	if !ok {
		err = errors.New("funciton delcaration must have a name")
		return
	}

	decl.Name, ok = n.(string)
	if !ok {
		err = errors.New("funciton delcaration name must be string")
		return
	}

	p, ok := f["params"]
	if ok {
		decl.Params, ok = p.([]string)
		if !ok {
			err = errors.New("funciton delcaration params must be a string list")
			return
		}
	}

	def, ok := f["definition"]
	if !ok {
		err = errors.New("funciton delcaration must have definitions")
		return
	}

	exp, ok := def.(map[string]interface{})
	if !ok {
		err = errors.New("funciton delcaration must be a map")
		return
	}

	decl.Definition, err = parseExpression(exp)
	return

}
