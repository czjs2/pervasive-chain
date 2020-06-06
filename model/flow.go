package model

type TotalFlow struct {
	Id   string `json:"id" bson:"_id"`
	Time int64  `json:"time" bson:"time"`
	In   int    `json:"in" bson:"in"`
	Out  int    `json:"out" bson:"out"`
}
