package interpreter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/peter9207/unischeme/interpreter"
)

var _ = Describe("FunctionCall", func() {

	Describe("built in functions", func() {

		It("plus", func() {

			f := interpreter.FunctionCall{
				Name: "plus",
				Params: []interpreter.Expression{
					interpreter.IntValue{
						Value: 1,
					},
					interpreter.IntValue{
						Value: 1,
					},
				},
			}

			v, err := f.Resolve(map[string]interpreter.Expression{}, map[string]interpreter.FunctionDeclaration{})
			Ω(err).Should(BeNil())
			Ω(v.String()).Should(Equal("2"))
		})

	})

})
