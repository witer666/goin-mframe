package library

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type App struct {
}

/**
 *
 * 注册路由GET/POST请求
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	r gin.Engine
 * @param 	method string 请求方式
 * @param 	uri	string 请求uri
 * @param 	action 请求处理函数
 * @return  void
 *
 */
func (app *App) RegisterRouter(r *gin.Engine, method string, uri string, action func(c *gin.Context)) {
	switch strings.ToUpper(method) {
	case "GET":
		r.GET(uri, action)
	case "POST":
		r.POST(uri, action)
	}
}

/**
 *
 * 运行APP
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	r gin.Engine
 * @return  void
 *
 */
func (app *App) Dispatch(r *gin.Engine) {
	r.Run()
}
