package utils

import (
	"github.com/jinzhu/gorm"
	"log"
)

// 获取数据库链接
func GetConnection() *gorm.DB {
	return rootDB
}

// 数据库
var rootDB *gorm.DB

// 初始化数据库
func InitDB() error {
	db, err := gorm.Open("mysql",
		"root:admin@tcp(127.0.0.1:32768)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	rootDB = db
	return err
}

// 关闭数据库
func closeDB() {
	if rootDB != nil {
		err := rootDB.Close()
		log.Printf("Close gorm.db error, error=%+v", err)
	}
}
