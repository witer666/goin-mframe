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
		"post:/v1/api/user/list": func(c *gin.Context) {
			_actionList := api.NewActionUserList(c)
			_actionList.Invoke(c)
		},
		"get:/v1/api/user/create": func(c *gin.Context) {
			_actionCreate := api.ActionUserCreate{}
			_actionCreate.Invoke(c)
		},
		"get:/v1/api/user/update": func(c *gin.Context) {
			_actionUpdate := api.ActionUserUpdate{}
			_actionUpdate.Invoke(c)
		},
		"get:/v1/api/user/view": func(c *gin.Context) {
			_actionView := api.ActionUserView{}
			_actionView.Invoke(c)
		},
	}
}
