package config

import (
	"github.com/BurntSushi/toml"
)

// 系统配置
type appConfig struct {
	Address string
}

// 数据库配置
type mysqlConfig struct {
	IpPort       string
	UserPassword string
	Database     string
}

// rdis 配置
type redisConfig struct {
	Addr     string
	Password string
	DB       int
}

type TomlConfig struct {
	AppConfig   appConfig
	MysqlConfig mysqlConfig
	RedisConfig redisConfig
}

var Cfg *TomlConfig

func init() {
	Cfg = new(TomlConfig)

	_, err := toml.DecodeFile("app.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
