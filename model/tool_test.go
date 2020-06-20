package model

import (
	"fmt"
	"testing"
)

func TestTools(t *testing.T){
	block := Block{}
	obj := ObjToMap(block)
	fmt.Println(obj)
}
