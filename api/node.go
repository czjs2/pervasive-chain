package api

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/form"
	"pervasive-chain/service/impl"
)

func HeartBeatHandler(c *gin.Context) {
	var heartBeatFrom form.HeartBeatFrom
	err := c.BindJSON(&heartBeatFrom)
	if err != nil {
		FailResponse(c)
		return
	}
	nodeService := impl.NewNodeService()
	_, err = nodeService.UpdateNodeInfo(heartBeatFrom)
	if err != nil {
		FailResponse(c)
		return
	}
	SuccessResponse(c, nil)
}
