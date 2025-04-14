package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"product-api/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {

	var err error
	connStr := "user=postgres password=... dbname=products port=5432 sslmode=disable client_encoding=UTF8"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//проверка для определения к нужной ли БД подключается
	//fmt.Println("Подключение к БД:", connStr)
	//log.Println("Количество продуктов в БД:", getProductCount(db))
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}

//для проверки
//func getProductCount(db *sql.DB) int {
//	var count int
//	err := db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return count
//}

func main() {
	//Инициализация БД
	initDB()
	defer db.Close()

	//маршрутизатор
	router := gin.Default()

	// Продукты
	router.GET("/product", getProducts)
	router.GET("/product/:id", getProduct)
	router.POST("/product", createProduct)
	router.PUT("/product/:id", updateProduct)
	router.DELETE("/product/:id", deleteProduct)

	// Единицы измерения
	router.GET("/measure", getMeasures)
	router.GET("/measure/:id", getMeasure)
	router.POST("/measure", createMeasure)
	router.PUT("/measure/:id", updateMeasure)
	router.DELETE("/measure/:id", deleteMeasure)

	//Запуск сервера
	router.Run("localhost:8080")
}

// Получить все продукты
func getProducts(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	rows, err := db.Query(`
    SELECT 
        id, 
        convert_from(name::bytea, 'UTF8') as name,
        quantity, 
        unit_cost, 
        measure_id 
    FROM products
`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.UnitCost, &p.MeasureID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, p)
	}

	c.JSON(http.StatusOK, products)
}

// Получить продукт по ID
func getProduct(c *gin.Context) {
	id := c.Param("id")

	var p models.Product
	err := db.QueryRow("SELECT id, name, quantity, unit_cost, measure_id FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Quantity, &p.UnitCost, &p.MeasureID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

// Создать продукт
func createProduct(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int
	err := db.QueryRow(
		"INSERT INTO products (name, quantity, unit_cost, measure_id) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Name, p.Quantity, p.UnitCost, p.MeasureID,
	).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// Обновить продукт
func updateProduct(c *gin.Context) {
	id := c.Param("id")

	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec(
		"UPDATE products SET name = $1, quantity = $2, unit_cost = $3, measure_id = $4 WHERE id = $5",
		p.Name, p.Quantity, p.UnitCost, p.MeasureID, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Удалить продукт
func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	result, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Получить все единицы измерения
func getMeasures(c *gin.Context) {
	rows, err := db.Query("SELECT id, name FROM measures")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var measures []models.Measure
	for rows.Next() {
		var m models.Measure
		if err := rows.Scan(&m.ID, &m.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		measures = append(measures, m)
	}

	c.JSON(http.StatusOK, measures)
}

// Получить единицу измерения по ID
func getMeasure(c *gin.Context) {
	id := c.Param("id")

	var m models.Measure
	err := db.QueryRow("SELECT id, name FROM measures WHERE id = $1", id).
		Scan(&m.ID, &m.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Measure not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, m)
}

// Создать единицу измерения
func createMeasure(c *gin.Context) {
	var m models.Measure
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int
	err := db.QueryRow(
		"INSERT INTO measures (name) VALUES ($1) RETURNING id",
		m.Name,
	).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// Обновить единицу измерения
func updateMeasure(c *gin.Context) {
	id := c.Param("id")

	var m models.Measure
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec(
		"UPDATE measures SET name = $1 WHERE id = $2",
		m.Name, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Measure not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Удалить единицу измерения
func deleteMeasure(c *gin.Context) {
	id := c.Param("id")

	result, err := db.Exec("DELETE FROM measures WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Measure not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
