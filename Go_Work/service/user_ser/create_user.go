package user_ser

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/ctype"
	"Go_Work/utils/pwd"
	"errors"
	"fmt"
)

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	// 判断用户名是否存在
	var userModel models.UserModel
	fmt.Println(userName)
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
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
		IP:       ip,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
