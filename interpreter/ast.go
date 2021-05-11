package interpreter

import "github.com/peter9207/unischeme/lexer"
import "errors"

var ErrUnknownValue = errors.New("unknown value type")
var ErrUnknownExpression = errors.New("unknown expression")
var ErrParametersMustBeIdentifiers = errors.New("parameters must be identifiers")
var ErrFnDeclarationWrongParameterCount = errors.New("function declaration should have 2 parameters")
var ErrInvalidFnDecl = errors.New("invalid fn declaration")

func ToAST(expressions []lexer.Expression) (result []ASTNode, err error) {
	for _, e := range expressions {
		var t ASTNode

		if e.FnCall != nil {
			// if e.FnCall != nil {
			// 	result, err = parseFunctionCall(e.FnCall)
			// 	return
			// }
			continue
		}

		t, err = parseExpression(e)
		if err != nil {
			return
		}
		result = append(result, t)
	}
	return
}

func parseExpression(e lexer.Expression) (result Expression, err error) {
	if e.Value != nil {
		result, err = parseValue(e.Value)
		return
	}
	if e.Identifier != nil {
		result, err = parseIdentifier(e.Identifier)
		return
	}

	err = ErrUnknownExpression
	return
}

func parseFunctionDeclaration(fn *lexer.FnCall) (f FunctionDeclaration, err error) {
	if len(fn.Parameters) != 2 {
		err = ErrFnDeclarationWrongParameterCount
		return
	}

	def := fn.Parameters[0]

	if def.FnCall == nil {
		err = ErrInvalidFnDecl
		return
	}

	f.Name = def.FnCall.Name.Name
	parameterList := []string{}
	for _, p := range def.FnCall.Parameters {
		if p.Identifier == nil {
			err = ErrParametersMustBeIdentifiers
			return
		}
		parameterList = append(parameterList, p.Identifier.Name)
	}
	f.Params = parameterList

	block := fn.Parameters[1]
	f.Definition, err = parseExpression(block)
	// f.Definition = block

	return
}

func parseFunctionCall(fn *lexer.FnCall) (node ASTNode, err error) {
	if fn.Name.Name == "def" {
		node, err = parseFunctionDeclaration(fn)
		return
	}

	fnCall := FunctionCall{
		Name: fn.Name.Name,
	}

	for _, exp := range fn.Parameters {

		var n Expression
		n, err = parseExpression(exp)
		if err != nil {
			return
		}

		fnCall.Params = append(fnCall.Params, n)
	}

	node = fnCall
	return
}

func parseIdentifier(v *lexer.Identifier) (node Expression, err error) {
	node = Identifier{
		Name: v.Name,
	}
	return
}

func parseValue(v *lexer.Value) (node Expression, err error) {

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
