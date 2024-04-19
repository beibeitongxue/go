package routers

import (
	"Go_Work/api"
	"Go_Work/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("user_login", app.LoginView)
	//	router.POST("login", app.QQLoginView)
	//	router.GET("qq_login_path", app.QQLoginLinkView) // qq登录的跳转地址
	router.POST("users", middleware.JwtAdmin(), app.UserCreateView)             //管理创建用户
	router.POST("user_create", middleware.JwtAuth(), app.UserCreate)            //创建用户
	router.GET("users", middleware.JwtAuth(), app.UserListView)                 //用户信息列表
	router.PUT("user_role", middleware.JwtAdmin(), app.UserUpdateRole)          //更新权限
	router.PUT("user_password", middleware.JwtAuth(), app.UserUdatePassword)    //更新密码
	router.POST("logout", middleware.JwtAuth(), app.LogoutView)                 //注销
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)           //删除用户
	router.POST("user_bind_email", middleware.JwtAuth(), app.UserBindEmailView) //绑定邮箱
	router.GET("user_info", middleware.JwtAuth(), app.UserInfoView)             //查看用户信息
	router.PUT("user_info", middleware.JwtAuth(), app.UserUpdateNickName)       //更改用户名称
}
