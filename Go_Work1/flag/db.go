package flag

import (
	"Go_Work/global"
	"Go_Work/models"
)

func Makemigrations() {

	//var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "NodeModel", &models.NodeModel{})
	//if err != nil {
	//	return
	//}
	//global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成四张表的表结构
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.NodeModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
