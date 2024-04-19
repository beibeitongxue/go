package settings_api

import (
	"Go_Work/config"
	"Go_Work/core"
	"Go_Work/global"
	"Go_Work/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsNodeInfoUpdateView(c *gin.Context) {
	var cr config.Node
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//res.OkWithData(global.Config.SiteInfo, c)
	global.Config.Node = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
