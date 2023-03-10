package actions

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/utils"
)

type ActionBase struct {
	Params map[string][]string
}

/**
 *
 * 多种请求方式请求参数处理
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	c gin.Context
 * @return 	map
 *
 */
func HandlerRequestParams(c *gin.Context) map[string][]string {
	var params map[string][]string
	method := c.Request.Method
	switch method {
	case "GET":
		params = c.Request.URL.Query()
	case "POST":
		//POST Raw请求
		var jsonParam map[string]interface{}
		c.Request.ParseMultipartForm(256)
		body, _ := c.GetRawData()
		if len(body) > 0 {
			params = make(map[string][]string)
			_ = json.Unmarshal(body, &jsonParam)
			for k, v := range jsonParam {
				params[k] = append(params[k], utils.TypeInterface2String(v))
			}
		} else {
			//POST UrlEncode请求
			params = c.Request.Form
		}
	default:
	}

	return params
}
