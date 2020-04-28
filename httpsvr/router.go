package httpsvr

import (
	"github.com/gin-gonic/gin"
)

// 仅限用户使用api
func UserApi(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/api/v1/")
	return group
}
