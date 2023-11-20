package main

import (
	"go_redis/strings_and_hash"
)

func main() {
	strings_and_hash.SimpleSet()
	strings_and_hash.SetWithTTL()
	strings_and_hash.Incr()
	strings_and_hash.GetRange()
	strings_and_hash.HSet()
	strings_and_hash.HMSet()
}
