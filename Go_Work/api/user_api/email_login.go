package user_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/res"
	"Go_Work/plugins/log_stash"
	"Go_Work/utils"
	"Go_Work/utils/jwts"
	"Go_Work/utils/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// LoginView 登陆界面
func (UserApi) LoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	log := log_stash.NewLogByGin(c)
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 账号不存在
		global.Log.Warn("UserName Not Exist")
		log.Warn(fmt.Sprintf("%s UserName Not Exist", cr.UserName))
		res.FailWithCode(res.UserNameError, c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("Password Not Match")
		log.Warn(fmt.Sprintf("Password Not Match %s %s", cr.UserName, cr.Password))
		res.FailWithCode(res.UserNameError, c)
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		UserName: userModel.UserName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		log.Error(fmt.Sprintf("Token Build failure %s", err.Error()))
		res.FailWithCode(res.TokenslideError, c)
		return
	}
	ip, addr := utils.GetAddrByGin(c)
	log = log_stash.New(ip, token)
	log.Info("登录成功")

	global.DB.Create(&models.LoginDataModel{
		UserID:   userModel.ID,
		IP:       ip,
		NickName: userModel.NickName,
		Token:    token,
		Device:   "",
		Addr:     addr,
	})

	res.OkWithData(token, c)

}
