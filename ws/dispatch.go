package ws

import (
	"fmt"
	"pervasive-chain/model"
	"pervasive-chain/service"
)

type Dispatch struct {
}

func NewDisPatch() *Dispatch {
	return &Dispatch{}
}

func (d *Dispatch) Execute(cmd model.Cmd) ([]byte, error) {
	switch cmd.Uri {
	case BlockInfoCmd:
		return d.doBlockInfo(cmd)
	case ChainInfoCmd:
		return d.doChainInfo(cmd)
	case ExecuteCmd:
		return d.GenCmd(cmd)
	default:
		return nil, fmt.Errorf("%s unsupport error  ", cmd.Uri)
	}
}

func (d *Dispatch) doBlockInfo(cmd model.Cmd) ([]byte, error) {
	blockService := service.NewBlockService()
	latestBlock, err := blockService.LatestBlock()
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	return NewSuccessResp(cmd, latestBlock)

}

// 生成命令
func (d *Dispatch) GenCmd(cmd model.Cmd) ([]byte, error) {
	nodeService := service.NewNodeService()
	_, err := nodeService.UpdateOnLineNodeCmd(cmd.Body)
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	return NewSuccessResp(cmd, nil)

}

func (d *Dispatch) doChainInfo(cmd model.Cmd) ([]byte, error) {
	// 总带宽
	totalFlowService := service.NewTotalFlowService()
	totalFlowList, _, err := totalFlowService.FlowList()
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	// 各种链的信息
	chainService := service.NewChainService()
	chainList, _, err := chainService.ChainList()
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	// 链的详细信息
	chain, err := chainService.Chain("")
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	// 链总体信息
	totalChainService := service.NewTotalChainService()
	totalChainList, _, err := totalChainService.TotalFlowList()
	p := model.P{
		"totalFlowList":  totalFlowList,
		"chainList":      chainList,
		"totalChainList": totalChainList,
		"chain":          chain,
	}
	return NewSuccessResp(cmd, p)
}
