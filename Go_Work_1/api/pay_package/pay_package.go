package pay_package

//
//import (
//	"Go_Work/global"
//	"Go_Work/models"
//	"Go_Work/models/res"
//	"Go_Work/utils/jwts"
//	"github.com/gin-gonic/gin"
//)
//
//type UserApi struct {
//}
//type UpdatePasswordRequest struct {
//	OldDate string // 剩余流量
//	Date    string // 购买流量包
//}
//
//func (UserApi) UserAddPackge(c *gin.Context) {
//	_claims, _ := c.Get("claims")
//	claims := _claims.(*jwts.CustomClaims)
//	var adddate UpdatePasswordRequest
//	if err := c.ShouldBindJSON(&adddate); err != nil {
//		res.FailWithError(err, &adddate, c)
//		return
//	}
//	var user models.UserModel
//	err := global.DB.Take(&user, claims.UserID).Error
//	if err != nil {
//		res.FailWithMessage("用户不存在", c)
//		return
//	}
//	err = global.DB.Model(&user).Updates(map[string]any{
//		"role": adddate.Date,
//	}).Error
//	if err != nil {
//		global.Log.Error(err)
//		res.FailWithMessage("流量包购买失败", c)
//		return
//	}
//	res.OkWithMessage("流量包购买成功", c)
//}
