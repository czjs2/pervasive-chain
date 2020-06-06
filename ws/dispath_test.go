package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"pervasive-chain/model"
	"testing"
	"time"
)

func TestDispatch_DoChainInfo(t *testing.T) {
	cmd := model.Cmd{
		Uri:   "blockInfo",
		MsgId: time.Now().String(),
	}
	res, e := NewDisPatch().DoChainInfo(cmd)
	if e!=nil{
		fmt.Println(e)
		return
	}
	var buf bytes.Buffer
	_ = json.Indent(&buf, res, "", "     ")
	fmt.Println(buf.String())

}

func TestDispatch_DoBlockInfo(t *testing.T) {
	cmd := model.Cmd{
		Uri:   "chainInfo",
		MsgId: time.Now().String(),
	}
	res, e := NewDisPatch().DoBlockInfo(cmd)
	if e!=nil{
		fmt.Println(e)
		return
	}
	var buf bytes.Buffer
	_ = json.Indent(&buf, res, "", "     ")
	fmt.Println(buf.String())

}
