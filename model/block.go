package model

import "time"

type LatestBlock struct {
	Type     string    `json:"type" bson:"type"`
	ChainKey string    `json:"chainKey" bson:"chain_key"`
	Height   int       `json:"height" bson:"height"`
	Time     time.Time `json:"time" bson:"time"`
	Interval int       `json:"interval" bson:"interval"`
	Trans    int       `json:"trans" bson:"trans"`
	Tps      int       `json:"tps" bson:"tps"`
	Size     int       `json:"size" bson:"size"`
}

type Node struct {
	NodeId     string    `json:"nodeId"`
	Type       string    `json:"type"`
	ChainKey   string    `json:"chainKey"`
	LatestTime time.Time `json:"latestTime"`
	Cmd        CmdInfo   `json:"cmd"`
	CmdTime    time.Time `json:"cmdTime"`
}
type CmdInfo struct {
	Key    string `json:"key"`
	Params struct {
		Value int `json:"value"`
	} `json:"params"`
}
