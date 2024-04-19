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
		// 没找到
		global.Log.Warn("用户名不存在")
		log.Warn(fmt.Sprintf("%s 用户名不存在", cr.UserName))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		log.Warn(fmt.Sprintf("用户名密码错误 %s %s", cr.UserName, cr.Password))
		res.FailWithMessage("用户名或密码错误", c)
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
		log.Error(fmt.Sprintf("token生成失败 %s", err.Error()))
		res.FailWithMessage("token生成失败", c)
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
