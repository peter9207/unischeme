package interpreter

import (
	"encoding/json"
	"fmt"
)

type IntValue struct {
	Value int `json:"value"`
}

func (f IntValue) MarshalJSON() (data []byte, err error) {

	i := map[string]interface{}{
		"value": f.Value,
		"type":  f.Type(),
	}

	data, err = json.Marshal(i)
	return
}

func (f IntValue) UnmarshalJSON() (data []byte, err error) {

	i := map[string]interface{}{}

	err = json.Unmarshal(data, &i)
	if err != nil {
		return
	}

	return
}
func (i IntValue) Node() interface{} {
	return i
}

func (i IntValue) Type() string {
	return "intValue"
}

func (i IntValue) String() string {
	return fmt.Sprintf("%v", i.Value)
}

func (i IntValue) Resolve(_ map[string]Expression) (Value, error) {
	return i, nil
}
