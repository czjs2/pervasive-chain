package service

import (
	"fmt"
	"pervasive-chain/form"
	"testing"
)

func TestStatisticService_AllChain(t *testing.T) {
	service := NewStatisticService()
	chain, e := service.AllChain()
	fmt.Println(chain,e)
}

func TestStatisticService_CountFlow(t *testing.T) {
	statisticService := NewStatisticService()
	flow, e := statisticService.CountFlow(form.ReportFlowForm{})
	fmt.Println(flow,e)
}

func TestStatisticService_CountChain(t *testing.T) {
	service := NewStatisticService()
	chain, e := service.CountChain("100003", "s")
	fmt.Println(chain,e)
}

func TestStatisticService_CountNode(t *testing.T) {
	service := NewStatisticService()
	node, e := service.CountNode()
	fmt.Println(node,e)
}

func TestStatisticService_CountTps(t *testing.T) {

	service := NewStatisticService()
	tps, e := service.CountTps()
	fmt.Println(tps,e)
}
