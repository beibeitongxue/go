package settings_api

import (
	"Go_Work/global"
	"Go_Work/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsNodeInfoView(c *gin.Context) {
	res.OkWithData(global.Config.Node, c)
}
