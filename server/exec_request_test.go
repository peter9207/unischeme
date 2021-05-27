package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"fmt"
	"github.com/peter9207/unischeme/interpreter"
	"github.com/peter9207/unischeme/server"
)

var _ = Describe("ExecRequest", func() {

	Describe("can serialize and deserialize empty scopes", func() {
		name := "testName"
		url := "testurl"
		varScope := map[string]interpreter.Value{}
		fnScope := map[string]interpreter.FunctionDeclaration{}
		params := []interpreter.Value{}

		It("should produce the same result", func() {
			req, err := server.MakeExecRequest(url, varScope, fnScope, name, params)
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
			"f": interpreter.StringValue{
				Value: "some string",
			},
			"g": interpreter.IntValue{
				Value: 5,
			},
		}

		fnScope := map[string]interpreter.FunctionDeclaration{}
		params := []interpreter.Value{}
		name := "testName"

		It("should produce the same result", func() {
			req, err := server.MakeExecRequest(url, varScope, fnScope, name, params)
			Ω(err).Should(BeNil())

			data, err := json.Marshal(req)
			Ω(err).Should(BeNil())

			fmt.Println(string(data))
			fmt.Println(varScope)

			actual := server.ExecRequest{}
			err = json.Unmarshal(data, &actual)
			Ω(err).Should(BeNil())
			Ω(actual).Should(Equal(req))
		})
	})

	Describe("can serialize and deserialize function scope", func() {
		url := "testurl"
		varScope := map[string]interpreter.Value{}
		fnScope := map[string]interpreter.FunctionDeclaration{
			"foo": interpreter.FunctionDeclaration{
				Name:       "testname",
				Params:     []string{"a", "b"},
				Definition: interpreter.IntValue{Value: 5},
			},
		}
		params := []interpreter.Value{}
		name := "testName"

		It("should produce the same result", func() {
			req, err := server.MakeExecRequest(url, varScope, fnScope, name, params)
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
