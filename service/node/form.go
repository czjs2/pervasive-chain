package node

import (
	"pervasive-chain/utils"
)

type HeartBeatFrom struct {
	Type     string `form:"type" binding:"required"`     //[b|r|s], 链类型
	ChainKey string `form:"chainKey" binding:"required"` // 链编号
	NodeId   string `form:"nodeId" binding:"required"`   // 节点id
	Time     string `form:"time" binding:"required"`     // 时间

}

//type GenCmdFrom struct {
//	Type string `json:"type"`
//	Cmd  struct {
//		Key    string `json:"key"`
//		Params struct {
//			Amount int `json:"amount"`
//		} `json:"params"`
//	} `json:"cmd"`
//}
//
//func (g *GenCmdFrom) Valid() (bool, error) {
//	if g.Cmd.Params.Amount == 0 {
//		return false, nil
//	}
//	return true, nil
//}

func (h *HeartBeatFrom) Valid() (bool, error) {
	if !utils.IsValidChain(h.Type) {
		return false, nil
	}
	if !utils.IsValidChainKey(h.ChainKey) {
		return false, nil
	}
	if len(h.NodeId) == 0 {
		return false, nil
	}
	if len(h.Type) == 0 {
		return false, nil
	}
	if len(h.Time) == 0 {
		return false, nil
	}
	return true, nil
}
