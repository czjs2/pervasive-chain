package model

import (
	"encoding/json"
	"time"
)

type TotalFlow struct {
	Id   string `json:"id" bson:"_id"`
	Time time.Time  `json:"time" bson:"time"`
	In   int    `json:"in" bson:"in"`
	Out  int    `json:"out" bson:"out"`
}

func (c TotalFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(ObjToMap(c))
}
