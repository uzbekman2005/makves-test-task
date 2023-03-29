package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production
	CtxTimeout  int    // context timeout in seconds
	LogLevel    string
	HTTPPort    string
	FilePath    string
}

// Load loads environment vars and inflates Config
func Load() Config {
	err := godotenv.Load("./config/.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	c := Config{}

	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(GetOrReturnDefault("HTTP_PORT", "8000"))
	c.FilePath = cast.ToString(GetOrReturnDefault("FILE_PATH", "./helper/file_storage/ueba.csv"))
	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
