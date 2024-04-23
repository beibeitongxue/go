package user_ser

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/ctype"
	"Go_Work/utils/pwd"
	"github.com/ser163/WordBot/generate"
)

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, deviceId string) error {
	// 判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name=?", userName).Error
	if err == nil {
		w, _ := generate.GenRandomWorld(10, "none")
		userName = w.Word
	}
	// 对密码进行hash
	hashPwd := pwd.HashPwd(password)

	// 入库
	err = global.DB.Create(&models.UserModel{
		NickName: nickName,
		UserName: userName,
		Password: hashPwd,
		Email:    email,
		Role:     role,
		//IP:       ip,
		DeviceID: deviceId,
		Link:     userName,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
