package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"pervasive-chain/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//MyGinLogger gin自定义logrus日志
func MyGinLogger(config *config.WebConfig) gin.HandlerFunc {
	logFilePath := config.LogPath
	logFileName := "web.log"
	fileName := path.Join(logFilePath, logFileName)
	var writers []io.Writer
	if config.Debug {
		writers = append(writers, os.Stdout)
	} else {
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		writers = append(writers, src)
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logger := logrus.New()
	logger.Out = fileAndStdoutWriter
	logger.SetLevel(logrus.DebugLevel)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		//rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(lfHook)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqURI,
		}).Info()
	}
}

var Logger *logrus.Logger

func MyLogicLogger(config *config.WebConfig) (*logrus.Logger, error) {
	logFilePath := config.LogPath
	logFileName := "logic.log"
	fileName := path.Join(logFilePath, logFileName)
	var writers []io.Writer
	if config.Debug {
		writers = append(writers, os.Stdout)
	} else {
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, err
		}
		writers = append(writers, src)
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logger := logrus.New()
	logger.Out = fileAndStdoutWriter
	logger.SetLevel(logrus.DebugLevel)
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		//rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		return nil, err
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(lfHook)
	Logger = logger
	return logger, nil
}
