package user_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/ctype"
	"Go_Work/models/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required, one of=1 2 3 4" msg:"权限参数错误"`
	NiceName string     `json:"nice_name" binding:"required"  ` //防止非法昵称
	UserName string     `json:"user_name" binding:"required" msg:"用户名错误" `
}

// UserUpdateRole 用户权限变更
func (UserApi) UserUpdateRole(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//验证用户信息
	var user models.UserModel
	err := global.DB.Take(&user, "user_name = ?", cr.UserName).Error
	if err != nil {
		res.FailWithCode(res.UserNotExit, c)
		return
	}
	//修改权限或昵称
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NiceName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.UpdateFailed, c)
		return
	}
	res.OkWithMessage("Update Success", c)
}
