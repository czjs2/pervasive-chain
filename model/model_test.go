package model

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_node(t *testing.T) {
	cmdInfo := fmt.Sprintf(`{"uri":"cmd","body":{"key":{"trans":1000}},"msgId":"msgId%d"}`, time.Now().Nanosecond())
	cmd:= Cmd{}
	err := json.Unmarshal([]byte(cmdInfo), &cmd)
	fmt.Println(cmd,err)
}
