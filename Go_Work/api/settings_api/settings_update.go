package settings_api

import (
	"Go_Work/config"
	"Go_Work/core"
	"Go_Work/global"
	"Go_Work/models/res"
	"github.com/gin-gonic/gin"
)

//	func (SettingsApi) SettingsSiteInfoUpdateView(c *gin.Context) {
//		var cr config.SiteInfo
//		err := c.ShouldBindJSON(&cr)
//		if err != nil {
//			res.FailWithCode(res.ArgumentError, c)
//			return
//		}
//		//res.OkWithData(global.Config.SiteInfo, c)
//		fmt.Println("before", global.Config)
//		global.Config.SiteInfo = cr
//		fmt.Println("after", global.Config)
//		err = core.SetYaml()
//		if err != nil {
//			global.Log.Error(err)
//			res.FailWithMessage(err.Error(), c)
//			return
//		}
//		res.OkWith(c)
//	}
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = info
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

	core.SetYaml()
	res.OkWith(c)
}
