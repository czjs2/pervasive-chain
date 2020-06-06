package ws

import (
	"fmt"
	"pervasive-chain/config"
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
		return d.DoBlockInfo(cmd)
	case ChainInfoCmd:
		return d.DoChainInfo(cmd)
	case ExecuteCmd:
		return d.GenCmd(cmd)
	default:
		return nil, fmt.Errorf("%s unsupport error  ", cmd.Uri)
	}
}

func (d *Dispatch) DoBlockInfo(cmd model.Cmd) ([]byte, error) {
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

func (d *Dispatch) DoChainInfo(cmd model.Cmd) ([]byte, error) {

	blockService := service.NewBlockService()
	// 信标链 最新区块列表
	blockList, _, err := blockService.BlockList(config.BChain, "")
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	totalFlowService := service.NewTotalFlowService()
	// 总带宽
	totalFlowList, _, err := totalFlowService.FlowList()
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	statisticService := service.NewStatisticService()
	// 节点总数
	countNode, err := statisticService.CountNode()
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	// 各种链信息
	allChain, err := statisticService.AllChain()
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	p := model.P{
		"beaconBlockList": blockList,
		"totalFlow":       totalFlowList,
		"countNode":       countNode,
		"chainList":       allChain,
	}
	return NewSuccessResp(cmd, p)
}
