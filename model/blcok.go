package model

import "time"

type Block struct {
	Type     string      `form:"type" binding:"required"`     //[b|r|s], 链类型
	Number   string      `form:"number" binding:"required"`   // 链编号
	Id       string      `form:"id" binding:"required"`       // 节点id
	Height   int         `form:"height" binding:"required"`   //当前区块高度
	Father   string      `form:"father" binding:"required"`   //父区块hash
	Hash     string      `form:"hash" binding:"required"`     //区块hash
	Vrf      string      `form:"vrf" binding:"required"`      //VRF
	Time     time.Time       `form:"time" binding:"required"`     //当前产生时间
	Interval int         `form:"interval" binding:"required"` //出块间隔
	Trans    int         `form:"trans" binding:"required"`    //交易数量
	Size     int         `form:"size" binding:"required"`     //区块大小
	Detail   interface{} `form:"detail" binding:"required"`   //详情 (需详细定义)
}

type ChainTps struct {
	Tps int `json:"tps" bson:"tps"`
}

type NodeBlock struct {
	Block *Block       `json:"block"`
	Child []*NodeBlock `json:"child"` // 该链的子链
}

type Block1 struct {
	Id BlockType `json:"id" bson:"_id"`
}

type BlockType struct {
	Id   string `json:"id" bson:"id"`
	Type string `json:"type" bson:"type"`
}
