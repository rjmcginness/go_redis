package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_redis/src"
	"go_redis/src/logging"
	"go_redis/src/model"
	"go_redis/src/persistence"
	strings_and_hash2 "go_redis/src/strings_and_hash"
	"net/http"
)

func examples() {
	strings_and_hash2.SimpleSet()
	strings_and_hash2.SetWithTTL()
	strings_and_hash2.Incr()
	strings_and_hash2.GetRange()
	strings_and_hash2.HSet()
	strings_and_hash2.HMSet()
	strings_and_hash2.MSet()
	strings_and_hash2.HGet()
	strings_and_hash2.HMGet()
	strings_and_hash2.HGetAll()
	strings_and_hash2.Scan()
	strings_and_hash2.HKeys()
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
	req := c.Request
	var user model.User
	json.NewDecoder(req.Body).Decode(&user)

	newId := redis.Incr("userId").Val()

	err := redis.HSet(fmt.Sprintf("user:%d", newId), "name", user.Name).Err()
	if err != nil {
		log.Errorf("Add User: %#v with id %d - %#v", user, newId, err)
		c.JSON(http.StatusInternalServerError, user)
	} else {
		c.JSON(http.StatusOK, gin.H{"created": "SUCCESS"})
	}
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	strCmd := redis.HGet(fmt.Sprintf("user:%s", id), "name")
	if strCmd.Err() != nil {
		log.Errorf("Get User: id %s - %#v", id, strCmd.Err())
		c.JSON(http.StatusNoContent, gin.H{"id": id})
	} else {
		c.JSON(http.StatusOK, gin.H{"name": strCmd.Val()})
	}
}
