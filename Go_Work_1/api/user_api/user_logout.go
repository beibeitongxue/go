package user_api

import (
	"Go_Work/global"
	"Go_Work/models/res"
	"Go_Work/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// LogoutView 注销
func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	println(diff)
	err := global.Redis.Set(fmt.Sprintf("logout_%s", token), "", diff).Err()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("Logout fail", c)
		return
	}
	res.OkWithMessage("Logout Success", c)

}
