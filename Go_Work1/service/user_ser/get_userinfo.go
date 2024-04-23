package user_ser

import (
	"Go_Work/core"
	"Go_Work/global"
	"Go_Work/models"
	"database/sql"
	"fmt"
)

type UserCreateInfo struct {
	UserName string `json:"user_name"` // 用户名
	NickName string `json:"nick_name"` // 昵称
	Password string `json:"password"`  // 密码
	DeviceID string `json:"device_id"` //
}

func (UserService) GetUserinfo(DeviceID string) error {
	//var request UserService
	var userModel models.UserModel
	//验证节点是否存在
	err := global.DB.Take(&userModel, "device_id=?", DeviceID).Error
	if err == nil {
		return err
	}
	core.InitConf()
	dsn := global.Config.Mysql.Dsn()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	sqlstr := "select * from user_models where id=?"
	// 准备 SQL 语句
	row := db.QueryRow(sqlstr, DeviceID)
	var (
		username string
		nickname string
		password string
		id       string
		email    string
		role     string
	)
	err = row.Scan(&username, &nickname, &password, &id, &email, &role /* 继续定义其他列 */)
	if err != nil {
		return err
	}
	defer db.Close()

	//// 执行插入操
	request := models.UserModel{
		UserName: username,
		Password: password,
		Email:    email,
	}
	fmt.Println(request)
	return err

}
