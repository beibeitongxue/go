package settings_api

import (
	"Go_Work/global"
	"Go_Work/models/res"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示某一项的配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//查看对应配置项信息
	switch cr.Name {
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "system":
		res.OkWithData(global.Config.System, c)
	case "mysql":
		res.OkWithData(global.Config.Mysql, c)
	case "jwy":
		res.OkWithData(global.Config.Jwy, c)
	case "node":
		res.OkWithData(global.Config.Node, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
}
