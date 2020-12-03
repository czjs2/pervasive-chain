package api

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/form"
	"pervasive-chain/service/impl"
)

func ReportBlockInfoHandler(c *gin.Context) {
	var reportBlockFrom form.ReportBlockForm
	err := c.BindJSON(&reportBlockFrom)
	if err != nil {
		FailResponse(c)
		return
	}
	blockService := impl.NewBlockService()
	_, code, err := blockService.UpdateBlock(reportBlockFrom)
	if err != nil {
		ResponseWithCode(c, code)
		return
	}
	SuccessResponse(c, nil)
}


