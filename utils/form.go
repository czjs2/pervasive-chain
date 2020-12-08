package utils

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/config"
)

func MustParams(c *gin.Context,obj interface{}) {
	err := c.BindJSON(obj)
	if err!=nil{
		panic(err)
	}
}



func IsValidChainKey(v string) bool {
	return len(v) == 3 || len(v) == 5

}

func IsValidChain(v string) bool {
	switch v {
	case config.BeaconType:
		return true
	case config.RelayType:
		return true
	case config.SharedType:
		return true
	default:
		return false
	}
}