package model

import "time"

type LatestBlock struct {
}

type Node struct {
	/*
	       nodeId:String,//id为节点id
	   	type:String, //类型[B|R|S]
	   	chainKey:String,//链编号 中继为 FF  分片为 FFFF
	   	lastTime:Date,//最近一次上报心跳时间
	   	cmd:{key:String,params:Object},//命令 :{执行码，参数}
	   	cmdTime:Date, //命令产生时间
	*/
	NodeId     string    `json:"nodeId"`
	Type       string    `json:"type"`
	ChainKey   string    `json:"chainKey"`
	LatestTime string    `json:"latestTime"`
	Cmd        CmdInfo   `json:"cmd"`
	CmdTime    time.Time `json:"cmdTime"`
}
type CmdInfo struct {
	Key    string `json:"key"`
	Params struct {
		Value int `json:"value"`
	} `json:"params"`
}
