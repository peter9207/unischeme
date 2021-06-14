package server

import (
	"github.com/peter9207/unischeme/interpreter"
)

type InterpretRequest struct {
	URL  string                 `json:"url"`
	Body interpreter.Expression `json:"body"`
}

func (i *InterpretRequest) UmarshalJSON(data []byte) (err error) {
	return
}
