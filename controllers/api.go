package controllers

import (
	"github.com/gin-gonic/actions/api"
	"github.com/gin-gonic/gin"
)

/**
 *
 * RestfulApi接口配置类
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 *
 */
type ControllerApi struct {
	Actions map[string]func(c *gin.Context)
}

func (ctlApi *ControllerApi) Init() {
	ctlApi.Actions = map[string]func(c *gin.Context){
		"get:/v1/api/user/list": func(c *gin.Context) {
			_actionList := api.NewActionUserList(c)
			_actionList.Invoke(c)
		},
		"post:/v1/api/user/create": func(c *gin.Context) {
			_actionCreate := api.NewActionUserCreate(c)
			_actionCreate.Invoke(c)
		},
		"post:/v1/api/user/update": func(c *gin.Context) {
			_actionUpdate := api.NewActionUserUpdate(c)
			_actionUpdate.Invoke(c)
		},
		"get:/v1/api/user/view": func(c *gin.Context) {
			_actionView := api.NewActionUserView(c)
			_actionView.Invoke(c)
		},
	}
}
