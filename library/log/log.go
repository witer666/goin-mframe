package log

import (
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/utils"
	log "github.com/sirupsen/logrus"
)

/**
 *
 * goin日志操作类
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/11 10:47:35
 *
 */
type GoinLog struct {
	Logger *log.Entry
}

/**
 *
 * 初始化logrus环境变量
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	void
 * @return  GoinLog
 *
 */
func Init(c *gin.Context) *GoinLog {
	glog := new(GoinLog)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.ErrorLevel)
	initFields := InitFields(c)
	glog.Logger = log.WithFields(initFields)

	return glog
}

/**
 *
 * 初始化日志通用字段
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	c gin.Context
 * @return  map
 *
 */
func InitFields(c *gin.Context) map[string]interface{} {
	var initFields map[string]interface{} = map[string]interface{}{}

	//获取调用者信息
	pc := make([]uintptr, 1)
	runtime.Callers(3, pc)
	if len(pc) > 0 {
		funcPc := runtime.FuncForPC(pc[0])
		initFields["func"] = funcPc.Name()
		fileName, lineNo := funcPc.FileLine(pc[0])
		initFields["Line"] = lineNo
		initFields["File"] = fileName
	}
	if c == nil {
		return initFields
	}
	initFields["Ip"] = c.ClientIP()
	initFields["UserName"] = c.Request.URL.User.Username()
	initFields["Referer"] = c.Request.Referer()
	initFields["Host"] = c.Request.Host
	initFields["Method"] = c.Request.Method
	initFields["Url"] = c.Request.URL.String()
	initFields["UserAgent"] = c.Request.UserAgent()
	initFields["Proto"] = c.Request.Proto

	return initFields
}

/**
 *
 * 日志文件分级保存
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	level string 日志级别
 * @return  void
 *
 */
func InitOutput(level string) {
	fileName := "log/" + utils.CurrentNumericDate() + "." + level + ".log"
	logFile, errOpen := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if errOpen == nil {
		log.SetOutput(logFile)
	} else {
		log.SetOutput(os.Stdout)
	}
}

func (glog *GoinLog) Info(fields map[string]interface{}, args ...interface{}) {
	InitOutput("info")
	glog.Logger.WithFields(fields).Info(args...)
}

func (glog *GoinLog) Warn(fields map[string]interface{}, args ...interface{}) {
	InitOutput("warn")
	glog.Logger.WithFields(fields).Warn(args...)
}

func (glog *GoinLog) Warning(fields map[string]interface{}, args ...interface{}) {
	InitOutput("warning")
	glog.Logger.WithFields(fields).Warning(args...)
}

func (glog *GoinLog) Debug(fields map[string]interface{}, args ...interface{}) {
	InitOutput("debug")
	glog.Logger.WithFields(fields).Debug(args...)
}

func (glog *GoinLog) Error(fields map[string]interface{}, args ...interface{}) {
	InitOutput("error")
	glog.Logger.WithFields(fields).Error(args...)
}

func (glog *GoinLog) Fatal(fields map[string]interface{}, args ...interface{}) {
	InitOutput("fatal")
	glog.Logger.WithFields(fields).Fatal(args...)
}

func (glog *GoinLog) Panic(fields map[string]interface{}, args ...interface{}) {
	InitOutput("panic")
	glog.Logger.WithFields(fields).Panic(args...)
}

func (glog *GoinLog) Trace(fields map[string]interface{}, args ...interface{}) {
	InitOutput("trace")
	glog.Logger.WithFields(fields).Trace(args...)
}
