package redis

import (
	"github.com/go-redis/redis/v8"
	"xframe/config"
)

func New() *redis.Client {
	conf := config.Conf.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Auth,
		DB:       conf.DB,
	})
	return rdb
}
