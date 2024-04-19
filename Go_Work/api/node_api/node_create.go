package node_api

import (
	"Go_Work/core"
	"Go_Work/global"
	"Go_Work/models/res"
	"Go_Work/service/node_ser"
	"database/sql"
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

//const (
//	username = "root"
//	password = "2864//"
//	dbname   = "db"
//)

func (NodeApi) NodeCreate(c *gin.Context) {
	var request NodeCreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		res.FailWithError(err, &request, c)
		return
	}
	//验证节点是否存在
	err := node_ser.NodeService{}.CreateUser(request.NodeName)
	if err != nil {
		res.FailWithError(err, &request, c)
		return
	}
	core.InitConf()
	dsn := global.Config.Mysql.Dsn()
	//dsn := username + ":" + password + "@tcp(localhost:3306)/" + dbname + "?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		res.FailWithError(err, &request, c)
		return
	}
	defer db.Close()

	// 准备 SQL 语句
	stmt, err := db.Prepare("INSERT INTO node_models (nid, node_name, country, bandwidth, addr, rate, feature, nodetype) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		res.FailWithError(err, &request, c)
		return
	}
	defer stmt.Close()

	// 执行插入操作
	_, err = stmt.Exec(request.Nid, request.NodeName, request.Country, request.Bandwidth, request.Addr, request.Rate, request.Feature, request.Nodetype)
	if err != nil {
		res.FailWithError(err, &request, c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("节点%s创建成功!", request.NodeName), c)
	return

}
