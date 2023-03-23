package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActionAdminForm struct {
}

func (action *ActionAdminForm) Initialize() bool {
	return true
}

func (action *ActionAdminForm) Invoke(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.form.page.tpl", nil)
}
