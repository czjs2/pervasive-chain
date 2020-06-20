package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTools(t *testing.T){
	cmd := Cmd{}
	bytes, err := json.Marshal(cmd)
	fmt.Println(string(bytes),err)
	res:=`{"uri":"cmd","body":{"type":"b","cmd":{"key":"transfer","params":[100]}},"msgId":"msgId%d"}`
	err = json.Unmarshal([]byte(res), &cmd)
	fmt.Println(cmd)
}
