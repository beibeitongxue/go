package service

import (
	"Go_Work/service/node_ser"
	"Go_Work/service/user_ser"
)

type ServiceGroup struct {
	//ImageService image_ser.ImageService
	UserService user_ser.UserService
	NodeService node_ser.NodeService
}

var ServiceApp = new(ServiceGroup)
