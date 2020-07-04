package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// 用户表
type User struct {
	// 主键ID
	Id uint64
	// 用户名
	Username string
	// 密码
	Password string
	// 创建时间
	CreatedAt time.Time
	// 最后更新时间
	UpdatedAt time.Time
}

// gorm 获取表名
func (n User) TableName() string {
	return "t_user"
}

// 保存用户信息
func SaveUser(db *gorm.DB, user User) error {
	return db.Save(&user).Error
}

// 根据用户名查询用户信息
func FindUserByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return User{}, fmt.Errorf("find user by username error, username=%v, error=%+v", username, err)
	}
	return user, nil
}
