package utils

import "log"

// 初始化信息
func init() {
	// 初始化数据库信息
	err := InitDB()
	if err != nil {
		log.Printf("init db error, error=%+v", err)
		panic("init db error")
	}
	// 关闭数据库
	defer closeDB()
}
