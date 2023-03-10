package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 *
 * db操作封状类
 *
 * @author  linux_chen<linux_chen@163.com>
 * @version 2023/03/07 23:19:35
 * @param	void
 * @return  db object
 *
 */
type DB struct {
	Db *gorm.DB
}

func (mdb *DB) Open() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", "root", "123456", "127.0.0.1", 3306, "test", "10s")
	mdb.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return err
}

func (mdb *DB) Insert(m interface{}) (row int64) {
	mdb.Open()
	res := mdb.Db.Create(m)
	return res.RowsAffected
}

func (mdb *DB) Update(m interface{}, values interface{}) (row int64) {
	mdb.Open()
	res := mdb.Db.Model(m).Updates(values)
	return res.RowsAffected
}

func (mdb *DB) Delete(value interface{}) (row int64) {
	mdb.Open()
	res := mdb.Db.Delete(value)
	return res.RowsAffected
}

func (mdb *DB) GetList(m interface{}, offset int, limit int, order interface{}, conds ...interface{}) (err error) {
	mdb.Open()
	res := mdb.Db.Offset(offset).Limit(limit).Order(order).Find(m, conds...)
	return res.Error
}

func (mdb *DB) GetOnly(m interface{}, conds ...interface{}) (err error) {
	mdb.Open()
	res := mdb.Db.First(m, conds...)
	return res.Error
}

func (mdb *DB) Close() (err error) {
	err = mdb.Close()
	return err
}
