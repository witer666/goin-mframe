package api

import (
	"github.com/gin-gonic/actions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/models/dao"
	"github.com/gin-gonic/models/service/page"
	"github.com/gin-gonic/utils"
)

type ActionUserCreate struct {
	actions.ActionBase
}

func NewActionUserCreate(c *gin.Context) *ActionUserCreate {
	action := new(ActionUserCreate)
	action.Params = actions.HandlerRequestParams(c)

	return action
}

func (action *ActionUserCreate) Invoke(c *gin.Context) {
	pUser := page.Service_Page_Users{}
	user := &dao.Dao_Users{
		UserName: action.Params["uname"][0],
		Password: action.Params["pwd"][0],
	}
	res := pUser.InsertUser(user)
	bolRet := true
	if res < 1 {
		bolRet = false
	}
	utils.ReturnJson(c, bolRet, 0, "")
}
