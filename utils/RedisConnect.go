package utils

import (
	"atao-go-blog/config"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func init() {
	redisconfig := &config.Cfg.RedisConfig
	RDB = redis.NewClient(&redis.Options{
		Addr:     redisconfig.Addr,
		Password: redisconfig.Password, // no password set
		DB:       redisconfig.DB,       // use default DB
	})
}
