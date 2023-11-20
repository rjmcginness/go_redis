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
	//redis := persistence.GetRedis()
	//
	//var kvPairs []interface{}

	//kvPairs = make()
	//
	//fmt.Println(redis.MSet("name": "user2", ))
}

func HMSet() {
	redis := persistence.GetRedis()

	data := make(map[string]interface{})
	data["name"] = "user3"
	data["email"] = "user3@example.com"
	fmt.Println(redis.HMSet("user:3", data))
}
