package main

import (
	"fmt"
	"log"
	"test_task/internal/config"
	"test_task/internal/repositories"
	"test_task/internal/routes"
)

func main() {
	//загрузка конфигурации
	cfg := config.Load()
	//log.Printf("DB config: %+v", cfg)

	//строчка подключения
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	// Инициализация БД
	if err := repositories.InitDB(connStr); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Настройка роутера
	router := routes.SetupRouter()

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

//package main
//
//import (
//	"database/sql"
//	"fmt"
//	"log"
//	"net/http"
//	"os"
//	"product-api/models"
//
//	"github.com/gin-gonic/gin"
//	"github.com/joho/godotenv"
//	_ "github.com/lib/pq"
//)
//
//var db *sql.DB
//
//func loadConfig() (*Config, error) {
//	cfg := &Config{
//		DB_USER:     os.Getenv("DB_USER"),
//		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
//		DB_NAME:     os.Getenv("DB_NAME"),
//		DB_PORT:     os.Getenv("DB_PORT"),
//		DB_SSL_MODE: os.Getenv("DB_SSL_MODE"),
//	}
//
//	if cfg.DB_USER == "" || cfg.DB_PASSWORD == "" || cfg.DB_NAME == "" {
//		return nil, fmt.Errorf("не заданы обязательные параметры БД")
//	}
//
//	return cfg, nil
//}
//
//func initDB() (*sql.DB, error) {
//
//	//загрузка конфигурации
//	cfg, err := loadConfig()
//	if err != nil {
//		return nil, fmt.Errorf("Ошибка загрузки конфигурации: %w", err)
//	}
//
//	//строка подключения
//	connStr := fmt.Sprintf(
//		"user=%s password=%s dbname=%s port=%s sslmode=%s",
//		cfg.DB_USER,
//		cfg.DB_PASSWORD,
//		cfg.DB_NAME,
//		cfg.DB_PORT,
//		cfg.DB_SSL_MODE,
//	)
//
//	//подключение к БД
//	db, err = sql.Open("postgres", connStr)
//	if err != nil {
//		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
//	}
//
//	// проверка соединения
//	if err = db.Ping(); err != nil {
//		db.Close() // Закрываем соединение при ошибке
//		return nil, fmt.Errorf("ошибка проверки соединения: %w", err)
//	}
//	log.Printf("Успешное подключение к БД %s на порту %s", cfg.DB_NAME, cfg.DB_PORT)
//	return db, nil
//}
//
//func main() {
//	// загрузка .env
//	_ = godotenv.Load() // Игнорируем ошибку, если файла нет
//
//	//инициализация БД
//	db, err := initDB()
//    if err != nil {
//        log.Fatalf("Ошибка инициализации БД: %v", err)
//    }
//    defer db.Close()
//
//	//маршрутизатор
//	router := gin.Default()
//
//	// Продукты
//	router.GET("/product", getProducts)
//	router.GET("/product/:id", getProduct)
//	router.POST("/product", createProduct)
//	router.PUT("/product/:id", updateProduct)
//	router.DELETE("/product/:id", deleteProduct)
//
//	// Единицы измерения
//	router.GET("/measure", getMeasures)
//	router.GET("/measure/:id", getMeasure)
//	router.POST("/measure", createMeasure)
//	router.PUT("/measure/:id", updateMeasure)
//	router.DELETE("/measure/:id", deleteMeasure)
//
//	//Запуск сервера
//	router.Run("localhost:8080")
//}
//
//// Получить все продукты
//func getProducts(c *gin.Context) {
//	c.Header("Content-Type", "application/json; charset=utf-8")
//	rows, err := db.Query(`
//    SELECT
//        id,
//        convert_from(name::bytea, 'UTF8') as name,
//        quantity,
//        unit_cost,
//        measure_id
//    FROM products
//`)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	defer rows.Close()
//
//	var products []models.Product
//	for rows.Next() {
//		var p models.Product
//		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.UnitCost, &p.MeasureID); err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		products = append(products, p)
//	}
//
//	c.JSON(http.StatusOK, products)
//}
//
//// Получить продукт по ID
//func getProduct(c *gin.Context) {
//	id := c.Param("id")
//
//	var p models.Product
//	err := db.QueryRow("SELECT id, name, quantity, unit_cost, measure_id FROM products WHERE id = $1", id).
//		Scan(&p.ID, &p.Name, &p.Quantity, &p.UnitCost, &p.MeasureID)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
//			return
//		}
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, p)
//}
//
//// Создать продукт
//func createProduct(c *gin.Context) {
//	var p models.Product
//	if err := c.ShouldBindJSON(&p); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	var id int
//	err := db.QueryRow(
//		"INSERT INTO products (name, quantity, unit_cost, measure_id) VALUES ($1, $2, $3, $4) RETURNING id",
//		p.Name, p.Quantity, p.UnitCost, p.MeasureID,
//	).Scan(&id)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{"id": id})
//}
//
//// Обновить продукт
//func updateProduct(c *gin.Context) {
//	id := c.Param("id")
//
//	var p models.Product
//	if err := c.ShouldBindJSON(&p); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	result, err := db.Exec(
//		"UPDATE products SET name = $1, quantity = $2, unit_cost = $3, measure_id = $4 WHERE id = $5",
//		p.Name, p.Quantity, p.UnitCost, p.MeasureID, id,
//	)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	rowsAffected, _ := result.RowsAffected()
//	if rowsAffected == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"status": "success"})
//}
//
//// Удалить продукт
//func deleteProduct(c *gin.Context) {
//	id := c.Param("id")
//
//	result, err := db.Exec("DELETE FROM products WHERE id = $1", id)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	rowsAffected, _ := result.RowsAffected()
//	if rowsAffected == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"status": "success"})
//}
//
//// Получить все единицы измерения
//func getMeasures(c *gin.Context) {
//	rows, err := db.Query("SELECT id, name FROM measures")
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	defer rows.Close()
//
//	var measures []models.Measure
//	for rows.Next() {
//		var m models.Measure
//		if err := rows.Scan(&m.ID, &m.Name); err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		measures = append(measures, m)
//	}
//
//	c.JSON(http.StatusOK, measures)
//}
//
//// Получить единицу измерения по ID
//func getMeasure(c *gin.Context) {
//	id := c.Param("id")
//
//	var m models.Measure
//	err := db.QueryRow("SELECT id, name FROM measures WHERE id = $1", id).
//		Scan(&m.ID, &m.Name)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			c.JSON(http.StatusNotFound, gin.H{"error": "Measure not found"})
//			return
//		}
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, m)
//}
//
//// Создать единицу измерения
//func createMeasure(c *gin.Context) {
//	var m models.Measure
//	if err := c.ShouldBindJSON(&m); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	var id int
//	err := db.QueryRow(
//		"INSERT INTO measures (name) VALUES ($1) RETURNING id",
//		m.Name,
//	).Scan(&id)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{"id": id})
//}
//
//// Обновить единицу измерения
//func updateMeasure(c *gin.Context) {
//	id := c.Param("id")
//
//	var m models.Measure
//	if err := c.ShouldBindJSON(&m); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	result, err := db.Exec(
//		"UPDATE measures SET name = $1 WHERE id = $2",
//		m.Name, id,
//	)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	rowsAffected, _ := result.RowsAffected()
//	if rowsAffected == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Measure not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"status": "success"})
//}
//
//// Удалить единицу измерения
//func deleteMeasure(c *gin.Context) {
//	id := c.Param("id")
//
//	result, err := db.Exec("DELETE FROM measures WHERE id = $1", id)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	rowsAffected, _ := result.RowsAffected()
//	if rowsAffected == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Measure not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"status": "success"})
//}
//
