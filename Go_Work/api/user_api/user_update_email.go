package user_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/res"
	"Go_Work/service/user_ser"
	"Go_Work/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UpdateEmailRequest struct {
	//OldEmail string `json:"old_email" binding:"required" msg:"请输入邮箱"`
	Email string `json:"email" binding:"required,email" msg:"邮箱非法"`
}

// UserUpdateEmail 修改登陆人邮箱
func (UserApi) UserUpdateEmail(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdateEmailRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	err = user_ser.UserService{}.CheckEmail(cr.Email)
	if err != nil {
		res.FailWithCode(res.EmailExit, c)
		return
	}
	err = global.DB.Model(&user).Update("email", cr.Email).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("Update Success", c)
	return
}
