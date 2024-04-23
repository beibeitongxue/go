package routers

import (
	"Go_Work/api"
	"Go_Work/middleware"
)

func (router RouterGroup) NodeRouter() {
	Node := api.ApiGroupApp.NodeApi
	router.POST("node_create", middleware.JwtAdmin(), Node.NodeCreate)
	router.PUT("node_update", middleware.JwtAdmin(), Node.NodeUpdate)
	router.GET("nodes", middleware.JwtAdmin(), Node.NodeList)
	router.GET("node_info", middleware.JwtAdmin(), Node.NodeInfoView)
}
