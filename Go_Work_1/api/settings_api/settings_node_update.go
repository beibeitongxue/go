package settings_api

import (
	"Go_Work/config"
	"Go_Work/core"
	"Go_Work/global"
	"Go_Work/models/res"

	"github.com/gin-gonic/gin"
)

// SettingsNodeInfoUpdateView 修改Node配置项信息
func (SettingsApi) SettingsNodeInfoUpdateView(c *gin.Context) {
	var cr config.Node
	//绑定输出
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	global.Config.Node = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
