package interpreter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/peter9207/unischeme/interpreter"
	"github.com/peter9207/unischeme/lexer"
)

var _ = Describe("Execute", func() {

	Describe("can execute simple functions", func() {

		It("can parse integer values", func() {
			program, err := lexer.Parse("5")
			Ω(err).Should(BeNil())
			nodes, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())
			Ω(len(nodes)).Should(Equal(1))

			n := nodes[0]
			Ω(n.Type()).Should(Equal("intValue"))
			intNode, ok := n.(interpreter.IntValue)
			Ω(ok).Should(Equal(true))
			Ω(intNode.Value).Should(Equal(5))
		})

		It("can parse string values", func() {
			program, err := lexer.Parse(`"some string"`)
			Ω(err).Should(BeNil())
			nodes, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())
			Ω(len(nodes)).Should(Equal(1))

			n := nodes[0]
			Ω(n.Type()).Should(Equal("stringValue"))
			intNode, ok := n.(interpreter.StringValue)
			Ω(ok).Should(Equal(true))
			Ω(intNode.Value).Should(Equal("some string"))
		})

		It("can parse function delcarations", func() {
			program, err := lexer.Parse(`(def (foo i) 5)`)
			Ω(err).Should(BeNil())
			nodes, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())
			Ω(len(nodes)).Should(Equal(1))

			n := nodes[0]
			Ω(n.Type()).Should(Equal("functionDeclaration"))
			fnNode, ok := n.(interpreter.FunctionDeclaration)
			Ω(ok).Should(Equal(true))
			Ω(fnNode.Name).Should(Equal("foo"))
			Ω(len(fnNode.Params)).Should(Equal(1))
			p := fnNode.Params[0]
			Ω(p).Should(Equal("i"))
		})

		It("can parse function calls", func() {
			program, err := lexer.Parse(`(foo i)`)
			Ω(err).Should(BeNil())
			nodes, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())
			Ω(len(nodes)).Should(Equal(1))
			n := nodes[0]
			Ω(n.Type()).Should(Equal("functionCall"))
			fnNode, ok := n.(interpreter.FunctionCall)
			Ω(ok).Should(Equal(true))
			Ω(fnNode.Name).Should(Equal("foo"))
			Ω(len(fnNode.Params)).Should(Equal(1))
			p := fnNode.Params[0]
			param, ok := p.(interpreter.Identifier)
			Ω(ok).Should(Equal(true))
			Ω(param.Name).Should(Equal("i"))
		})
	})

})
