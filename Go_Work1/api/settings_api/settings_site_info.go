package settings_api

import (
	"Go_Work/global"
	"Go_Work/models/res"
	"github.com/gin-gonic/gin"
)

// SettingsSiteInfoView 显示网站信息
func (SettingsApi) SettingsSiteInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
