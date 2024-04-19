package routers

import (
	"Go_Work/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings/site", settingsApi.SettingsSiteInfoView)
	router.PUT("settings/site", settingsApi.SettingsSiteUpdateView)
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
}
