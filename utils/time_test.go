package utils

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T){
	//src := "2012-11-01T22:08:41+00:00"
	src := "2020-01-01T00:00:00.927+08:00"
	time, err := ParseRFCTime(src)
	fmt.Println(time,err)
}
