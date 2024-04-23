package user_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/res"
	"Go_Work/utils/jwts"
	"github.com/fatih/structs"

	"github.com/gin-gonic/gin"
	"strings"
)

type UserUpdateNicknameRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Sign     string `json:"sign" structs:"sign"`
	Link     string `json:"link" structs:"link"`
}

// UserUpdateNickName 修改登陆人昵称
func (UserApi) UserUpdateNickName(c *gin.Context) {
	var cr UserUpdateNicknameRequest
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var newMaps = map[string]interface{}{}
	maps := structs.Map(cr)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[key] = val
		}
	}
	var userModel models.UserModel
	err = global.DB.Debug().Take(&userModel, claims.UserID).Error
	if err != nil {
		res.FailWithCode(res.UserNotExit, c)
		return
	}
	err = global.DB.Model(&userModel).Updates(newMaps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.UpdateFailed, c)
		return
	}
	res.OkWithMessage("Update Success", c)

}
