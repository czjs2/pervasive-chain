package model

type TotalFlow struct {
	Id   string `json:"id" bson:"_id"`
	Time string `json:"time" bson:"time"`
	In   int    `json:"in" bson:"in"`
	Out  int    `json:"out" bson:"out"`
}
