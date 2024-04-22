package user_ser

import (
	"Go_Work/global"
	"Go_Work/models"
)

func (UserService) CheckDeviceID(DeviceID string) error {
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "device_id=?", DeviceID).Error
	if err == nil {
		return err
	}
	return err
}
func (UserService) GetUserInfo(DeviceID string) error {
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "device_id=?", DeviceID).Error
	if err == nil {
		return err
	}
	return err
}
