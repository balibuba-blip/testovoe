package app

import (
	"database/sql"
	"log"
	"testovoe/config"
	"testovoe/database"
	mservice "testovoe/internal/measures/service"
	mtransport "testovoe/internal/measures/transport"
	mhandlers "testovoe/internal/measures/transport/http/handlers"
	pservice "testovoe/internal/products/service"
	ptransport "testovoe/internal/products/transport"
	phandlers "testovoe/internal/products/transport/http/handlers"
	"testovoe/internal/repository"

	"github.com/gin-gonic/gin"
)

type App struct {
	Config *config.DBConfig
	DB     *sql.DB
	Router *gin.Engine
}

func Initialize() *App {
	cfg := config.Load()

	if err := database.InitDB(cfg.GetConnectionString()); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	app := &App{
		Config: cfg,
		DB:     database.GetDB(),
		Router: gin.Default(),
	}

	app.initializeRoutes()
	return app
}

func (a *App) initProducts() {
	log.Println("Initializing products routes...")
	repo := repository.NewRepository(a.DB)
	productService := pservice.NewService(repo.Product)
	productHandler := phandlers.NewHandler(productService)
	ptransport.NewRouter(productHandler).RegisterRoutes(a.Router)
}

func (a *App) initMeasures() {
	log.Println("Initializing measures routes...")
	repo := repository.NewRepository(a.DB)
	measureService := mservice.NewService(repo.Measure)
	measureHandler := mhandlers.NewHandler(measureService)
	mtransport.NewRouter(measureHandler).RegisterRoutes(a.Router)
}

func (a *App) initializeRoutes() {
	if a.DB == nil {
		log.Fatal("DB connection is nil")
	}

	a.initProducts()
	a.initMeasures()

	log.Println("All routes initialized successfully")
}

func (a *App) Close() {
	if err := database.CloseDB(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}
