package cmd

import (
	"encoding/json"
	"fmt"
	"testing"
)


func Test01(t *testing.T){
	src:=`{"uri":"chainInfo","body":{"type":"b","number":"1111","hash":"sdfsadsf"},"msgId":"111111"}`
	cmd := BaseCmd{}
	err := json.Unmarshal([]byte(src), &cmd)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(cmd.Uri,cmd)
}


func TestBaseCmd(t *testing.T){
	//src:=`{"uri":"","body":{"type":"","number":"","hash":""},"msgId":""}`
	cmd := ChainInfoCmd{}
	bytes, e := json.Marshal(cmd)
	fmt.Println(string(bytes),e)
}
