package redis

import (
	"manger/internal/model"

	"github.com/go-redis/redis"
)

// NewCounterRepo 创建一个API缓存存储服务
func NewCounterRepo(redis *redis.Client) model.CounterRepo {
	return &counterRepo{
		redis: redis,
	}
}

type counterRepo struct {
	redis *redis.Client
}

func (c *counterRepo) Key(name string) string {
	return "manger/counter:" + name
}

func (c *counterRepo) Count(counter *model.Counter) (int64, error) {
	key := c.Key(counter.Name)

	intCmd := c.redis.IncrBy(key, counter.Count)
	if err := intCmd.Err(); err != nil {
		return 0, err
	}

	if err := c.redis.Expire(key, counter.TTL).Err(); err != nil {
		// 尝试删除数据
		c.redis.Del(key)
		return 0, err
	}
	return intCmd.Val(), nil
}
