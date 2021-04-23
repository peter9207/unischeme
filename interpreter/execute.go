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

func Exec(program lexer.Program) (result []ASTNode, err error) {
	for _, e := range program.Expressions {
		var t ASTNode
		t, err = parseExpression(e)
		if err != nil {
			return
		}
		result = append(result, t)
	}
	return
}

var ErrParametersMustBeValues = errors.New("parameters must be values")
var ErrUnknownValue = errors.New("unknown value type")
var ErrUnknownExpression = errors.New("unknown expression")
var ErrParametersMustBeIdentifiers = errors.New("parameters must be identifiers")

func parseExpression(e lexer.Expression) (result ASTNode, err error) {
	if e.Value != nil {
		result, err = parseValue(e.Value)
		return
	}
	if e.Identifier != nil {
		result, err = parseIdentifier(e.Identifier)
		return
	}
	if e.FnCall != nil {
		result, err = parseFunctionCall(e.FnCall)
		return
	}
	err = ErrUnknownExpression
	return
}

func parseFunctionCall(fn *lexer.FnCall) (node ASTNode, err error) {

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
	}
	return
}

func parseIdentifier(v *lexer.Identifier) (node ASTNode, err error) {

	node = Identifier{
		Name: v.Name,
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
