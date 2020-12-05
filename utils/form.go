package utils

import "github.com/gin-gonic/gin"

func MustParams(c *gin.Context,obj interface{}) {
	err := c.BindJSON(obj)
	if err!=nil{
		panic(err)
	}
}
