package redis

import (
	"manger/pkg/misc/config"

	"github.com/go-redis/redis"
)

// New 创建redis连接池
func New(conf *config.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})

	err := client.Ping().Err()
	return client, err
}
