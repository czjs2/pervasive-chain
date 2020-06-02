package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pervasive-chain/code"
	"pervasive-chain/form"
	"pervasive-chain/log"
	"pervasive-chain/service"
	"pervasive-chain/utils"
)

// 心跳
func ReportHeadBeatHandler(c *gin.Context) {
	var heartForm form.HeartBeatFrom
	err := c.ShouldBind(&heartForm)
	if err != nil {
		log.Logger.Errorln(c.Request.URL, "heart report parameter is error ", err.Error())
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	nodeService := service.NewNodeService()
	node, err := nodeService.FindAndUpdate(heartForm)
	if err != nil {
		if err.Error() != code.NoDocumentError {
			log.Logger.Errorln(c.Request.URL, "heart report insert error", err.Error())
			c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
			return
		}
	}
	// 每次有心跳上报时更新链信息
	statisticService := service.NewStatisticService()
	_, err = statisticService.CountChain()
	if err != nil {
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	if node != nil && node.Cmd != nil {
		c.JSONP(http.StatusOK, utils.SuccessResponse(node.Cmd))
	} else {
		c.JSONP(http.StatusOK, utils.ResponseWithCode(code.NoCmd, "没有命令下发", nil))
	}
}

// 块信息
func ReportBlockHandler(c *gin.Context) {
	var blockForm form.ReportBlockForm
	err := c.ShouldBind(&blockForm)
	if err != nil {
		log.Logger.Errorln("block report parameter is error", err.Error())
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	blockService := service.NewBlockService()
	_, err = blockService.UpdateBlockInfo(blockForm)
	if err != nil {
		log.Logger.Errorln(c.Request.URL, "block report insert error ", err.Error())
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	// todo 需要事务处理两张表 ?
	historyBlockService := service.NewHistoryBlockService()
	_, err = historyBlockService.UpdateBlockInfo(blockForm)
	if err != nil {
		log.Logger.Errorln(c.Request.URL, "block report history block insert error ", err.Error())
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	// 更新tps信息
	statisticService := service.NewStatisticService()
	_, err = statisticService.CountChain()
	if err != nil {
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}

	c.JSONP(http.StatusOK, utils.SuccessResponse(nil))
}

// 流量信息
func ReportFlowHandler(c *gin.Context) {
	var flowForm form.ReportFlowForm
	err := c.ShouldBind(&flowForm)
	if err != nil {
		log.Logger.Errorln(c.Request.URL, "flow report parameter is error ", err.Error())
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	flowService := service.NewFlowService()
	_, err = flowService.UpdateFlowInfo(flowForm)
	if err != nil {
		log.Logger.Errorln(c.Request.URL, "flow report insert is error ", err.Error())
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	// 更新整体流量信息
	statisticService := service.NewStatisticService()
	_, err = statisticService.CountFlow()
	if err != nil {
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}

	c.JSONP(http.StatusOK, utils.SuccessResponse(nil))
}
