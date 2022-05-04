package pkg

import (
	"github.com/go-redis/redis/v8"
	"test/config"
)

func NewRedis() *redis.Client {
	conf := config.Conf.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Auth,
		DB:       conf.DB,
	})
	return rdb
}
