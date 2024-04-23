package node_ser

import (
	"Go_Work/global"
	"Go_Work/models"
	"Go_Work/models/res"
)

func (NodeService) CreateNode(Nid, NodeName, Country, Bandwidth, Addr, Rate, Feature, Nodetype string) error {
	// 判断节点是否存在
	var nodeModel models.NodeModel
	err := global.DB.Take(&nodeModel, "node_name = ?", NodeName).Error
	if err == nil {
		err := res.NodeExit
		return err
	}

	err = global.DB.Create(&models.NodeModel{
		NodeName:  Nid,
		Country:   Country,
		Bandwidth: Bandwidth,
		//Bear:
		Addr:    Addr,
		Rate:    Rate,
		Feature: Feature,
		//Link:
		Nodetype: Nodetype,
	}).Error
	return err
}
