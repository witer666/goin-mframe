package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/actions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/models/dao"
	"github.com/gin-gonic/models/service/page"
	"github.com/gin-gonic/utils"
)

type ActionAdminPost struct {
	actions.ActionBase
}

func (action *ActionAdminPost) Initialize() bool {
	return true
}

func (action *ActionAdminPost) Invoke(c *gin.Context) {
	var _daoPost dao.Dao_Post
	var _daoDbPost dao.Dao_DB_Post

	/*if err := c.ShouldBindJSON(&_daoPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}*/
	/*if err := c.BindQuery(&_daoPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}*/
	if err := c.Bind(&_daoPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/*if err := c.BindUri(&_daoPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var _arrLikes []string
	json.Unmarshal([]byte(_daoPost.Likes[0]), &_arrLikes)*/
	_daoDbPost.Id = _daoPost.Id
	_daoDbPost.UserName = _daoPost.UserName
	_daoDbPost.Sex = _daoPost.Sex
	_daoDbPost.Province = _daoPost.Province
	_likes, _ := json.Marshal(_daoPost.Likes) //_arrLikes
	_daoDbPost.Likes = string(_likes)
	pPost := page.Service_Page_Post{}
	pPost.InsertPost(&_daoDbPost)
	fmt.Println(_daoDbPost)
	utils.ReturnJson(c, nil, http.StatusOK, "")
}
