package service

import (
	"fmt"
	"go-web-gin-ant/model"
	"go-web-gin-ant/utils"
)

// 登录
func Login(username, password string) error {
	db := utils.GetConnection()
	user, err := model.FindUserByUsername(db, username)
	if err != nil {
		return err
	}

	if user.Password != password {
		return fmt.Errorf("用户名不正确")
	}
	return nil
}