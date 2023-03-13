package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/library"
)

func main() {
	r := gin.New()

	//加载模板
	r.LoadHTMLGlob("template/*")

	//初始化路由
	library.InitAppRouter(r)

	//gin日志
	logFile, _ := os.OpenFile("log/errors.log", os.O_RDWR|os.O_CREATE, 0755)
	gin.DefaultWriter = io.MultiWriter(logFile)

	//启动服务
	app := library.App{}
	app.Dispatch(r)
}
