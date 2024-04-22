package user_api

import (
	"Go_Work/global"
	"Go_Work/models/res"
	"Go_Work/service/user_ser"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ser163/WordBot/generate"
)

type UserCreateRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	NickName string `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	Password string `json:"password" binding:"required" msg:"请输入密码"`   // 密码
	//	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限  1 管理员  2 普通用户  3 游客
	DeviceID string `json:"device_id" binding:"required" msg:"请输入设备号"` //获取设备id
}

// UserCreateView 游客登陆创建用户
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	word, err := generate.GenRandomWorld(8, "none")
	cr.UserName = word.Word
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err = user_ser.UserService{}.CreateUser(cr.UserName, "", cr.DeviceID, 4, "", cr.DeviceID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("User%s Creat Succes!", cr.UserName), c)
	return
}
