package app

import (
	"database/sql"
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

	// Инициализация БД
	if err := database.InitDB(cfg.GetConnectionString()); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	app := &App{
		Config: cfg,
		DB:     database.GetDB(),
		Router: gin.Default(),
	}

	app.initializeRoutes()

	return app
}

func (a *App) initializeRoutes() {
	// Инициализация компонентов продуктов
	productRepo := prepo.NewRepository(a.DB)
	productService := pservice.NewService(productRepo)
	productHandler := phandlers.NewHandler(productService)
	productRouter := ptransport.NewRouter(productHandler)
	productRouter.RegisterRoutes(a.Router)

	// Инициализация компонентов мер
	measureRepo := mrepo.NewRepository(a.DB)
	measureService := mservice.NewService(measureRepo)
	measureHandler := mhandlers.NewHandler(measureService)
	measureRouter := mtransport.NewRouter(measureHandler)
	measureRouter.RegisterRoutes(a.Router)
}

func (a *App) Close() {
	if err := database.CloseDB(); err != nil {
		panic("Failed to close database: " + err.Error())
	}
}
