package model

import (
	"encoding/json"
	"time"
)

type TotalChain struct {
	RelayNum  int `json:"relayNum" bson:"relayNum"`
	SharedNum int `json:"sharedNum" bson:"sharedNum"`
	NodeNum   int `json:"nodeNum" bson:"nodeNum"`
	TotalNum  int `json:"totalNum" bson:"totalNum"`
}

type Chain struct {
	Type   string `json:"type" bson:"type"`     //[b|r|s], 链类型
	Number string `json:"number" bson:"number"` // 链编号
	Id     string `json:"id" bson:"id"`         // 节点id
	Time   time.Time  `json:"time" bson:"time"`     // 时间
}


func (c Chain) MarshalJSON() ([]byte, error) {
	return json.Marshal(ObjToMap(c))
}

// 链类型总数
type ChainType struct {
	Id    string `json:"id" bson:"_id"`
	Total int    `json:"total" bson:"total"`
}

type Tps struct {
	Id  string `json:"id" bson:"_id"`
	Tps int64  `json:"tps" bson:"tps"`
}
type TotalTps struct {
	Id  int64 `json:"id" bson:"_id"`
	Tps int64 `json:"tps" bson:"tps"`
}
