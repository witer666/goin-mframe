package api

import (
	"strconv"

	"github.com/gin-gonic/actions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/models/dao"
	"github.com/gin-gonic/models/service/page"
	"github.com/gin-gonic/utils"
)

type ActionUserUpdate struct {
	actions.ActionBase
}

func NewActionUserUpdate(c *gin.Context) *ActionUserUpdate {
	action := new(ActionUserUpdate)
	action.Params = actions.HandlerRequestParams(c)

	return action
}

func (action *ActionUserUpdate) Invoke(c *gin.Context) {
	pUser := page.Service_Page_Users{}
	id, err := strconv.ParseInt(action.Params["id"][0], 10, 64)
	if err != nil {
		id = 0
	}
	user := &dao.Dao_Users{
		Id: id,
	}
	values := &dao.Dao_Users{
		UserName: action.Params["uname"][0],
		Password: action.Params["pwd"][0],
	}
	res := pUser.UpdateUser(user, values)
	bolRet := true
	if res < 0 {
		bolRet = false
	}
	utils.ReturnJson(c, bolRet, 0, "")
}
