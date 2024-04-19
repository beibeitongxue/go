package user_api

import (
	"Go_Work/models"
	"Go_Work/models/ctype"
	"Go_Work/models/res"
	"Go_Work/service/common"
	"Go_Work/utils/desens"
	"Go_Work/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	models.UserModel
	RoleID int `json:"role_id"`
}

type UserListRequest struct {
	models.PageInfo
	Role int `json:"role" form:"role"`
}

func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var page UserListRequest
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []UserResponse
	list, count, _ := common.ComList(models.UserModel{Role: ctype.Role(page.Role)}, common.Option{
		PageInfo: page.PageInfo,
		Likes:    []string{"nick_name"},
	})

	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员
			user.UserName = ""
			user.Tel = desens.DesensitizationTel(user.Tel)
			user.Email = desens.DesensitizationEmail(user.Email)
		}

		// 脱敏
		users = append(users, UserResponse{
			UserModel: user,
			RoleID:    int(user.Role),
		})
	}

	res.OkWithList(users, count, c)
}

//func (UserApi) UserListView(c *gin.Context) {
//
//	token := c.Request.Header.Get("token")
//	if token == "" {
//		res.FailWithMessage("未携带token", c)
//		return
//	}
//	claims, err := jwts.ParseToken(token)
//	if err != nil {
//		res.FailWithMessage("token错误", c)
//		return
//	}
//	var page UserListRequest
//	var users []UserResponse
//	list, count, _ := common.ComList(models.UserModel{Role: ctype.Role(page.Role)}, common.Option{
//		PageInfo: page.PageInfo,
//		Likes:    []string{"nick_name"},
//	})
//
//	for _, user := range list {
//		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
//			// 管理员
//			user.UserName = ""
//		}
//		// 脱敏
//		user.Tel = desens.DesensitizationTel(user.Tel)
//		//user.Email = desens.DesensitizationEmail(user.Email)
//		users = append(users, UserResponse{
//			UserModel: user,
//			RoleID:    int(user.Role),
//		})
//	}
//
//	res.OkWithList(users, count, c)
//}
