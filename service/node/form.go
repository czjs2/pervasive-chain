package node

import (
	"fmt"
	"pervasive-chain/utils"
)

type HeartBeatFrom struct {
	Type     string `form:"type" binding:"required"`     //[b|r|s], 链类型
	ChainKey string `form:"chainKey" binding:"required"` // 链编号
	NodeId   string `form:"nodeId" binding:"required"`   // 节点id
	Time     string `form:"time" binding:"required"`     // 时间

}

type GenCmdFrom struct {
	Type string `json:"type"`
	Cmd  struct {
		Key    string `json:"key"`
		Params struct {
			Amount int `json:"amount"`
		} `json:"params"`
	} `json:"cmd"`
}

func (g *GenCmdFrom) Valid() (bool, error) {
	if !utils.IsValidChain(g.Type) {
		return false, fmt.Errorf("type is error %v \n", g.Type)
	}
	if g.Cmd.Key != "transfer" {
		return false, fmt.Errorf("key value is not transfer %v \n", g.Cmd.Key)
	}
	if g.Cmd.Params.Amount == 0 {
		return false, fmt.Errorf("amount is zero  \n")
	}
	return true, nil
}

func (h *HeartBeatFrom) Valid() (bool, error) {
	if !utils.IsValidChain(h.Type) {
		return false, fmt.Errorf("chain type is error %v \n", h.Type)
	}
	if !utils.IsValidChainKey(h.ChainKey, h.Type) {
		return false, fmt.Errorf("chainKey is error %v \n", h.ChainKey)
	}
	if !utils.IsValidNodeId(h.NodeId) {
		return false, fmt.Errorf("nodeId is error %v \n", h.NodeId)
	}
	if !utils.IsRFC339Time(h.Time) {
		return false, fmt.Errorf("time is error %v \n", h.Time)
	}
	return true, nil
}
