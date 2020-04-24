package discovery

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisPool struct {
	pool *redis.Pool
}

func NewRedisPool(host string) *RedisPool {
	return &RedisPool{
		pool: &redis.Pool{
			MaxIdle:     100,
			MaxActive:   50,
			IdleTimeout: 60 * time.Second,

			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", host)
				if err != nil {
					return nil, err
				}
				return c, err
			},

			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		},
	}
}
