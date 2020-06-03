package model

type Node struct {
	Type     string `json:"type" bson:"type"`         //[b|r|s], 链类型
	Number   string `json:"number" bson:"number"`     // 链编号
	KeyId    string `json:"keyId" bson:"keyId"`       // 链编号
	LastTime string `json:"lastTime" bson:"lastTime"` //最新的更新时间
	Id       string `json:"id" bson:"id"`             // 节点id
	Cmd      string   `json:"cmd" bson:"cmd"`
	CmdTime  string `json:"cmdTime" bson:"cmdTime"`
}

type NodeNum struct {
	Id    string `json:"id" bson:"_id"`
	Total int    `json:"total" bson:"total"`
}
