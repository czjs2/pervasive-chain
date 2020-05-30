package ws

import (
	"fmt"
	"pervasive-chain/service"
)

type Dispatch struct {
}

func NewDisPatch() *Dispatch {
	return &Dispatch{}
}

func (d *Dispatch) Execute(cmd Cmd) ([]byte, error) {
	switch cmd.Uri {
	case BlockInfoCmd:
		return d.doBlockInfo(cmd)
	case ChainInfoCmd:
		return d.doChainInfo(cmd)
	default:
		return nil, fmt.Errorf("%s unsupport error  ", cmd.Uri)
	}
}

func (d *Dispatch) doBlockInfo(cmd Cmd) ([]byte, error) {
	blockService := service.NewBlockService()
	block, err := blockService.LatestBlock()
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	return NewSuccessResp(cmd, block)
}

func (d *Dispatch) doChainInfo(cmd Cmd) ([]byte, error) {
	nodeService := service.NewNodeService()
	list, _, err := nodeService.ChainList()
	fmt.Println(list)
	if err != nil {
		return NewRespErr(cmd, err.Error())
	}
	return NewSuccessResp(cmd, list)
}
