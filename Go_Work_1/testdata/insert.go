package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//type NodeCreateRequest struct {
//	Nid       string ` json:"nid" binding:"required" msg:"请输入节点ID"`      // 节点ID
//	NodeName  string ` json:"node_name" binding:"required" msg:"请输入节点名"` // 节点名
//	Country   string `  json:"country" binding:"required" msg:"请输入落地国家"` // 节点国家
//	Bandwidth string `  json:"bandwidth" binding:"required" msg:"请输入带宽"` // 带宽
//	//	Bear      string ` json:"tel"`                                                  // 节点负荷
//	Addr    string ` json:"addr" binding:"required" msg:"请输入节点地址"` // 节点地址
//	Rate    string ` json:"rate" binding:"required" msg:"请输入倍率"`   // 倍率
//	Feature string ` json:"ip" binding:"required" msg:"请输入节点标签"`   // 线路标签
//	//	Link      string ` json:"link,select(info)" binding:"required" msg:"请输入节点名"`     // 我的链接地址
//	Nodetype string ` json:"nodetype" binding:"required" msg:"节点是否收费"` // 权限  1 管理员  2 普通用户  3 游客
//}

type NodeCreateRequest struct {
	Nid       string ` json:"nid" binding:"required" msg:"请输入节点ID"`      // 节点ID
	NodeName  string ` json:"node_name" binding:"required" msg:"请输入节点名"` // 节点名
	Country   string `  json:"country" binding:"required" msg:"请输入落地国家"` // 节点国家
	Bandwidth string `  json:"bandwidth" binding:"required" msg:"请输入带宽"` // 带宽
	//	Bear      string ` json:"tel"`                                                  // 节点负荷
	Addr    string ` json:"addr" binding:"required" msg:"请输入节点地址"` // 节点地址
	Rate    string ` json:"rate" binding:"required" msg:"请输入倍率"`   // 倍率
	Feature string ` json:"ip" binding:"required" msg:"请输入节点标签"`   // 线路标签
	//	Link      string ` json:"link,select(info)" binding:"required" msg:"请输入节点名"`     // 我的链接地址
	Nodetype string ` json:"nodetype" binding:"required" msg:"节点是否收费"` // 权限  1 管理员  2 普通用户  3 游客
}

const (
	username = "root"
	password = "2864//"
	dbname   = "db"
)

func NodeCreate1(request NodeCreateRequest) error {

	dsn := username + ":" + password + "@tcp(localhost:3306)/" + dbname + "?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// 准备 SQL 语句
	stmt, err := db.Prepare("INSERT INTO node_models (nid, node_name, country, bandwidth, addr, rate, feature, nodetype) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 执行插入操作
	_, err = stmt.Exec(request.Nid, request.NodeName, request.Country, request.Bandwidth, request.Addr, request.Rate, request.Feature, request.Nodetype)
	if err != nil {
		return err
	}

	return nil

}
func main1() {
	//crr := NodeCreateRequest{
	//	Nid:       "1",
	//	NodeName:  "riben",
	//	Country:   "riebn",
	//	Bandwidth: "10M",
	//	Addr:      "127.0.0.1",
	//	Rate:      "sda",
	//	Feature:   "1233",
	//	Nodetype:  "收费",
	//}
	////c := NodeCreate(crr)
	//fmt.Println(c)
}
