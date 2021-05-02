package interpreter

import "github.com/peter9207/unischeme/lexer"
import "errors"
import "fmt"

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

type FunctionCall struct {
	Name       string
	Parameters []ASTNode
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
var ErrFnDeclarationWrongParameterCount = errors.New("function declaration should have 2 parameters")

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

var ErrInvalidFnDecl = errors.New("invalid fn declaration")

func parseFunctionCall(fn *lexer.FnCall) (node ASTNode, err error) {
	fmt.Println("start parsing for funciton call")

	if fn.Name.Name == "def" {
		if len(fn.Parameters) != 2 {
			err = ErrFnDeclarationWrongParameterCount
			return
		}

		def := fn.Parameters[0]

		if def.FnCall == nil {
			err = ErrInvalidFnDecl
			return
		}

		f := FunctionDeclaration{}
		f.Name = def.FnCall.Name.Name
		parameterList := []string{}
		for _, p := range def.FnCall.Parameters {
			fmt.Println(p)
			if p.Identifier == nil {
				err = ErrParametersMustBeIdentifiers
				return
			}
			parameterList = append(parameterList, p.Identifier.Name)
		}
		f.Params = parameterList

		fmt.Println("got here")
		block := fn.Parameters[1]
		f.Definition, err = parseExpression(block)
		if err != nil {
			return
		}
		node = f
		return
	}

	fnCall := FunctionCall{
		Name: fn.Name.Name,
	}

	for _, exp := range fn.Parameters {

		var n ASTNode
		n, err = parseExpression(exp)
		if err != nil {
			return
		}

		fnCall.Parameters = append(fnCall.Parameters, n)
	}

	node = fnCall
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
