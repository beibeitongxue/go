package main

import (
	"Go_Work/core"
	"Go_Work/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Level string

type LogStashModel struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	IP        string    `gorm:"size:32" json:"ip"`
	Addr      string    `gorm:"size:64" json:"addr"`
	Level     Level     `gorm:"size:4" json:"level"`     // 日志的等级
	Content   string    `gorm:"size:128" json:"content"` // 日志消息内容
	UserID    uint      `json:"user_id"`                 // 登录用户的用户id，需要自己在查询的时候做关联查询
}

func main() {
	core.InitConf()
	dsn := global.Config.Mysql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate will create the table based on the LogStashModel struct
	db.AutoMigrate(&LogStashModel{})
}
