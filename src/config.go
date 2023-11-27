package src

import (
	"os"
	"strconv"
)

type Config struct {
	APP_PORT              string
	LOG_PATH              string
	REDIS_CONNECT_RETRIES int
}

func GetConfig() Config {
	appPort, ok := os.LookupEnv("APP_PORT")
	if !ok {
		appPort = "8089"
	}
	logPath, ok := os.LookupEnv("LOG_FILE")
	if !ok {
		logPath = "./logs/log.txt"
	}
	var redisConnectRetries int
	envRetries, ok := os.LookupEnv("REDIS_CONNECT_RETRIES")
	if !ok {
		redisConnectRetries = 10
	} else {
		var err error
		redisConnectRetries, err = strconv.Atoi(envRetries)
		if err == nil {
			redisConnectRetries = 10
		}
	}

	return Config{
		APP_PORT:              appPort,
		LOG_PATH:              logPath,
		REDIS_CONNECT_RETRIES: redisConnectRetries,
	}
}
