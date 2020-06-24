package service

import (
	"fmt"
	"testing"
)

func TestNodeService_OnLineList(t *testing.T) {
	list, i, e := NewNodeService().OnLineList()
	fmt.Println(list,i,e)
}
