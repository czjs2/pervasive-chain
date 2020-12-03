package api

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/form"
)

func ReportFlowHandler(c *gin.Context){
	var reportFlowForm form.ReportFlowForm
	err := c.BindJSON(&reportFlowForm)
	if err!=nil{
		FailResponse(c)
		return
	}
}
