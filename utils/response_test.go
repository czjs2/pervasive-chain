package utils

import (
	"fmt"
	"pervasive-chain/config"
	"testing"
	"time"
)


func Test01(t *testing.T){
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.UTC())
	//location, _ := time.LoadLocation("")
	parse, e := time.Parse(config.SysTimefrom, "2020-12-07 07:05:47.8066936")
	fmt.Println(parse,e)

}



func TestResponse(t *testing.T) {


}

