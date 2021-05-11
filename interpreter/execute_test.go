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

	})

})
