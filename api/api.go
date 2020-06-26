package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pervasive-chain/code"
	"pervasive-chain/config"
	"pervasive-chain/form"
	"pervasive-chain/log"
	"pervasive-chain/model"
	"pervasive-chain/service"
	"pervasive-chain/utils"
	"pervasive-chain/ws"
	"time"
)

func TestPing(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{"info:": "pong"})
}

// 清空下发命令
func ClearCmdHandler(c *gin.Context) {
	nodeService := service.NewNodeService()
	_, err := nodeService.ClearCmd()
	if err != nil {
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
	} else {
		c.JSONP(http.StatusOK, utils.SuccessResponse("success"))
	}
}

// 心跳
func ReportHeadBeatHandler(c *gin.Context) {
	var heartForm form.HeartBeatFrom
	err := c.ShouldBind(&heartForm)
	log.Logger.Debug(heartForm)
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
	if cmdCanSend(node) {
		c.JSONP(http.StatusOK, utils.SuccessResponse(node.Cmd))
	} else {
		c.JSONP(http.StatusOK, utils.SuccessResponse("没有命令下发"))
	}
}

// 判断命令是否能下发
func cmdCanSend(node *model.Node) bool {
	if node == nil {
		return false
	}
	if node.Type != config.SChain {
		return false
	}
	// 心跳时间和命令时间 效验
	if time.Now().Sub(node.LastTime).Seconds() < config.HeartBeatTime && time.Now().Sub(node.CmdTime).Seconds() < config.HeartBeatTime {
		return true
	}
	return false
}

// 块信息
func ReportBlockHandler(c *gin.Context) {
	var blockForm form.ReportBlockForm
	err := c.ShouldBind(&blockForm)
	log.Logger.Debug(blockForm)
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
	historyBlockService := service.NewHistoryBlockService()
	_, err = historyBlockService.UpdateBlockInfo(blockForm)
	if err != nil {
		log.Logger.Errorln(c.Request.URL, "block report history block insert error ", err.Error())
		c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
		return
	}

	// 广播消息 todo 写到中间件里更好 ？
	ws.BroadcastBlock(blockForm)

	c.JSONP(http.StatusOK, utils.SuccessResponse(nil))
}

// 流量信息
func ReportFlowHandler(c *gin.Context) {
	var flowForm form.ReportFlowForm
	err := c.ShouldBind(&flowForm)
	log.Logger.Debug(flowForm)
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
	_, err = statisticService.CountFlow(flowForm)
	if err != nil {
		if err.Error() != code.NoDocumentError {
			c.JSONP(http.StatusOK, utils.FailResponse(err.Error()))
			return

		}
	}
	c.JSONP(http.StatusOK, utils.SuccessResponse(nil))
}
