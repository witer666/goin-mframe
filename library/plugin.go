package library

import (
	"github.com/gin-gonic/actions"
	"github.com/gin-gonic/gin"
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
