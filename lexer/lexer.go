package lexer

import (
	"github.com/alecthomas/participle"
)

type Program struct {
	Expressions []Expression `@@*`
}

type Expression struct {
	FnCall     *FnCall     `@@`
	Value      *Value      `| @@`
	Identifier *Identifier `| @@`
}

type FnCall struct {
	Name       Identifier   `"(" @@`
	Parameters []Expression `@@* ")"`
}

type Value struct {
	String *string  `@String`
	Float  *float64 `| @Float`
	Int    *int     `| @Int`
}

type Identifier struct {
	Name string `@Ident`
}

func Parse(data string) (p Program, err error) {
	parser, err := participle.Build(&p)
	if err != nil {
		return
	}

	err = parser.ParseString(data, &p)
	return
}
