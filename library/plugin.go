package library

import (
	"net/http"

	"github.com/gin-gonic/actions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/library/log"
	"github.com/gin-gonic/utils"
)

/**
 *
 * 请求数据安全效验中间件函数
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	void
 * @return 	gin.HandlerFunc
 *
 */
func HandlerUseRequestParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqParams := actions.HandlerRequestParams(c)
		for _, value := range reqParams {
			if len(value) == 1 && utils.GetValueType(value) == "slice" {
				value[0] = utils.Quote(value[0])
			}
		}
	}
}

/**
 *
 * 自定义错误恢复中间件
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	c gin.Context
 * @param	recovered any
 * @return 	void
 *
 */
func HandlerUseCustomRecovery(c *gin.Context, recovered any) {
	fields := map[string]interface{}{
		"err": recovered,
	}
	log.Init(c).Error(fields)
	utils.ReturnJson(c, "", http.StatusInternalServerError, "internal code error")
	c.Abort()
}
