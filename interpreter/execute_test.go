package interpreter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peter9207/unischeme/interpreter"
	"github.com/peter9207/unischeme/lexer"
)

var _ = Describe("Simple Exec", func() {

	Describe("Simple values", func() {
		It("can exec function delcarations", func() {
			program, err := lexer.Parse(`(def (foo i) 5) (foo 2)`)
			Ω(err).Should(BeNil())

			value, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())
			Ω(len(value)).Should(Equal(1))
			Ω(value[0]).Should(Equal("5"))

		})

		It("can exec function delcarations with simple values", func() {
			program, err := lexer.Parse(`(def (foo i) i) (foo 2)`)
			Ω(err).Should(BeNil())

			value, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())
			Ω(len(value)).Should(Equal(1))
			Ω(value[0]).Should(Equal("2"))

		})

		It("can exec function delcarations with simple values 2 functions", func() {
			program, err := lexer.Parse(`(def (foo i) i) (def (bar j) j) (foo (bar 2))`)

			Ω(err).Should(BeNil())
			value, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())
			Ω(len(value)).Should(Equal(1))
			Ω(value[0]).Should(Equal("2"))

		})

		It("can exec function delcarations with simple values 2 functions", func() {
			program, err := lexer.Parse(`(def (foo i) i) (def (bar j) j) (foo (bar 2))`)

			Ω(err).Should(BeNil())
			value, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())
			Ω(len(value)).Should(Equal(1))
			Ω(value[0]).Should(Equal("2"))

		})

	})

	Describe("simple builtin functions", func() {
		It("can exec function delcarations with simple values 2 functions", func() {
			program, err := lexer.Parse(`(def (foo i) (plus i 2)) (foo 2)`)

			Ω(err).Should(BeNil())
			value, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())
			Ω(len(value)).Should(Equal(1))
			Ω(value[0]).Should(Equal("4"))

		})

		It("add and subtract with functions", func() {
			program, err := lexer.Parse(`(def (foo i) (plus i 2)) (def (bar i) (subtract i 1)) (foo (bar 5))`)

			Ω(err).Should(BeNil())
			value, err := interpreter.Exec(program)
			Ω(err).Should(BeNil())
			Ω(len(value)).Should(Equal(1))
			Ω(value[0]).Should(Equal("6"))

		})

	})

})
