package app

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
	productRepo := prepo.NewRepository(a.DB)
	productService := pservice.NewService(productRepo)
	productHandler := phandlers.NewHandler(productService)
	ptransport.NewRouter(productHandler).RegisterRoutes(a.Router)
}

func (a *App) initMeasures() {
	log.Println("Initializing measures routes...")
	measureRepo := mrepo.NewRepository(a.DB)
	measureService := mservice.NewService(measureRepo)
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
