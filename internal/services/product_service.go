package services

import (
	"errors"
	"fmt"
	"strconv"
	"test_task/internal/models"
	"test_task/internal/repositories"
)

// возвращает 1 продукт по id
func GetProductByID(id string) (*models.Product, error) {
	if err := validateProductID(id); err != nil {
		return nil, err
	}
	return repositories.GetProductByID(id)
}

// возвращает все продукты
func GetAllProducts() ([]models.Product, error) {
	return repositories.GetAllProducts()
}

// id созданной сущности, ошибка
func CreateProduct(product *models.Product) (int, error) {
	if product.Name == "" {
		return 0, errors.New("product name cannot be empty")
	}
	if product.Quantity < 0 {
		return 0, errors.New("quantity cannot be negative")
	}

	// Создание продукта
	return repositories.CreateProduct(product)
}

func UpdateProduct(id int, updateData *models.Product) (*models.Product, error) {
	if updateData.Name == "" {
		return nil, errors.New("product name cannot be empty")
	}
	if updateData.Quantity < 0 {
		return nil, errors.New("quantity cannot be negative")
	}

	//проверка существует ли продукт
	if _, err := GetProductByID(strconv.Itoa(id)); err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	return repositories.UpdateProduct(id, updateData)
}

func DeleteProduct(id int) error {
	//проверка существует ли продукт
	if _, err := GetProductByID(strconv.Itoa(id)); err != nil {
		return fmt.Errorf("delete validation failed: %w", err)
	}

	_, err := repositories.DeleteProduct(id)
	return err
}
