package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"github.com/peter9207/unischeme/server"
)

var _ = Describe("ExecRequest", func() {

	Describe("can serialize and deserialize empty scopes", func() {
		url := "testurl"
		varScope := map[string]interpreter.Value{}
		fnScope := map[string]interpreter.Expression{}

		It("should produce the same result", func() {
			req, err := server.MakeExecRequest(url, varScope, fnScope)
			Ω(err).Should(BeNil())

			data, err := json.Marshal(req)
			Ω(err).Should(BeNil())

			actual := server.ExecRequest{}
			err = json.Unmarshal(data, &actual)
			Ω(err).Should(BeNil())
			Ω(actual).Should(Equal(req))
		})
	})

	Describe("can serialize and deserialize value scopes", func() {
		url := "testurl"
		varScope := map[string]interpreter.Value{
			"f": interpret.StringValue{
				Value: "some string",
			},
			"g": interpret.IntValue{
				Value: "some string",
			},
		}
		fnScope := map[string]interpreter.Expression{}

		It("should produce the same result", func() {
			req, err := server.MakeExecRequest(url, varScope, fnScope)
			Ω(err).Should(BeNil())

			data, err := json.Marshal(req)
			Ω(err).Should(BeNil())

			actual := server.ExecRequest{}
			err = json.Unmarshal(data, &actual)
			Ω(err).Should(BeNil())
			Ω(actual).Should(Equal(req))
		})
	})

	Describe("can serialize and deserialize value scopes", func() {
		url := "testurl"
		varScope := map[string]interpreter.Value{}
		fnScope := map[string]interpreter.Expression{}

		It("should produce the same result", func() {
			req, err := server.MakeExecRequest(url, varScope, fnScope)
			Ω(err).Should(BeNil())

			data, err := json.Marshal(req)
			Ω(err).Should(BeNil())

			actual := server.ExecRequest{}
			err = json.Unmarshal(data, &actual)
			Ω(err).Should(BeNil())
			Ω(actual).Should(Equal(req))
		})
	})

})
