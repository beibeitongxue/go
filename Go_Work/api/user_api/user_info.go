package user_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/res"
	"Go_Work/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (UserApi) UserInfoView(c *gin.Context) {

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var userInfo models.UserModel
	err := global.DB.Take(&userInfo, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	res.OkWithData(filter.Select("info", userInfo), c)

}
