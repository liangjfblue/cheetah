package discovery

import (
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestNewRedisPool(t *testing.T) {
	redisPool := NewRedisPool("172.16.7.16:6379")

	conn := redisPool.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", "name", "liangjf"); err != nil {
		t.Fatal(err)
	}

	resp, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}
