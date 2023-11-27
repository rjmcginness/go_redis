package strings_and_hash

import (
	"fmt"
	"go_redis/persistence"
)

func HSet() {
	redis := persistence.GetRedis()

	fmt.Println(redis.HSet("user:1", "name", "user1"))
}

func MSet() {
	redis := persistence.GetRedis()

	fmt.Println(redis.MSet("name", "user2", "email", "user2@example.com"))

	var kvPairs []string

	//kvPairs = make([]interface{}, ...)
	kvPairs = append(kvPairs, "name", "user2")
	kvPairs = append(kvPairs, "email", "user2@example.com")

	fmt.Println(redis.MSet(kvPairs))
}

func HMSet() {
	redis := persistence.GetRedis()

	data := make(map[string]interface{})
	data["name"] = "user3"
	data["email"] = "user3@example.com"
	fmt.Println(redis.HMSet("user:3", data))
}

func HGet() {
	redis := persistence.GetRedis()

	fmt.Println(redis.HGet("user:3", "name"))
}

func HMGet() {
	redis := persistence.GetRedis()

	fmt.Println(redis.HMGet("user:3", "name", "email"))
}

func HGetAll() {
	redis := persistence.GetRedis()

	fmt.Println(redis.HGetAll("user:3"))
}

func Scan() {
	redis := persistence.GetRedis()

	var cursor uint64

	keys, cursor, err := redis.Scan(cursor, "user:*", 10).Result()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(keys, cursor)
}

func HKeys() {
	redis := persistence.GetRedis()

	fmt.Println(redis.HKeys("user:3"))
}
