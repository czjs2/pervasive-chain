package service

import (
	"fmt"
	"testing"
	"time"
)

func TestS(t *testing.T){
	now := time.Now()
	i := now.UnixNano() / 1e6
	toTime := millisecondToTime(i)
	unix := time.Unix(i/1000, (i%1000)*1e9)
	fmt.Println(now,toTime)
	fmt.Println(unix)
}
