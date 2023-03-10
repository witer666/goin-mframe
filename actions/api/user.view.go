package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/models/service/page"
	"github.com/gin-gonic/utils"
)

type ActionUserView struct {
}

func (action *ActionUserView) Invoke(c *gin.Context) {
	pUser := page.Service_Page_Users{}
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		id = 0
	}
	userModel := pUser.GetUserById(id)
	utils.ReturnJson(c, userModel, 0, "")
}
