package lexer

import (
	"github.com/alecthomas/participle"
)

type Program struct {
	Expressions []Expression `@@*`
}

type Expression struct {
	// Name       string   `"(" @Ident`
	// Parameters []Values `@@ ")"`
	Value  *Value  `@@`
	FnCall *FnCall `| @@`
}

type FnCall struct {
	Name       string  `"(" @Ident`
	Parameters []Value `@@* ")"`
}

type Value struct {
	String *string  `@String`
	Float  *float64 `| @Float`
	Int    *int     `| @Int`
}

func Parse(data string) (p Program, err error) {
	parser, err := participle.Build(&p)
	if err != nil {
		return
	}

	err = parser.ParseString(data, &p)
	return
}
