package dao

/**
 *
 * Post model类
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	void
 * @return  object
 *
 */
type Dao_Post struct {
	Id       int64
	UserName string   `form:"uname" json:"uname" uri:"uname"`
	Sex      string   `form:"sex" json:"sex" uri:"sex"`
	Likes    []string `form:"likes" json:"likes" uri:"likes"`
	Province string   `form:"province" json:"province" uri:"province"`
}

type Dao_DB_Post struct {
	Id       int64
	UserName string
	Sex      string
	Province string
	Likes    string
}

func (Dao_DB_Post) TableName() string {
	return "test.posts"
}
