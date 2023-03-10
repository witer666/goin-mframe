package api

import (
	"github.com/gin-gonic/actions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/models/dao"
	"github.com/gin-gonic/models/service/page"
	"github.com/gin-gonic/utils"
)

type ActionUserList struct {
	actions.ActionBase
}

func NewActionUserList(c *gin.Context) *ActionUserList {
	action := new(ActionUserList)
	action.Params = actions.HandlerRequestParams(c)

	return action
}

func (action *ActionUserList) Invoke(c *gin.Context) {
	pUser := page.Service_Page_Users{}
	userName := ""
	if _, ok := action.Params["uname"]; ok {
		userName = action.Params["uname"][0]
	}
	var userModel *[]dao.Dao_Users
	if userName == "" {
		userModel = pUser.GetUserList(0, 10)
	} else {
		userModel = pUser.GetUserList(0, 10, "user_name = ?", userName)
	}
	utils.ReturnJson(c, userModel, 0, "")
}
