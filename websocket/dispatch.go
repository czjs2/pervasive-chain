package websocket

import (
	"fmt"
	"pervasive-chain/model"
)

type Dispatch struct {
}

func NewDisPatch() *Dispatch {
	return &Dispatch{}
}

// todo
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
	panic(d)

}

// 生成命令
func (d *Dispatch) GenCmd(cmd model.Cmd) ([]byte, error) {
	panic(d)

}
func (d *Dispatch) WsBlockInfo(cmd model.Cmd) ([]byte, error) {
	panic(d)
}

func (d *Dispatch) WsChainInfo(cmd model.Cmd) ([]byte, error) {
	panic(d)
}

func (d *Dispatch) DoChainInfo(cmd model.Cmd) ([]byte, error) {
	panic(d)
}

func (d *Dispatch) ChainInfoById(cmd model.Cmd) ([]byte, error) {
	panic(d)
}
