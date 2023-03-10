package page

import (
	"github.com/gin-gonic/models/dao"
	"github.com/gin-gonic/models/service/data"
)

type Service_Page_Users struct {
}

/**
 *
 * 获取用户列表数据
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	conds interface 查询条件
 * @return  dao.Dao_Users
 *
 */
func (pUser *Service_Page_Users) GetUserList(conds ...interface{}) *[]dao.Dao_Users {
	dUser := data.Service_Data_Users{}
	m := dUser.GetUserList(0, 3, "id DESC", conds...)
	return m
}

/**
 *
 * 插入用户数据
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	user dao.Dao_Users 用户数据
 * @return  int64 影响行数
 *
 */
func (pUser *Service_Page_Users) InsertUser(user *dao.Dao_Users) int64 {
	dUser := data.Service_Data_Users{}
	res := dUser.InsertUser(user)
	return res
}

/**
 *
 * 插入用户数据
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	user dao.Dao_Users 用户数据
 * @param	values interface 更新数据
 * @return  int64 影响行数
 *
 */
func (pUser *Service_Page_Users) UpdateUser(user *dao.Dao_Users, values interface{}) int64 {
	dUser := data.Service_Data_Users{}
	res := dUser.UpdateUser(user, values)
	return res
}

/**
 *
 * 根据主键Id获取用户数据
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	id int64 主键id
 * @return  dao.Dao_Users
 *
 */
func (pUser *Service_Page_Users) GetUserById(id int64) *dao.Dao_Users {
	dUser := data.Service_Data_Users{}
	m := dUser.GetUser("id = ?", id)
	return m
}
