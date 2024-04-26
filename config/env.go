package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Env          string
	Host         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	maxOpenConnsStr := getEnv("MAX_OPEN_CONN", "25")
	maxOpenConns, _ := strconv.Atoi(maxOpenConnsStr)

	maxIdleConnsStr := getEnv("MAX_IDLE_CONN", "25")
	maxIdleConns, _ := strconv.Atoi(maxIdleConnsStr)

	maxIdleTimeStr := getEnv("MAX_IDLE_TIME", "15")
	maxIdleTime, _ := strconv.Atoi(maxIdleTimeStr)

	return Config{
		Port:         getEnv("API_PORT", "4000"),
		Env:          getEnv("ENVIRONMENT", "development"),
		Host:         getEnv("HOST", "postgres://green:light@postgres:5432/greenlight?sslmode=disable"),
		MaxOpenConns: maxOpenConns,
		MaxIdleConns: maxIdleConns,
		MaxIdleTime:  time.Duration(maxIdleTime),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
