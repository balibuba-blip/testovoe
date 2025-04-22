package config

import "os"

type DBConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSL_MODE string //не уверен что нужно
}

func Load() *DBConfig {
	return &DBConfig{
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     getEnv("DB_PORT", "5432"),
		DB_USER:     getEnv("DB_USER", "postgres"),
		DB_PASSWORD: getEnv("DB_PASSWORD", "gfhjkm777"),
		DB_NAME:     getEnv("DB_NAME", "testovik"),
		DB_SSL_MODE: getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
