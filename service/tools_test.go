package service

import (
	"fmt"
	"testing"
	"time"
)

func TestS(t *testing.T){
	now := time.Now()
	toTime := millisecondToTime(now.UnixNano()/1e6)
	fmt.Println(now,toTime)
	fmt.Println(now.UnixNano(),toTime.UnixNano())
}
