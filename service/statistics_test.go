package service

import (
	"fmt"
	"pervasive-chain/form"
	"pervasive-chain/utils"
	"testing"
	"time"
)

func TestStatisticService_AllChain(t *testing.T) {
	service := NewStatisticService()
	chain, e := service.AllChain()
	fmt.Println(utils.JsonFormat(chain), e)
}

func TestStatisticService_CountFlow(t *testing.T) {
	form := form.ReportFlowForm{
		Type:   "s",
		Number: "10000",
		Id:     "1000",
		Time:   100000,
		In:     100,
		Out:    100,
	}
	statisticService := NewStatisticService()
	flow, e := statisticService.CountFlow(form)
	fmt.Println(flow, e)
}

func TestStatisticService_CountChain(t *testing.T) {
	service := NewStatisticService()
	chain, e := service.CountChain("100003", "s")
	fmt.Println(chain, e)
}

func TestStatisticService_CountNode(t *testing.T) {
	service := NewStatisticService()
	node, e := service.CountNode()
	fmt.Println(node, e)
}

func TestStatisticService_CountTps(t *testing.T) {

	service := NewStatisticService()
	tps, e := service.CountTps()
	fmt.Println(tps, e)
}

func TestNanTime(t*testing.T){
	now := time.Now()
	toTime := millisecondToTime(now.UnixNano())
	fmt.Println(toTime)
}
