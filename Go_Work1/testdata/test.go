package main

import (
	"Go_Work/models/ctype"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	//Creat    string `json:"creat"`
	ID       uint   `json:"id" gorm:"primary_key;column:id"`         //用户id
	NickName string `gorm:"size:36" json:"nick_name,select(c|info)"` // 昵称
	UserName string `gorm:"size:36" json:"user_name"`                // 用户名
	Password string `gorm:"size:128" json:"-"`                       // 密码
	Email    string `gorm:"size:128" json:"email"`                   // 邮箱
	Tel      string `gorm:"size:18" json:"tel"`                      // 手机号
	Addr     string `gorm:"size:64" json:"addr,select(c|info)"`      // 地址
	Token    string `gorm:"size:64" json:"token"`                    // 其他平台的唯一id
	//IP       string     `gorm:"size:20" json:"ip,select(c)"`               // ip地址
	Role     ctype.Role `gorm:"size:4;default:1" json:"role,select(info)"` // 权限  1 管理员  2 普通用户  3 游客
	Integral string     `gorm:"default:0" json:"integral,select(info)"`    // 我的积分
	Link     string     `gorm:"size:128" json:"link,select(info)"`         // 我的链接地址
	DeviceID string     `gorm:"size:128" json:"device_id,select(info)"`    //设备ID
}

func main() {
	// 连接到数据库
	db, err := sql.Open("mysql", "root:Root1234//@tcp(localhost:3306)/db+&parseTime=True&loc=Asia%2FShanghai")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
