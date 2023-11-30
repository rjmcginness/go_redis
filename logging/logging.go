package logging

import (
	"go.uber.org/zap"
	"go_redis/src"
	"log"
	"os"
	"strings"
)

var logger *zap.Logger

func createLogFile() (string, error) {
	cfg := src.GetConfig()

	idx := strings.LastIndex(cfg.LOG_PATH, "/")
	dir := cfg.LOG_PATH[:idx]
	var file *os.File

	_, err := os.Stat(dir)
	if os.IsNotExist(err) && len(dir) != 0 {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return cfg.LOG_PATH, err
		}
		file, err = os.Create(cfg.LOG_PATH)
	}
	defer file.Close()
	return cfg.LOG_PATH, err
}

func newLogger() (*zap.Logger, error) {
	logPath, err := createLogFile()
	if err != nil {
		panic(err)
	}
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		logPath,
	}
	return cfg.Build()
}

func init() {
	loggr, err := newLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger = loggr
}

func GetLogger() *zap.SugaredLogger {
	return logger.Sugar()
}
