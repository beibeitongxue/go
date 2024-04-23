package user_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/res"
	"Go_Work/utils/jwts"
	"Go_Work/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入密码"`
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`
}

// UserUpdatePassword 修改登陆人密码
func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithCode(res.UserNotExit, c)
		return
	}
	//判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithCode(res.PasswordError, c)
		return
	}
	Hashpwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("password", Hashpwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.UserNotExit, c)
		return
	}
	res.OkWithMessage("Update Success", c)
	return
}
