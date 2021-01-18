package model

import "time"

type Node struct {
	NodeId     string    `json:"nodeId" bson:"nodeId"`
	Type       string    `json:"type" bson:"type"`
	ChainKey   string    `json:"chainKey" bson:"chainKey"`
	LatestTime time.Time `json:"latestTime" bson:"latestTime"`
	Cmd        CmdInfo   `json:"cmd" bson:"cmd"`
	CmdTime    time.Time `json:"cmdTime" bson:"cmdTime"`
}

func NewEmptyCmdInfo() CmdInfo {
	return CmdInfo{Key: "other", Params: Params{Amount: 0}}
}

type CmdInfo struct {
	Key    string `json:"key" bson:"key"`
	Params Params `json:"params" bson:"params"`
}

type Params struct {
	Amount int `json:"amount" bson:"amount"`
}
