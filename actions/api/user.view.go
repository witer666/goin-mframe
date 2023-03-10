package api

import (
	"strconv"

	"github.com/gin-gonic/actions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/models/service/page"
	"github.com/gin-gonic/utils"
)

type ActionUserView struct {
	actions.ActionBase
}

func NewActionUserView(c *gin.Context) *ActionUserView {
	action := new(ActionUserView)
	action.Params = actions.HandlerRequestParams(c)

	return action
}

func (action *ActionUserView) Invoke(c *gin.Context) {
	pUser := page.Service_Page_Users{}
	id, err := strconv.ParseInt(action.Params["id"][0], 10, 64)
	if err != nil {
		id = 0
	}
	userModel := pUser.GetUserById(id)
	utils.ReturnJson(c, userModel, 0, "")
}
