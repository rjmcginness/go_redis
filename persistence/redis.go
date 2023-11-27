package persistence

import (
	"github.com/go-redis/redis"
	"go_redis/src"
	"time"
)

//type Value interface{}

var redis_client *redis.Client

func init() {
	redis_client = redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	cfg := src.GetConfig()
	for i := 0; i < cfg.REDIS_CONNECT_RETRIES; i++ {
		err := redis_client.Ping().Err()
		if err != nil {
			time.Sleep(3 * time.Second)
		} else {
			return
		}
	}
	panic("Unable to connect to Redis")
}

func GetRedis() *redis.Client {
	return redis_client
}

//func (v *Value) UnmarshalBinary(data []byte) error {
//	return json.Unmarshal(data, v)
//}
//
//func (v *Value) MarshalBinary() ([]byte, error) {
//	return json.Marshal(v)
//}
