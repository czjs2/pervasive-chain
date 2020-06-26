package ws

import (
	"fmt"
	"pervasive-chain/model"
	"pervasive-chain/utils"
	"testing"
	"time"
)

func TestDispatch_WsBlockInfo(t *testing.T) {
	cmd := model.Cmd{
		Uri:   "blockInfo",
		Body:  model.ReqCmd{Type: "b", Number: "0"},
		MsgId: time.Now().String(),
	}
	res, e := NewDisPatch().WsBlockInfo(cmd)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(utils.JsonBeautFormat(res), e)
}

func TestDispatch_WsChainInfo(t *testing.T) {
	cmd := model.Cmd{
		Uri:   "chainInfo",
		MsgId: time.Now().String(),
	}
	res, e := NewDisPatch().WsChainInfo(cmd)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(utils.JsonBeautFormat(res), e)
}

// ws chainInfo
func TestDispatch_DoChainInfo(t *testing.T) {
	cmd := model.Cmd{
		Uri:   "chainInfo",
		MsgId: time.Now().String(),
	}
	res, e := NewDisPatch().DoChainInfo(cmd)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(utils.JsonBeautFormat(res), e)

}

func TestDispatch_DoBlockInfo(t *testing.T) {
	cmd := model.Cmd{
		Uri:   "chainInfo",
		MsgId: time.Now().String(),
	}
	res, e := NewDisPatch().DoBlockInfo(cmd)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(utils.JsonBeautFormat(res), e)

}

func TestDispatch_chainInfoById(t *testing.T) {
	cmd := model.Cmd{
		Uri:   "cmd",
		MsgId: time.Now().String(),
		Body: model.ReqCmd{

			Type: "s",
			Cmd: model.PyCmd{
				Key:    "transfer",
				Params: model.Params{Amount:100},
			},
			Number: "0",
		},
	}
	res, e := NewDisPatch().GenCmd(cmd)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(utils.JsonBeautFormat(res), e)
}

func TestDispatch_GenCmd(t *testing.T) {
	cmd := model.Cmd{
		Uri: "cmd",
		Body: model.ReqCmd{

			Type:   "ChainInfoById",
			Cmd:    model.PyCmd{},
			Number: "10000",
		},
		MsgId: time.Now().String(),
	}
	res, e := NewDisPatch().GenCmd(cmd)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(utils.JsonBeautFormat(res), e)
}

func TestChain111(t *testing.T) {
	msg := make(chan struct{}, 1)
	go func() {
		select {
		case _, ok := <-msg:
			if !ok {
				fmt.Println("exit")

			}
		}
	}()
	time.Sleep(1 * time.Second)
	close(msg)
	time.Sleep(10 * time.Second)

}
