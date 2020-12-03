package new

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestShardInfo(t *testing.T) {

	groupInfos := []GroupInfo{GroupInfo{}}

	groupInfosv1 := []GroupInfo{GroupInfo{Groups:groupInfos}}

	block := NewBlock{
		LockBlocks:   []LockBlock{LockBlock{}},
		CrossTrans:   groupInfosv1,
		ConfirmTrans: groupInfos,
	}
	bytes, _ := json.MarshalIndent(block, " ", " ")
	fmt.Println(string(bytes))
}

func TestTrans(t *testing.T){
	trans := Trans{}
	bytes, _ := json.MarshalIndent(trans, " ", " ")
	fmt.Println(string(bytes))
}
