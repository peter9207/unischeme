package server_test

//build +integrations

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/peter9207/unischeme/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("ExecRequest", func() {

	testName := "testServer"
	url := "http://localhost:1234"
	server := server.New(testName, url)

	It("handles health checks", func() {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/ping", nil)
		Ω(err).Should(BeNil())
		server.Router.ServeHTTP(w, req)
		Ω(w.Code).Should(Equal(200))
	})

	It("handles simple requests", func() {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/do", strings.NewReader("5"))
		Ω(err).Should(BeNil())
		server.Router.ServeHTTP(w, req)
		Ω(w.Code).Should(Equal(200))

		data, err := ioutil.ReadAll(w.Body)
		Ω(string(data)).Should(Equal(5))

	})

})
