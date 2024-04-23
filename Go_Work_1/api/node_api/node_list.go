package node_api

import (
	"Go_Work/models"
	"Go_Work/models/ctype"
	"Go_Work/models/res"
	"Go_Work/service/common"
	"Go_Work/utils/jwts"
	"github.com/gin-gonic/gin"
)

//type UserResponse struct {
//	models.UserModel
//	RoleID int `json:"role_id"`
//}

type UserListRequest struct {
	models.PageInfo
	NodeName string `json:"node_name" form:"node_name"`
}

func (NodeApi) NodeList(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var page UserListRequest
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var Nodes []models.NodeModel

	Nodes, count, _ := common.ComList(models.NodeModel{NodeName: page.NodeName}, common.Option{
		PageInfo: page.PageInfo,
		Likes:    []string{"node_name"},
		Debug:    true,
	})
	if ctype.Role(claims.Role) == ctype.PermissionAdmin {
		res.OkWithList(Nodes, count, c)
	}
}
