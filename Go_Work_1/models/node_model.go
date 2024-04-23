package models

type NodeModel struct {
	MODEL
	NodeName  string `gorm:"size:36" json:"node_name"`               // 节点名
	Country   string `gorm:"size:128" json:"country" `               // 节点国家
	Bandwidth string `gorm:"size:128" json:"bandwidth" `             // 带宽
	Bear      string `gorm:"size:18" json:"tel" `                    // 节点负荷
	Addr      string `gorm:"size:64" json:"addr,select(c|info)" `    // 节点地址
	Rate      string `gorm:"size:64" json:"token" `                  // 倍率
	Feature   string `gorm:"size:20" json:"ip,select(c)"`            // 线路标签
	Link      string `gorm:"size:128" json:"link,select(info)" `     // 我的链接地址
	Nodetype  string `gorm:"size:128" json:"nodetype,select(info)" ` //节点类型是否收费
}
