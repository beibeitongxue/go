package user_ser

import (
	"Go_Work/global"
	"Go_Work/models"
	"errors"
)

// CheckEmail 检测邮箱是否已存在
func (UserService) CheckEmail(Email string) error {
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "email=?", Email).Error
	if err == nil {
		return errors.New("email exit")
	}
	return nil
}
