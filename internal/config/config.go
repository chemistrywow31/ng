package config

import (
	"os"
)

type Config struct {
	Server ServerConfig
	JWT    JWTConfig
}

type ServerConfig struct {
	Port string
	Mode string // debug, release, test
}

type JWTConfig struct {
	Secret string
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "dev-secret-change-in-prod"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
