package utils

import (
	"fmt"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	//now := time.Now()
	//fmt.Println(now)
	//fmt.Println(now.UTC())
	////location, _ := time.LoadLocation("")
	// 2006-01-02T15:04:05Z07:00
	parse, err := ParseRFCTime("2012-11-01T22:08:41+00:00")
	fmt.Println(parse,err)
	fmt.Println(parse.UTC())

}

var timeZone = []string{
	"-1200",
	"-1000",
	"-0900",
	"-0800",
	"-0700",
	"-0600",
	"-0500",
	"-0400",
	"-0300",
	"-0200",
	"-0100",
	"+0000",
	"+0100",
	"+0200",
	"+0300",
	"+0400",
	"+0500",
	"+0600",
	"+0700",
	"+0800",
	"+0900",
	"+1000",
	"+1100",
	"+1200",
}

//2014-03-25T06:26:01.927+0800
func ParseUtcTime(t string) (time.Time,error){

	return time.Time{},nil
}







func TestResponse(t *testing.T) {
	localTime, e := ParseLocalTime("2020-12-07 07:05:47.8066936")
	fmt.Println(localTime, e)
}
