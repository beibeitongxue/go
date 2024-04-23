package node_api

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strings"
)

type NodeUpdateRequest struct {
	Nid       string `json:"nid" binding:"required" msg:"请输入节点ID"`      // 节点ID
	NodeName  string `json:"node_name" binding:"required" msg:"请输入节点名"` // 节点名
	Country   string `json:"country" binding:"required" msg:"请输入落地国家"`  // 节点国家
	Bandwidth string `json:"bandwidth" binding:"required" msg:"请输入带宽"`  // 带宽
	//	Bear      string ` json:"tel"`                                                  // 节点负荷
	Addr     string `json:"addr" binding:"required" msg:"请输入节点地址"`    // 节点地址
	Rate     string `json:"rate" binding:"required" msg:"请输入倍率"`      // 倍率
	Feature  string `json:"feature" binding:"required" msg:"请输入节点标签"` // 线路标签
	Nodetype string `json:"nodetype" binding:"required" msg:"节点是否收费"` // 权限  1 管理员  2 普通用户  3 游客
}

func (NodeApi) NodeUpdate(c *gin.Context) {

	var cr NodeUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var request models.NodeModel
	err := global.DB.Take(&request, "node_name = ?", cr.NodeName).Error
	if err != nil {
		res.FailWithCode(res.NodeNotFound, c)
		return
	}
	var newMaps = map[string]interface{}{}
	maps := structs.Map(cr)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[key] = val
		}
	}
	var nodeModel models.NodeModel

	err = global.DB.Model(&nodeModel).Where("node_name = ?", cr.NodeName).Updates(newMaps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.UpdateFailed, c)
		return
	}
	res.OkWithData("Update Success", c)
	return

}
