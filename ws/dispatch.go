package ws

import (
	"fmt"
	"pervasive-chain/config"
	"pervasive-chain/model"
	"pervasive-chain/service"
	"time"
)

type Dispatch struct {
}

func NewDisPatch() *Dispatch {
	return &Dispatch{}
}

func (d *Dispatch) Execute(cmd model.Cmd) ([]byte, error) {
	switch cmd.Uri {
	case BlockInfoCmd:
		return d.WsBlockInfo(cmd)
	case ChainInfoCmd:
		return d.WsChainInfo(cmd)
	case ExecuteCmd:
		return d.GenCmd(cmd)
	case ChainInfoById:
		return d.ChainInfoById(cmd)
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
	nodeCmd, err := nodeService.LatestNodeCmd()
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	if time.Now().Sub(nodeCmd.CmdTime) < config.HeartBeatTime {
		return NewRespErr(cmd, "前一个命令还在下发中...")
	}
	_, total, err := nodeService.OnLineList()
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	if len(cmd.Body.Cmd.Params) == 0 {
		return NewRespErr(cmd, "参数为空")
	}
	totalTrans := cmd.Body.Cmd.Params[0]
	singTrans := totalTrans / float64(total)
	cmd.Body.Cmd.Params = []float64{singTrans}
	_, err = nodeService.UpdateOnLineNodeCmd(cmd.Body.Cmd)
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	return NewSuccessResp(cmd, nil)

}

func (d *Dispatch) WsBlockInfo(cmd model.Cmd) ([]byte, error) {
	statisticService := service.NewStatisticService()
	info, err := statisticService.BlockInfo(cmd.Body.Type, cmd.Body.Number)
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	return NewSuccessResp(cmd, info)
}

func (d *Dispatch) WsChainInfo(cmd model.Cmd) ([]byte, error) {
	statisticService := service.NewStatisticService()
	chainInfo, err := statisticService.ChainInfo()
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	return NewSuccessResp(cmd, chainInfo)
}

func (d *Dispatch) DoChainInfo(cmd model.Cmd) ([]byte, error) {

	// 获取每条链上最新的一个区块数据

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

// 指定链前 100个区块
func (d *Dispatch) ChainInfoById(cmd model.Cmd) ([]byte, error) {
	blockService := service.NewBlockService()
	list, _, err := blockService.BlockList(cmd.Body.Type, string(cmd.Body.Number))
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	return NewSuccessResp(cmd, list)
}
