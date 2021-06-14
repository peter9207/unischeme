package interpreter

import (
	"errors"
	"fmt"
	"github.com/peter9207/unischeme/lexer"
)

var ErrUnknownValue = errors.New("unknown value type")
var ErrUnknownExpression = errors.New("unknown expression")
var ErrParametersMustBeIdentifiers = errors.New("parameters must be identifiers")
var ErrFnDeclarationWrongParameterCount = errors.New("function declaration should have 2 parameters")
var ErrInvalidFnDecl = errors.New("invalid fn declaration")

func ToProgram(expressions []lexer.Expression) (p Program, err error) {

	for _, e := range expressions {

		if e.FnCall != nil {

			fn := e.FnCall

			switch fn.Name.Name {
			case "main":
				t, err := parseFunctionCall(e.FnCall)
				if err != nil {
					return p, err
				}
				p.Main = t
			case "def":

				t, err := parseFunctionDeclaration(e.FnCall)
				if err != nil {
					return p, err
				}
				_, exists := p.Scope[t.Name]
				if exists {
					err = errors.New("function redeclared")
					return p, err
				}

				p.Scope[t.Name] = t

			}
			continue
		}

		if e.Identifier != nil {
			i, err := parseIdentifier(e.Identifier)
			if err != nil {
				return p, err
			}
			p.Scope[i.Name] = i
			continue
		}

		if e.Value != nil {
			err = errors.New("orphaed value")
			continue
		}
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

	if e.FnCall != nil {
		result, err = parseFunctionCall(e.FnCall)
		return
	}

	err = fmt.Errorf("unknown expression %T", e)
	return
}

func parseFunctionDeclaration(fn *lexer.FnCall) (f *FunctionDeclaration, err error) {
	f = &FunctionDeclaration{}
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

	block, err := parseExpression(fn.Parameters[1])
	blockExp, ok := block.(Expression)
	if !ok {
		err = errors.New("function block must be an expression")
	}

	f.Definition = blockExp

	return
}

func parseFunctionCall(fn *lexer.FnCall) (fnCall *FunctionCall, err error) {

	fnCall = &FunctionCall{
		Name: fn.Name.Name,
	}

	for _, exp := range fn.Parameters {

		a, err := parseExpression(exp)
		if err != nil {
			return nil, err
		}
		expression, ok := a.(Expression)
		if !ok {
			err = errors.New("parameters must be expressions")
			return nil, err
		}

		fnCall.Params = append(fnCall.Params, expression)
	}
	return
}

func parseIdentifier(v *lexer.Identifier) (node Identifier, err error) {
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
