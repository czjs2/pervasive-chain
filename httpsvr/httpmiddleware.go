package httpsvr

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"pervasive-chain/statecode"
	"pervasive-chain/utils"
)

var DevIdError = errors.New("dev error")

// 通用参数验证
func ParamVerifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.RequestURI
		ok := validateManager.Exists(path)
		if ok {
			buf := new(bytes.Buffer)
			_, err := buf.ReadFrom(c.Request.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, nil)
				c.Abort()
				return
			}
			fmt.Println(buf.String())
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(buf.String())))
			err = validateManager.Execute(path, buf.String())
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code":statecode.Fail,"message":"参数错误"})
				c.Abort()
			}
		}
		c.Next()

	}
}

// 通信验证
func TokenVerifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		uid, err := CheckToken(token)
		if err == DevIdError {
			c.JSON(http.StatusOK, gin.H{"code": -2, "msg": "认证失败"})
			c.Abort()
		} else if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "认证失败"})
			c.Abort()
		} else {
			c.Request.ParseForm()
			c.Request.PostForm.Add("uid", uid)
			c.Request.Form.Add("uid", uid)
			c.Next()
		}
	}
}

func CheckToken(token string) (string, error) {
	uid, err := utils.ValidToken(token)
	if err != nil {
		return uid, err
	}
	return uid, nil
}
