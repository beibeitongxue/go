package settings_api

import (
	"Go_Work/global"
	"Go_Work/models/res"
	"github.com/gin-gonic/gin"
)

//	func (SettingsApi) SettingsInfoView(c *gin.Context) {
//		res.OkWithData(global.Config.SiteInfo, c)
//	}
type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示某一项的配置信息
// @Tags 系统管理
// @Summary 显示某一项的配置信息
// @Description 显示某一项的配置信息  site email qq qiniu jwt
// @Param name path string  true  "name"
// @Param token header string  true  "token"
// @Router /api/settings/{name} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (SettingsApi) SettingsInfoView(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

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
