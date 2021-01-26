package log

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

//MyGinLogger gin自定义logrus日志
func MyGinLogger(logPath string) gin.HandlerFunc {
	logFilePath := logPath
	logFileName := "web.log"
	fileName := path.Join(logFilePath, logFileName)
	fileSrc, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {

	}
	writers := []io.Writer{
		fileSrc,
		os.Stdout,
	}

	logger := logrus.New()
	logger.Out = io.MultiWriter(writers...)
	logger.SetReportCaller(true)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		//rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
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
	logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	}
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			c.Abort()
			return
		}
		Debug(c.Request.RequestURI, c.Request.Method, buf.String())
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(buf.String())))
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

func Info(msg ...interface{}) {
	Logger.Infoln(msg)
}

func Debug(msg ...interface{}) {
	Logger.Debugln(msg)
}

func Warn(msg ...interface{}) {
	Logger.Warnln(msg)
}
func Error(msg ...interface{}) {
	Logger.Errorln(msg)
}

// Logger 逻辑日志
var Logger *logrus.Logger

//MyLogicLogger 自定义logrus日志
func MyLogicLogger(logPath string) (*logrus.Logger, error) {
	logFilePath := logPath
	logFileName := "logic.log"
	fileName := path.Join(logFilePath, logFileName)
	fileSrc, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	writers := []io.Writer{
		fileSrc,
		os.Stdout,
	}

	logger := logrus.New()
	logger.Out = io.MultiWriter(writers...)
	logger.SetLevel(logrus.InfoLevel)
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		//rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		//	logrus.DebugLevel: logWriter,
		//logrus.WarnLevel:  logWriter,
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
