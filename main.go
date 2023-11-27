package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_redis/logging"
	"go_redis/persistence"
	"go_redis/src"
	"go_redis/strings_and_hash"
)

func examples() {
	strings_and_hash.SimpleSet()
	strings_and_hash.SetWithTTL()
	strings_and_hash.Incr()
	strings_and_hash.GetRange()
	strings_and_hash.HSet()
	strings_and_hash.HMSet()
	strings_and_hash.MSet()
	strings_and_hash.HGet()
	strings_and_hash.HMGet()
	strings_and_hash.HGetAll()
	strings_and_hash.Scan()
	strings_and_hash.HKeys()
}

var log = logging.GetLogger()
var redis = persistence.GetRedis()

func main() {
	//examples()
	cfg := src.GetConfig()

	log.Infof("Starting server on port %s ....", cfg.APP_PORT)

	r := gin.Default()
	r.GET("/users/:id", getUser)
	r.POST("users", addUser)

	r.Run(fmt.Sprintf(":%s", cfg.APP_PORT))
}

func addUser(c *gin.Context) {
	log.Info("POST")
}

func getUser(c *gin.Context) {
	log.Info("GET")
}
