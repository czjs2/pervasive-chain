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
	fmt.Println(utils.JsonBeautFormat(res),e)

}

func TestDispatch_GenCmd(t *testing.T) {
	cmd := model.Cmd{
		Uri: "cmd",
		Body: model.PyCmd{
			Key:    "transfer",
			Params: []interface{}{1000},
		},
		MsgId: time.Now().String(),
	}
	res, e := NewDisPatch().GenCmd(cmd)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(utils.JsonBeautFormat(res),e)
}
