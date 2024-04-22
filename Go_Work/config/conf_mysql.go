package config

import (
	"strconv"
)

type Mysql struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Config    string `yaml:"config"` // 高级配置，例如 charset
	DB        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	LogLevel  string `yaml:"log_level"` // 日志等级，debug就是输出全部sql，dev, release
	ParseTime string `yaml:"parse_time"`
}

func (m Mysql) Dsn() string {

	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config + "&parseTime=True&loc=Asia%2FShanghai"
}
