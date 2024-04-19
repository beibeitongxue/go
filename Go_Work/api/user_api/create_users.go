package user_api

import (
	"Go_Work/global"
	"Go_Work/models/res"
	"Go_Work/service/user_ser"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreate struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	NickName string `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	Password string `json:"password" binding:"required" msg:"请输入密码"`   // 密码
}

// UserCreateView 创建用户
func (UserApi) UserCreate(c *gin.Context) {
	var cr UserCreate
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err := user_ser.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, 4, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.UserName), c)
	return
}
