package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl       string
	Port        string
	RabbitURL   string
	RabbitQueue string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		DBUrl: mustGetEnv("DATABASE_URL"),

		Port:        getEnvWithDefault("PORT", "8081"),
		RabbitURL:   mustGetEnv("RABBITMQ_URL"),
		RabbitQueue: getEnvWithDefault("RABBITMQ_QUEUE", "JobQueue"),
	}
}

// mustGetEnv fetches an env var or fatally exits if missing
func mustGetEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	log.Fatalf("Environment variable %s is required but not set", key)
	return ""
}
func getEnvWithDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return defaultValue
}
