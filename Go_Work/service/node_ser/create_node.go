package node_ser

import (
	"Go_Work/global"
	"Go_Work/models"
	"errors"
	"fmt"
)

func (NodeService) CreateUser(Nodename string) error {
	// 判断用户名是否存在
	var nodeModel models.NodeModel
	fmt.Println(Nodename)
	err := global.DB.Take(&nodeModel, "node_name = ?", Nodename).Error
	if err == nil {
		return errors.New("节点名已存在")
	}
	return nil
}
