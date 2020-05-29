package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"pervasive-chain/form"
	"pervasive-chain/log"
	"pervasive-chain/service"
	"pervasive-chain/utils"
	"pervasive-chain/ws"
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
	// todo
	bytes, _ := json.Marshal(heartForm)

	ws.BroadCast(bytes);
	nodeService := service.NewNodeService()
	_, err = nodeService.UpdateNodeInfo(heartForm)
	if err != nil {
		log.Logger.Errorln(c.Request.URL, "heart report insert error", err.Error())
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	c.JSONP(http.StatusOK, utils.SuccessResponse(nil))
}

// 块信息
func ReportBlockHandler(c *gin.Context) {
	var blockForm form.ReportBlockForm
	err := c.ShouldBind(blockForm)
	if err != nil {
		log.Logger.Errorln("block report parameter is error", err.Error())
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	bytes, _ := json.Marshal(blockForm)

	ws.BroadCast(bytes);
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
	bytes, _ := json.Marshal(flowForm)

	ws.BroadCast(bytes);
	flowService := service.NewFlowService()
	_, err = flowService.UpdateFlowInfo(flowForm)
	if err != nil {
		log.Logger.Errorln(c.Request.URL, "flow report insert is error ", err.Error())
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}
	c.JSONP(http.StatusOK, utils.SuccessResponse(nil))
}
