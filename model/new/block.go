package new

import "time"

type NewBlock struct {
	Type         string      `json:"type"`     //[b|r|s], 链类型
	Number       string      `json:"number"`   // 链编号
	Id           string      `json:"id"`       // 节点id
	Height       int         `json:"height"`   //当前区块高度
	Father       string      `json:"father"`   //父区块hash
	Hash         string      `json:"hash"`     //区块hash
	Vrf          string      `json:"vrf"`      //VRF
	Time         time.Time   `json:"time"`     //当前产生时间
	Interval     int         `json:"interval"` //出块间隔
	Trans        int         `json:"trans"`    //交易数量
	Size         int         `json:"size"`     //区块大小
	Detail       interface{} `json:"detail"`   //详情 (需详细定义)
	LockBlocks   []LockBlock `json:"lockBlocks"`
	CrossTrans   []GroupInfo `json:"crossTrans"`
	ConfirmTrans []GroupInfo `json:"confirmTrans"`
}

type GroupInfo struct {
	Key    string      `json:"key"`
	Hash   string      `json:"hash"`
	Trans  int64       `json:"trans"`
	Groups []GroupInfo `json:"groups"`
}

type LockBlock struct {
	Type   string `json:"type"`
	Number string `json:"number"`
	Height int64  `json:"height"`
}

type Trans struct {
	From      string
	FromShard string `json:"fromShard"`//链编号
	To        string
	ToShard   string `json:"toShard"`//链编号
	Amount    string
}
