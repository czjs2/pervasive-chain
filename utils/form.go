package utils

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/config"
)

func MustParams(c *gin.Context, obj interface{}) {
	err := c.BindJSON(obj)
	if err != nil {
		panic(err)
	}
}

func IsValidNodeId(nodeId string) bool {
	return len(nodeId) == 52
}

func  IsValidChainKey(chainKey, chainType string) bool {
	if chainType == config.BeaconType && len(chainKey) == 1 {
		return true
	}
	if chainType == config.RelayType && len(chainKey) == 2 {
		return true
	}
	if chainType == config.SharedType && len(chainKey) == 4 {
		return true
	}
	return false
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
