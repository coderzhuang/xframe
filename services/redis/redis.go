package redis

import (
	"github.com/coderzhuang/core"
	"github.com/go-redis/redis/v8"
)

func New() *redis.Client {
	conf := core.Conf.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Auth,
		DB:       conf.DB,
	})
	return rdb
}
