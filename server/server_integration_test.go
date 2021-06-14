package server_test

import (
	// "bytes"
	// "encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/peter9207/unischeme/interpreter"
	// "github.com/peter9207/unischeme/lexer"
	"github.com/peter9207/unischeme/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("ExecRequest", func() {

	testName := "testDest"
	url := "http://localhost:1234"
	dest := server.New(testName, url)

	XIt("handles health checks", func() {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/ping", nil)
		Ω(err).Should(BeNil())
		dest.Router.ServeHTTP(w, req)
		Ω(w.Code).Should(Equal(200))
	})

	It("handles simple requests", func() {
		w := httptest.NewRecorder()

		program := "(main 5)"

		req, err := http.NewRequest("POST", "/interpret", strings.NewReader(program))
		Ω(err).Should(BeNil())
		dest.Router.ServeHTTP(w, req)

		body, err := ioutil.ReadAll(w.Body)
		Ω(string(body)).Should(Equal(5))
		Ω(w.Code).Should(Equal(200))

	})

})
