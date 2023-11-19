package persistence

import (
	"github.com/go-redis/redis"
)

var redis_client: redis

func init() {
	redis_client  = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
}

func connect() {
	redis.
}
