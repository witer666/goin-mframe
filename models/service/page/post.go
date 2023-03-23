package page

import (
	"github.com/gin-gonic/models/dao"
	"github.com/gin-gonic/models/service/data"
)

type Service_Page_Post struct {
}

/**
 *
 * 插入post数据
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	post dao.Dao_Post post数据
 * @return  int64 影响行数
 *
 */
func (pPost *Service_Page_Post) InsertPost(post *dao.Dao_DB_Post) int64 {
	dPost := data.Service_Data_Post{}
	res := dPost.InsertPost(post)
	return res
}
