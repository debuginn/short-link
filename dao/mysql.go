package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// 连接数据库
func InitDB() (err error) {
	driveName := "mysql"
	host := "localhost"
	port := "3306"
	database := "goshortlink"
	username := "root"
	password := "root"
	charset := "utf8"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	DB, err = gorm.Open(driveName, args)
	if err != nil {
		return
	}

	return DB.DB().Ping()
}

// 关闭 DB 实例方法
func CloseDB() {
	DB.Close()
}
