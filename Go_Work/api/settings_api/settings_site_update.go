package settings_api

import (
	"Go_Work/config"
	"Go_Work/core"
	"Go_Work/global"
	"Go_Work/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsSiteUpdateView(c *gin.Context) {
	var info config.SiteInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = info
	core.SetYaml()
	res.OkWithMessage("网站信息更新成功", c)
}
