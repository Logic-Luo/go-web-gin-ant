package controller

import "go-web-gin-ant/service"

// 登录
func Login(username, password string) error {
	return service.Login(username, password)
}
