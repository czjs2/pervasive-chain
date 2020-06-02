package model

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
	Time   string `json:"time" bson:"time"`     // 时间
}
