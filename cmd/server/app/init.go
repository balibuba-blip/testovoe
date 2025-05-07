package app

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"testovoe/config"
	"testovoe/database"
	"testovoe/internal/interfaces"
	mservice "testovoe/internal/measures/service"
	mtransport "testovoe/internal/measures/transport"
	mhandlers "testovoe/internal/measures/transport/http/handlers"
	pservice "testovoe/internal/products/service"
	ptransport "testovoe/internal/products/transport"
	phandlers "testovoe/internal/products/transport/http/handlers"
	"testovoe/internal/repository"
	"testovoe/internal/service"

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

func (a *App) initUniversalRoutes(repo *repository.UnifiedRepository) {
	if a.DB == nil {
		log.Fatal("DB connection is nil")
	}

	var _ interfaces.Repository = repo

	universalService := service.NewUniversalService(repo)

	a.Router.GET("/api/entities", func(c *gin.Context) {
		entityType := c.Query("type")
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

		res, err := universalService.GetAllEntities(
			interfaces.EntityType(entityType),
			c.Request.Context(),
			limit,
			offset,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	log.Println("Universal routes initialized successfully")
}

func (a *App) initializeRoutes() {
	if a.DB == nil {
		log.Fatal("DB connection is nil")
	}

	repo := repository.NewRepository(a.DB)
	a.initUniversalRoutes(repo) // только универсальные маршруты
	a.initProducts(repo)        // маршруты продуктов
	a.initMeasures(repo)        // маршруты мер

	log.Println("All routes initialized successfully")
}

func (a *App) initProducts(repo *repository.UnifiedRepository) {
	log.Println("Initializing products routes...")
	productService := pservice.NewService(repo.Product)
	productHandler := phandlers.NewHandler(productService)
	ptransport.NewRouter(productHandler).RegisterRoutes(a.Router)
}

func (a *App) initMeasures(repo *repository.UnifiedRepository) {
	log.Println("Initializing measures routes...")
	measureService := mservice.NewService(repo.Measure)
	measureHandler := mhandlers.NewHandler(measureService)
	mtransport.NewRouter(measureHandler).RegisterRoutes(a.Router)
}

func (a *App) Close() {
	if err := database.CloseDB(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}
