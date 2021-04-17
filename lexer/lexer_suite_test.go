package lexer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/peter9207/unischeme/lexer"
)

func TestLexer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lexer Suite")
}

var _ = Describe("programs with 1 expression", func() {
	Describe("parsing values", func() {
		p, err := lexer.Parse("5")
		It("should not return errors", func() {
			Ω(err).Should(BeNil())
		})

		It("should parse the correct amount of expressiosn", func() {
			Ω(len(p.Expressions)).Should(Equal(1))
		})

		It("Should parse the correct value type", func() {
			exp := p.Expressions[0]
			Ω(exp.Value.Int).ShouldNot(BeNil())
			Ω(*exp.Value.Int).Should(BeEquivalentTo(5))
		})

	})
	Describe("function with 1 parameter", func() {
		p, err := lexer.Parse("(foo 1)")
		It("should not return errors", func() {
			Ω(err).Should(BeNil())
		})
		It("should parse the correct amount of expressiosn", func() {
			Ω(len(p.Expressions)).Should(Equal(1))
		})

		It("Should parse the correct values", func() {
			exp := p.Expressions[0]
			Ω(exp.FnCall.Name).Should(Equal("foo"))

			Ω(len(exp.FnCall.Parameters)).Should(Equal(1))
			Ω(*exp.FnCall.Parameters[0].Int).Should(Equal(1))
		})

	})
})
