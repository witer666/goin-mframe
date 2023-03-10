package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActionAdminIndex struct {
}

func (action *ActionAdminIndex) Initialize() bool {
	return true
}

func (action *ActionAdminIndex) Invoke(c *gin.Context) {
	only := map[string]string{
		"uname": "go",
		"sex":   "男",
	}
	data := []map[string]string{}
	data = append(data, only)
	data = append(data, only)
	data = append(data, only)
	data = append(data, only)
	res := map[string]interface{}{
		"title": "Hello HTML",
		"list":  data,
	}
	c.HTML(http.StatusOK, "admin.index.page.tpl", res)
}
