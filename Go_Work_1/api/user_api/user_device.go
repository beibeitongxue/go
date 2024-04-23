package user_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/res"
	"Go_Work/utils/jwts"
	"github.com/gin-gonic/gin"
)

type DeviceInfo struct {
	DeviceID string `json:"device_id"`
}

func (UserApi) GetDeviceID(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr DeviceInfo
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		return
	}
	err = global.DB.Model(&user).Update("addr", cr.DeviceID).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	res.OkWithMessage("获取成功", c)
	return
}
