package block

import (
	"pervasive-chain/config"
	"pervasive-chain/utils"
)

type ReportBlockForm struct {
	Type     string `form:"type" binding:"required"`     //[b|r|s], 链类型
	ChainKey string `form:"chainKey" binding:"required"` // 链编号
	NodeId   string `form:"nodeId" binding:"required"`   // 节点id
	Height   int64  `form:"height" binding:"required"`   //当前区块高度
	Father   string `form:"father" binding:"required"`   //父区块hash
	Hash     string `form:"hash" binding:"required"`     //区块hash
	Vrf      string `form:"vrf" binding:"omitempty"`     //VRF
	Time     string `form:"time" binding:"required"`     //当前产生时间
	Interval int64  `form:"interval" binding:"required"` //出块间隔
	Trans    int64  `form:"trans" binding:"required"`    //交易数量
	Size     int64  `form:"size" binding:"required"`     //区块大小

	LockHash []LockHash `form:"lockHash" binding:"required"`

	UpHash   string      `form:"upHash" binding:"required"`
	DownHash string      `form:"downHash" binding:"required"`
	Detail   DetailBlock `form:"detail" binding:"required"`
}

type LockHash struct {
	ChainKey string `json:"chainKey"`
	Height   int    `json:"height"`
}

type DetailBlock struct {
	UpStream   []RelayTransGroup `json:"upStream"`
	DownStream []RelayTransGroup `json:"downStream"`
	Ss         []TransGroup      `json:"ss"`
}

type RelayTransGroup struct {
	Key  string `json:"key"`
	Hash string `json:"hash"`
}

type TransGroup struct {
	FromShard string `json:"fromShard"`
	ToShard   string `json:"toShard"`
	FromRelay string `json:"fromRelay"`
	ToRelay   string `json:"toRelay"`
	Hash      string `json:"hash"`
	Trans     []struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount string `json:"amount"`
	} `json:"trans"`
}

func (h *ReportBlockForm) Valid() (bool, error) {
	if !utils.IsValidChain(h.Type) {
		return false, nil
	}
	if !utils.IsValidChainKey(h.ChainKey) {
		return false, nil
	}
	if len(h.NodeId) == 0 {
		return false, nil
	}
	if h.Height < 0 {
		return false, nil
	}
	return h.checkParam(), nil
}

func (h *ReportBlockForm) checkParam() bool {
	switch h.Type {
	case config.BeaconType:
		if len(h.DownHash) == 0 {
			return false
		}
	case config.RelayType:
		if len(h.DownHash) == 0 {
			return false
		}
		if len(h.UpHash) == 0 {
			return false
		}
	case config.SharedType:
		if len(h.UpHash) == 0 {
			return false
		}
	default:
		return true
	}
	return true
}
