package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pervasive-chain/extern/ws"
	"testing"
)

func TestResponse(t *testing.T) {
	context := gin.Context{}
//	wsContext := ws.WsContext{}
	test(context)
}

func test(v interface{}) {
	switch v {
	case gin.Context{}:
		fmt.Println("gin context")
	case ws.WsContext{}:
		fmt.Println("ws context")
	default:
		fmt.Println("default")
	}
}
