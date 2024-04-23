package node_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/res"
	"github.com/gin-gonic/gin"
)

type NodeInfoRequest struct {
	NodeName string `json:"node_name" binding:"required" msg:"请输入节点名"`
}

func (NodeApi) NodeInfoView(c *gin.Context) {

	//_claims, _ := c.Get("claims")
	//claims := _claims.(*jwts.CustomClaims)
	var nodeInfo models.NodeModel
	if err := c.ShouldBindJSON(&nodeInfo); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	err := global.DB.Take(&nodeInfo, "node_name = ?", nodeInfo.NodeName).Error
	if err != nil {
		res.FailWithCode(res.UserNotExit, c)
		return
	}
	res.OkWithData(nodeInfo, c)

}
