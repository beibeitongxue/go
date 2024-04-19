package user_api

import (
	"Go_Work/global"
	"Go_Work/models/res"
	"Go_Work/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")
	//err := service.ServiceApp.UserService.Logout(claims, token)
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	println(diff)
	err := global.Redis.Set(fmt.Sprintf("logout_%s", token), "", diff).Err()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)

}
