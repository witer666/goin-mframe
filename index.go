package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/library"
)

func main() {
	r := gin.Default()

	//加载模板
	r.LoadHTMLGlob("template/*")

	//初始化路由
	library.InitAppRouter(r)

	//启动服务
	app := library.App{}
	app.Dispatch(r)
}
