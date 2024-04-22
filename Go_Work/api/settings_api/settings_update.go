package settings_api

import (
	"Go_Work/config"
	"Go_Work/core"
	"Go_Work/global"
	"Go_Work/models/res"
	"github.com/gin-gonic/gin"
)

// SettingsInfoUpdateView 配置信息动态路由
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//查看对应配置信息
	switch cr.Name {
	case "site_info":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info
	case "system":
		var info config.System
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.System = info
	case "mysql":
		var info config.Mysql
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Mysql = info
	case "jwt":
		var info config.Jwy
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwy = info
	case "node":
		var info config.Node
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Node = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}

	err = core.SetYaml()
	if err != nil {
		return
	}
	res.OkWith(c)
}
