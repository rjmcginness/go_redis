package strings_and_hash

import (
	"fmt"
	"go_redis/src/persistence"
	"time"
)

func SimpleSet() {
	redis := persistence.GetRedis()

	fmt.Println(redis.Set("foo", "bar", 0))
}

func SetWithTTL() {
	redis := persistence.GetRedis()

	// 2 second TTL
	fmt.Println(redis.Set("foo_short", "bar", 2*time.Second))

	// sleep 1 second
	time.Sleep(time.Second)
	fmt.Println(redis.Get("foo_short"))

	time.Sleep(2 * time.Second)
	fmt.Println(redis.Get("foo_short"))
}

func Incr() {
	redis := persistence.GetRedis()
	redis.Set("uid", 1, 0)
	redis.Incr("uid")
	redis.IncrBy("uid", 8)
	uid, err := redis.Get("uid").Result()

	if err != nil {
		fmt.Println("Get('uid') failed")
	}

	fmt.Println(uid)
}

func GetRange() {
	redis := persistence.GetRedis()

	redis.Set("long_str", "CoolKid McGinness", 0)
	name, _ := redis.GetRange("long_str", 0, 6).Result()
	fmt.Println(name)
}
