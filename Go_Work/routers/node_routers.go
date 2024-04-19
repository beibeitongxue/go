package routers

import "Go_Work/api"

func (router RouterGroup) NodeRouter() {
	Node := api.ApiGroupApp.NodeApi
	router.POST("node/create", Node.NodeCreate)
}
