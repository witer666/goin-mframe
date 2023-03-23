package data

import (
	"github.com/gin-gonic/models/dao"
)

type Service_Data_Post struct {
}

func (dPost *Service_Data_Post) InsertPost(post *dao.Dao_DB_Post) int64 {
	mdb := dao.DB{}
	res := mdb.Insert(post)

	return res
}
