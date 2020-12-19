package block

import (
	"fmt"
	"pervasive-chain/config"
	"pervasive-chain/utils"
	"strings"
)

type SingleBlockForm struct {
	Type     string `json:"type"`
	ChainKey string `json:"chainKey"`
	Height   uint64 `json:"height"`
	Hash     string `json:"hash"`
}

func (b *SingleBlockForm) Valid() (bool, error) {
	if b.Hash != "" {
		return true, nil
	}
	if strings.HasPrefix(b.Type,config.BeaconType){
		return true,nil
	}
	if b.Type == "" || b.ChainKey == "" {
		return false, nil
	}
	return true, nil
}


type ReportBlockForm struct {
	Type     string     `form:"type" binding:"required"`     //[b|r|s], 链类型
	ChainKey string     `form:"chainKey"` // 链编号
	NodeId   string     `form:"nodeId" binding:"required"`   // 节点id
	Height   uint64     `form:"height"`                      //当前区块高度
	Father   string     `form:"father" binding:"required"`   //父区块hash
	Hash     string     `form:"hash" binding:"required"`     //区块hash
	Vrf      string     `form:"vrf" binding:"required"`      //VRF
	Time     string     `form:"time" binding:"required"`     //当前产生时间
	Interval float64     `form:"interval"` //出块间隔
	Trans    uint64     `form:"trans"`    //交易数量
	Size     uint64     `form:"size"`     //区块大小
	Gas      string     `form:"gas"`
	LockHash []LockHash `form:"lockHash" binding:"required"`

	UpHash   string      `form:"upHash"`
	DownHash string      `form:"downHash"`
	Detail   DetailBlock `form:"detail" binding:"required"`
}

type LockHash struct {
	Type     string `json:"type"`
	ChainKey string `json:"chainKey"`
	Height   uint64 `json:"height"`
}

type DetailBlock struct {
	UpStream   []RelayTransGroup `json:"upStream"`
	DownStream []RelayTransGroup `json:"downStream"`
	Ss         []TransGroup      `json:"ss"`
}

type RelayTransGroup struct {
	FromKey string `json:"fromKey"`
	ToKey   string `json:"toKey"`
	Hash    string `form:"hash" binding:"required"`
}

type TransGroup struct {
	FromShard string `json:"fromShard"`
	ToShard   string `json:"toShard"`
	FromRelay string `json:"fromRelay"`
	ToRelay   string `json:"toRelay"`
	Hash      string `form:"hash" binding:"required"`
	Trans     []struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int `json:"amount"`
		Hash   string `form:"hash" binding:"required"`
	} `json:"trans"`
}

func (h *ReportBlockForm) Valid() (bool, error) {
	if !utils.IsValidChain(h.Type) { // 效验 type类型 B R S
		return false, fmt.Errorf("chain type is error %v \n", h.Type)
	}
	if !utils.IsValidChainKey(h.ChainKey, h.Type) { // chainKey
		return false, fmt.Errorf("chainkey is error %v  %v \n", h.ChainKey, h.Type)
	}
	if !utils.IsValidNodeId(h.NodeId) {
		return false, fmt.Errorf("nodeId is error %v \n", h.NodeId)
	}
	if !utils.IsRFC339Time(h.Time) {
		return false, fmt.Errorf("time is error %v \n", h.Time)
	}
	if h.Type == config.SharedType && (h.UpHash == "" || len(h.Detail.UpStream) == 0 || len(h.Detail.Ss) == 0) {
		return false, fmt.Errorf("shard type params error,upHash,upstream,ss can not empty")
	}
	if h.Type == config.RelayType && (h.UpHash == "" || h.DownHash == "" || len(h.Detail.UpStream) == 0 || len(h.Detail.DownStream) == 0) {
		return false, fmt.Errorf("relay type params error ,upHash downHash upstream downstream, can not empty")
	}
	if h.Type == config.BeaconType && (h.DownHash == "" || len(h.Detail.DownStream) == 0) {
		return false, fmt.Errorf("beacon type params downhasah,downstream,can not empty")
	}
	return true, nil

}
