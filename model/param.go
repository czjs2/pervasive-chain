package model

import (
	"encoding/json"
	"reflect"
)

type Param map[string]interface{}

func (u Param) MarshalJSON() ([]byte, error) {
	param := make(map[string]interface{})
	for k, v := range u {
		tk := k
		tv := v
		if reflect.TypeOf(v).Name() == "Time" {
			// todo
		}
		param[tk] = tv
	}
	return json.Marshal(param)
}
