package httpsvr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"pervasive-chain/config"
	"pervasive-chain/log"
	"pervasive-chain/utils"
	"time"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

//ListenAndServe 启动管理端webserver
func ListenAndServe(cfg *config.WebConfig) error {
	gin.SetMode(gin.ReleaseMode)
	httpRouter := gin.New()
	httpRouter.Use(cors())
	httpRouter.Use(log.MyGinLogger(cfg))
	httpRouter.Use(gin.Recovery())
	UseApi(httpRouter)
	exists, err := utils.FileExists(cfg.WebRoot)
	if err != nil {
		return err
	}
	if !exists {
		err := os.MkdirAll(cfg.WebRoot, os.ModePerm)
		if err != nil {
			return err
		}
	}
	exists, err = utils.FileExists(cfg.HtmlTemplate)
	if err != nil {
		return err
	}
	if !exists {
		err := os.MkdirAll(cfg.HtmlTemplate, os.ModePerm)
		if err != nil {
			return err
		}
	}
	httpRouter.LoadHTMLGlob(fmt.Sprintf("%s/*", cfg.HtmlTemplate))
	httpRouter.StaticFS("/static", http.Dir(cfg.WebRoot))
	httpSever := &http.Server{
		Addr:           cfg.HTTPListen,
		Handler:        httpRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("start %s ....\n", time.Now())
	err = httpSever.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
