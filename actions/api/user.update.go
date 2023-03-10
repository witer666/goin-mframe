package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/models/dao"
	"github.com/gin-gonic/models/service/page"
	"github.com/gin-gonic/utils"
)

type ActionUserUpdate struct {
}

func (action *ActionUserUpdate) Invoke(c *gin.Context) {
	pUser := page.Service_Page_Users{}
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		id = 0
	}
	user := &dao.Dao_Users{
		Id: id,
	}
	values := &dao.Dao_Users{
		UserName: c.Query("uname"),
		Password: "admin",
	}
	res := pUser.UpdateUser(user, values)
	bolRet := true
	if res < 0 {
		bolRet = false
	}
	utils.ReturnJson(c, bolRet, 0, "")
}
