package main

import (
	"fmt"
	"pervasive-chain/form"
	"pervasive-chain/utils"
)

type MockData struct {
}

var chainTypeList = []string{"b", "r", "s"}



func (m *MockData) GenRandHeart() form.HeartBeatFrom {
	f := form.HeartBeatFrom{
		Type:   chainTypeList[utils.Rand(3)],
		Number: fmt.Sprintf("%d", utils.Rand(100000)),
		Id:     fmt.Sprintf("%d", utils.Rand(100)),
		Time:   utils.GetNowTime(),
	}
	return f
}

func (m *MockData) GenBlock() form.ReportBlockForm {
	blockForm := form.ReportBlockForm{
		Type: chainTypeList[utils.Rand(3)],
		Number:fmt.Sprintf("%d",utils.Rand(100000)),
		Id:fmt.Sprintf("%d",utils.Rand(100)),
	}
	return blockForm
}
