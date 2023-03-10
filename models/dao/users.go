package dao

/**
 *
 * 用户model类
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	void
 * @return  object
 *
 */
type Dao_Users struct {
	Id       int64
	UserName string
	Password string
}

func (Dao_Users) TableName() string {
	return "test.users"
}
