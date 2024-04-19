package main

import (
	"Go_Work/core"
	"Go_Work/flag"
	"Go_Work/global"
	"Go_Work/routers"
	"time"
)

// @title gvb_server API文档
// @version 1.0
// @description gvb_server API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	//前端和数据库8小时时差
	time.Local = time.FixedZone("CST", 8*3600)
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	//连接redis
	global.Redis = core.ConnectRedis()
	//绑定命令行参数
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Info("DB运行在：", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}

	//
	//core.InitAddrDB()
	//defer global.AddrDB.Close()

	// 命令行参数绑定

	// 连接es
	//global.ESClient = core.EsConnect()

	//cron_ser.CronInit()
	//
	//router := routers.InitRouter()
	//
	//addr := global.Config.System.Addr()
	////
	//utils.PrintSystem()
	//
	//err := router.Run(addr)
	//if err != nil {
	//	global.Log.Fatalf(err.Error())
	//}
	//// 不要乱提交
}
