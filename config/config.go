package config

import (
	"short-it/internal/logger"
	"sync"

	"github.com/joho/godotenv"
)

var (
	server_config map[string]string
	once          sync.Once
)

func init() {
	once.Do(func() {
		config, err := godotenv.Read()
		if err != nil {
			logger.Error("Error loading .env file")
			panic(err)
		}
		server_config = config
	})
}

func Get(key string) string {
	if _, ok := server_config[key]; !ok {
		logger.Error("Key not found in server config")
		panic("Key not found in server config")
	}
	return server_config[key]
}
