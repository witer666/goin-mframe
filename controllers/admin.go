package controllers

import (
	"github.com/gin-gonic/actions/admin"
	"github.com/gin-gonic/gin"
)

/**
 *
 * Admin管理页面接口配置类
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 *
 */
type ControllerAdmin struct {
	Actions map[string]func(c *gin.Context)
}

func (ctlAdmin *ControllerAdmin) Init() {
	ctlAdmin.Actions = map[string]func(c *gin.Context){
		"get:/admin/index": func(c *gin.Context) {
			_actionAdminIndex := admin.ActionAdminIndex{}
			_actionAdminIndex.Invoke(c)
		},
	}
}
