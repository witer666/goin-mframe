package library

import (
	"strings"

	"github.com/gin-gonic/controllers"
	"github.com/gin-gonic/gin"
)

/**
 *
 * 初始化gin请求路由
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	r gin.Engine
 * @return  void
 *
 */
func InitAppRouter(r *gin.Engine) {
	_app := App{}
	var _urls []string
	var _url string
	var _method string

	//Restful Api控制器
	_apiController := controllers.ControllerApi{}
	_apiController.Init()
	for k, vfunc := range _apiController.Actions {
		_urls = strings.Split(k, ":")
		if len(_urls) == 1 {
			_url = _urls[0]
			_method = "GET"
		} else {
			_url = _urls[1]
			_method = _urls[0]
		}
		_app.RegisterRouter(r, _method, _url, vfunc)
	}

	//Admin管理页面控制器
	_adminController := controllers.ControllerAdmin{}
	_adminController.Init()
	for k, vfunc := range _adminController.Actions {
		_urls = strings.Split(k, ":")
		if len(_urls) == 1 {
			_url = _urls[0]
			_method = "GET"
		} else {
			_url = _urls[1]
			_method = _urls[0]
		}
		_app.RegisterRouter(r, _method, _url, vfunc)
	}
}
