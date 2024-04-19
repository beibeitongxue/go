package user_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/ctype"
	"Go_Work/models/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NiceName string     `json:"nice_name" binding:"required"  ` //防止非法昵称
	UserName string     `json:"user_name" binding:"required" msg:"用户名错误" `
}

// 用户权限变更
func (UserApi) UserUpdateRole(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, "user_name = ?", cr.UserName).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NiceName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改权限失败", c)
		return
	}
	res.OkWithMessage("修改权限成功", c)
}
