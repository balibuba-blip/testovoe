package main

import (
	"database/sql"
	"log"
	"testovoe/config"
	"testovoe/database"
	mrepo "testovoe/internal/measures/repository"
	mservice "testovoe/internal/measures/service"
	mtransport "testovoe/internal/measures/transport"
	mhandlers "testovoe/internal/measures/transport/http/handlers"
	prepo "testovoe/internal/products/repository"
	pservice "testovoe/internal/products/service"
	ptransport "testovoe/internal/products/transport"
	phandlers "testovoe/internal/products/transport/http/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	if err := database.InitDB(cfg.GetConnectionString()); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Получаем соединение через GetDB()
	db := database.GetDB()

	// Инициализация компонентов
	productRouter := initProductComponents(db)
	measureRouter := initMeasureComponents(db)

	// Настройка роутера
	router := setupRouter(productRouter, measureRouter)

	// Запуск сервера
	log.Printf("Starting server on port %s", cfg.DB_PORT)
	if err := router.Run(":" + cfg.DB_PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Инициализация продуктов
func initProductComponents(db *sql.DB) *ptransport.Router {
	repo := prepo.NewRepository(db)
	service := pservice.NewService(repo)
	handler := phandlers.NewHandler(service)
	return ptransport.NewRouter(handler)
}

// Инициализация мер
func initMeasureComponents(db *sql.DB) *mtransport.Router {
	repo := mrepo.NewRepository(db)
	service := mservice.NewService(repo)
	handler := mhandlers.NewHandler(service)
	return mtransport.NewRouter(handler)
}

// Настройка роутера
func setupRouter(productRouter *ptransport.Router, measureRouter *mtransport.Router) *gin.Engine {
	router := gin.Default()
	productRouter.RegisterRoutes(router)
	measureRouter.RegisterRoutes(router)
	return router
}
