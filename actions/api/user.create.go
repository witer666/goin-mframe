package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/models/dao"
	"github.com/gin-gonic/models/service/page"
	"github.com/gin-gonic/utils"
)

type ActionUserCreate struct {
}

func (action *ActionUserCreate) Invoke(c *gin.Context) {
	pUser := page.Service_Page_Users{}
	user := &dao.Dao_Users{
		UserName: c.Query("uname"),
		Password: "xxxxxxxx",
	}
	res := pUser.InsertUser(user)
	bolRet := true
	if res < 1 {
		bolRet = false
	}
	utils.ReturnJson(c, bolRet, 0, "")
}
