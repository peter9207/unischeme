package interpreter

import "github.com/peter9207/unischeme/lexer"
import "errors"

var globalScope map[string]FunctionDeclaration

type ASTNode interface {
	Node() interface{}
	Type() string
	Children() []ASTNode
}

type AST struct {
	Children []AST
	Type     string
	Value    string
}

type Function struct {
	Name   string
	Params []string
	Tree   AST
}

type FunctionDeclaration struct {
	Name       string
	params     []string
	Definition AST
}

type Identifier struct {
	Name string
}

func Exec(program lexer.Program) (result int, err error) {

	for _, e := range program.Expressions {

		if e.Value != nil {
			result = parseValue(e.Value)
			return
		}

		if e.FnCall != nil {

		}
	}

	return
}

var ErrParametersMustBeValues = errors.New("parameters must be values")
var ErrUnknownValue = errors.New("unknown value type")
var ErrParametersMustBeIdentifiers = errors.New("parameters must be identifiers")

func resolveFunctionCall(fn lexer.FnCall) (err error) {

	if fn.Name.Name == "def" {
		if _, ok := globalScope[fn.Name.Name]; ok {
			err = errors.New("function already defined")
			return
		}

		parameterList := []string{}
		for i := 0; i < len(fn.Parameters); i++ {
			p := fn.Parameters[i]
			if p.Value == nil {
				err = ErrParametersMustBeValues
				return
			}
			if p.Value.String == nil {
				err = ErrParametersMustBeIdentifiers
				return
			}
			parameterList = append(parameterList, *p.Value.String)
		}

		// f := FunctionDeclaration{
		// 	Name: fn.Name.Name,
		// }

	}

	return
}

func parseValue(v *lexer.Value) (node ASTNode, err error) {

	if v.Int != nil {
		node = IntValue{
			Value: *v.Int,
		}
		return
	}

	if v.String != nil {
		node = StringValue{
			Value: *v.String,
		}
		return

	}
	err = ErrUnknownValue
	return
}
