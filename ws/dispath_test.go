package ws

import (
	"fmt"
	"pervasive-chain/model"
	"pervasive-chain/utils"
	"testing"
	"time"
)

func TestDispatch_DoChainInfo(t *testing.T) {
	cmd := model.Cmd{
		Uri:   "blockInfo",
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

			Type:   "ChainInfoById",
			Cmd:    model.PyCmd{},
			Number: 100000,
		},
	}
	res, e := NewDisPatch().ChainInfoById(cmd)
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
			Number: 100000,
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
