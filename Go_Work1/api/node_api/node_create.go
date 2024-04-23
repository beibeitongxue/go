package node_api

import (
	"Go_Work/global"
	"Go_Work/models/res"
	"Go_Work/service/node_ser"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type NodeCreateRequest struct {
	Nid       string `json:"nid" binding:"required" msg:"请输入节点ID"`      // 节点ID
	NodeName  string `json:"node_name" binding:"required" msg:"请输入节点名"` // 节点名
	Country   string `json:"country" binding:"required" msg:"请输入落地国家"`  // 节点国家
	Bandwidth string `json:"bandwidth" binding:"required" msg:"请输入带宽"`  // 带宽
	//	Bear      string ` json:"tel"`                                                  // 节点负荷
	Addr    string `json:"addr" binding:"required" msg:"请输入节点地址"`    // 节点地址
	Rate    string `json:"rate" binding:"required" msg:"请输入倍率"`      // 倍率
	Feature string `json:"feature" binding:"required" msg:"请输入节点标签"` // 线路标签
	//	Link      string ` json:"link,select(info)" binding:"required" msg:"请输入节点名"`     // 我的链接地址
	Nodetype string `json:"nodetype" binding:"required" msg:"节点是否收费"` // 权限  1 管理员  2 普通用户  3 游客
}

// NodeCreate 创建新节点至数据库
func (NodeApi) NodeCreate(c *gin.Context) {
	var request NodeCreateRequest

	var cr NodeCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	err := node_ser.NodeService{}.CreateNode(cr.Nid, cr.NodeName, cr.Country, cr.Bandwidth, cr.Addr, cr.Rate, cr.Feature, cr.Nodetype)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.CreatedError, c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("Node Cteat Success!", request.NodeName), c)
	return

}
