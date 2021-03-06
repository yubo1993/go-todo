package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
)

// Mysql 供外部调用
var Mysql *myDB

// myDB 封装一些数据库常用的操作
type myDB struct {
	*gorm.DB
}

type mysql struct {
	db *myDB //为了实现关闭数据库，所以在内部持有一个DB对象
}

func url(conf map[string]interface{}) string {
	user := cast.ToString(conf["username"])
	password := cast.ToString(conf["password"])
	host := cast.ToString(conf["host"])
	port := cast.ToInt(conf["port"])
	db := cast.ToString(conf["db"])
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8", user, password, host, port, db)
}

func (*mysql) Name() string {
	return "mysql"
}

func (m *mysql) Init(conf map[string]interface{}) error {
	db, err := gorm.Open(m.Name(), url(conf))
	if err != nil {
		return err
	}
	Mysql = &myDB{db}
	return nil
}

func (m *mysql) Close() {
	m.db.Close()
}

func (db *myDB) Insert(value interface{}) error {
	return db.Model(value).Create(value).Error
}

func (db *myDB) FindOne(result interface{}, where ...interface{}) error {
	return db.First(result, where...).Error
}
