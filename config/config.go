package config

import (
	"fmt"
	"os"
)

const (
	DefaultDBPort     = "5432"
	DefaultServerPort = "8080"
)

type DBConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSL_MODE string //не уверен что нужно
}

func (c *DBConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DB_HOST,
		c.DB_PORT,
		c.DB_USER,
		c.DB_PASSWORD,
		c.DB_NAME,
		c.DB_SSL_MODE)
}

func Load() *DBConfig {
	return &DBConfig{
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     getEnv("DB_PORT", DefaultDBPort),
		DB_USER:     getEnv("DB_USER", "postgres"),
		DB_PASSWORD: getEnv("DB_PASSWORD", "gfhjkm777"),
		DB_NAME:     getEnv("DB_NAME", "testovik"),
		DB_SSL_MODE: getEnv("DB_SSLMODE", "disable"),
	}
}

func GetServerPort() string {
	return getEnv("SERVER_PORT", DefaultServerPort)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
