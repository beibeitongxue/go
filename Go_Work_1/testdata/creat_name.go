package main

import (
	"Go_Work/core"
	"Go_Work/global"
	"Go_Work/models"
	"fmt"
)

type UserCreateInfo struct {
	UserName string `json:"user_name"` // 用户名
	NickName string `json:"nick_name"` // 昵称
	Password string `json:"password"`  // 密码
	DeviceID string `json:"device_id"` //
}

func GetUserinfo() error {
	//var request UserService
	DeviceID := 1234
	//var userModel models.UserModel
	////验证节点是否存在
	//err := global.DB.Take(&userModel, "device_id=?", DeviceID).Error
	//if err == nil {
	//	return err
	//}
	core.InitConf()

	global.DB = core.InitGorm()
	var userModel models.UserModel
	cc := global.DB.Take(&userModel, "device_id = ?", DeviceID)
	fmt.Println(userModel)

	return cc.Error
}
func main() {
	a := GetUserinfo()
	fmt.Println(a)
}
