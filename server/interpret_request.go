package server

import (
	"github.com/peter9207/unischeme/interpreter"
	"github.com/peter9207/unischeme/lexer"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"time"
)

type InterpretRequest struct {
	URL  string              `json:"url"`
	Body interpreter.ASTNode `json:"body"`
}

func (i *InterpretRequest) UmarshalJSON(data []byte) (err error) {

}
