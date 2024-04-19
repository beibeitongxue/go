package routers

import (
	"Go_Work/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//路由分组
	apiRouterGroup := router.Group("api")

	routerGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置api
	routerGroupApp.SettingsRouter()
	routerGroupApp.NodeRouter()
	//routerGroupApp.AdvertRouter()
	//routerGroupApp.MenuRouter()
	routerGroupApp.UserRouter()
	//routerGroupApp.TagRouter()
	//routerGroupApp.MessageRouter()
	//routerGroupApp.ArticleRouter()
	//routerGroupApp.CommentRouter()
	//routerGroupApp.NewsRouter()
	//routerGroupApp.ChatRouter()
	//routerGroupApp.LogRouter()
	//routerGroupApp.DataRouter()
	return router
}
