package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"test_task/internal/models"
	"test_task/internal/services"

	"github.com/gin-gonic/gin"
)

func GetAllMeasures(c *gin.Context) {
	measures, err := services.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, measures)
}

func GetMeasure(c *gin.Context) {
	id := c.Param("id")

	product, err := services.GetMeasureByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateMeasure(c *gin.Context) {
	var measure models.Measure
	if err := c.ShouldBindJSON(&measure); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.CreateMeasure(&measure)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateMeasure(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid measure ID"})
		return
	}

	var updateData models.Measure
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMeasure, err := services.UpdateMeasure(id, &updateData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "measure not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedMeasure)
}

func DeleteMeasure(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid measure ID"})
		return
	}

	if err := services.DeleteMeasure(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "measure not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "measure deleted successfully",
		"id":      id,
	})
}

//!!!Допилить GetMeasure, UpdateMeasure, DeleteMeasure
// Аналогично для GetAllMeasures, UpdateMeasure, DeleteMeasure
