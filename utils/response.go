package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pervasive-chain/log"
	"pervasive-chain/statecode"
	"pervasive-chain/ws"
	"reflect"
)

func FailResponse(c *gin.Context,msg interface{}) {
	log.Error(c.Request.RequestURI,msg)
	c.JSON(http.StatusOK, gin.H{"code": statecode.Fail, "message": statecode.CodeInfo(statecode.Fail)})
}

func FailResponseWithMsg(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": statecode.Fail, "message": msg})
}
func SuccessResponse(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(http.StatusOK, gin.H{"code": statecode.Success, "message": statecode.CodeInfo(statecode.Success), "data": data})
}

func ResponseWithCode(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{"code": code, "message": statecode.CodeInfo(code)})
}

// todo
func WsFailResponse(c *ws.WsContext) {
	c.JSON(statecode.Fail, "操作失败")
}

func WsSuccessResponse(c *ws.WsContext, data interface{}) {
	c.JSON(statecode.Success, data)
}

func WsFailResponseWithMsg(c *ws.WsContext,data interface{}){
	c.JSON(statecode.Fail,data)
}

func WsResponseWithCode(c *ws.WsContext, code int) {
	c.JSON(code, nil)
}

func Response(v *interface{}) {
	reflect.TypeOf(*v)
}
