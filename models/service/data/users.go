package data

import (
	"github.com/gin-gonic/library/log"
	"github.com/gin-gonic/models/dao"
)

type Service_Data_Users struct {
}

func (dUser *Service_Data_Users) GetUserList(offset int, limit int, order interface{}, conds ...interface{}) *[]dao.Dao_Users {
	mdb := dao.DB{}
	m := new([]dao.Dao_Users)
	err := mdb.GetList(&m, offset, limit, order, conds...)
	if err != nil {
		//添加日志
		logFields := map[string]interface{}{
			"conds": conds,
			"err":   err.Error(),
		}
		log.Init(nil).Debug(logFields)

		return m
	}

	return m
}

func (dUser *Service_Data_Users) InsertUser(user *dao.Dao_Users) int64 {
	mdb := dao.DB{}
	res := mdb.Insert(user)

	return res
}

func (dUser *Service_Data_Users) UpdateUser(user *dao.Dao_Users, values interface{}) int64 {
	mdb := dao.DB{}
	res := mdb.Update(user, values)

	return res
}

func (dUser *Service_Data_Users) GetUser(conds ...interface{}) *dao.Dao_Users {
	mdb := dao.DB{}
	m := new(dao.Dao_Users)
	err := mdb.GetOnly(m, conds)
	if err != nil {
		//添加日志
		logFields := map[string]interface{}{
			"conds": conds,
			"err":   err.Error(),
		}
		log.Init(nil).Debug(logFields)

		return m
	}

	return m
}
