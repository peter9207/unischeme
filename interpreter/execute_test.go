package interpreter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/peter9207/unischeme/interpreter"
	"github.com/peter9207/unischeme/lexer"
)

var _ = Describe("Simple Exec", func() {

	Describe("Simple values", func() {

		It("should parse function with just int values", func() {
			program, err := lexer.Parse("5")
			Ω(err).Should(BeNil())
			values, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())

			Ω(len(values)).Should(Equal(1))
			v := values[0]
			Ω(v).Should(Equal("5"))
		})

		It("should parse function with just string values", func() {
			program, err := lexer.Parse(`"some_string"`)
			Ω(err).Should(BeNil())
			values, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())

			Ω(len(values)).Should(Equal(1))
			v := values[0]
			Ω(v).Should(Equal("some_string"))
		})

		It("should parse function declarations", func() {
			program, err := lexer.Parse(`(def (a i) i) (a 5)`)
			Ω(err).Should(BeNil())
			values, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())

			Ω(len(values)).Should(Equal(1))
			v := values[0]
			Ω(v).Should(Equal("5"))
		})

	})

})
