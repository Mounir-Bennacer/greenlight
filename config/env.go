package config

import "os"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}
type Config struct {
	Port string
	Env  string
	Host string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port: getEnv("API_PORT", "4000"),
		Env:  getEnv("ENVIRONMENT", "development"),
		Host: getEnv("HOST", "postgresql://admin:admin@localhost:5432/greenlight?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
