package interpreter

import "github.com/peter9207/unischeme/lexer"

func Exec(program lexer.Program) (result int, err error) {

	for _, e := range program.Expressions {

		if e.Value != nil {
			result = parseValue(e.Value)
			return
		}
	}
	return
}

func parseValue(v *lexer.Value) (value int) {

	if v.Int != nil {
		return *v.Int
	}

	return -1
}
