package main

import (
	"fmt"
	"log"
	"testovoe/internal/config"
	"testovoe/internal/repositories"
	"testovoe/internal/routes"
)

func main() {
	// Загрузка конфигурации
	cfg := config.Load()
	log.Printf("Starting with config: DB_HOST=%s, DB_NAME=%s", cfg.DB_HOST, cfg.DB_NAME)

	// Инициализация БД
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_SSL_MODE)

	if err := repositories.InitDB(connStr); err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	// Настройка роутера
	router := routes.SetupRouter()

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	router.Run(":5432")
}
