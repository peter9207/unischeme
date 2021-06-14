package interpreter

import (
	"encoding/json"
)

type StringValue struct {
	Value string `json:"value"`
}

func (f StringValue) MarshalJSON() (data []byte, err error) {

	i := map[string]interface{}{
		"value": f.Value,
		"type":  f.Type(),
	}

	data, err = json.Marshal(i)

	return
}

func (i StringValue) Node() interface{} {
	return i
}

func (i StringValue) Type() string {
	return "stringValue"
}

func (i StringValue) String() string {
	return i.Value
}

func (i StringValue) Resolve(_ map[string]Expression) (Value, error) {
	return i, nil
}
