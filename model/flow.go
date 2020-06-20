package model

import "time"

type TotalFlow struct {
	Id   string `json:"id" bson:"_id"`
	Time time.Time  `json:"time" bson:"time"`
	In   int    `json:"in" bson:"in"`
	Out  int    `json:"out" bson:"out"`
}
