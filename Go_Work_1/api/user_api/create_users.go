package user_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/res"
	"Go_Work/service"
	"Go_Work/service/user_ser"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ser163/WordBot/generate"
)

type UserCreate struct {
	DeviceID string `json:"device_id" binding:"required" msg:"请输入设备号"` //
}

// UserCreate  游客登录
func (UserApi) UserCreate(c *gin.Context) {
	var cr UserCreate
	//验证设备是否存在
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	device := service.ServiceApp.UserService.CheckDeviceID(cr.DeviceID)
	if device != nil {

		//入库用户信息
		w, err := generate.GenRandomWorld(8, "none")
		if err != nil {
			return
		}
		username := w.Word
		err = user_ser.UserService{}.CreateUser(username, "myname", cr.DeviceID, 4, "", cr.DeviceID)
		if err != nil {
			global.Log.Error(err)
			res.FailWithCode(res.CreatedError, c)
			return
		}
		res.OkWithMessage(fmt.Sprintf("User%s Creat Success!", username), c)
		return

	}

	var userModel models.UserModel
	err := global.DB.Take(&userModel, "device_id = ?", cr.DeviceID).Error
	if err != nil {
		res.FailWithError(err, &cr, c)
	}
	res.OkWithData(userModel, c)
	return

}
