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
type CmdInfo struct {
	Key    string `json:"key" bson:"key"`
	Params struct {
		Amount int `json:"amount" bson:"amount"`
	} `json:"params" bson:"params"`
}

