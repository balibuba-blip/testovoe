package interfaces

import "testovoe/internal/products/models"

type ProductRepository interface {
	GetAll(limit, offset int) ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Create(product *models.Product) (int, error)
	Update(id int, product *models.Product) (*models.Product, error)
	Delete(id int) error
}
