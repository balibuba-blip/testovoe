package interfaces

import "testovoe/internal/products/models"

type ProductRepository interface {
	GetByID(id int) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Create(product *models.Product) (int, error)
	Update(id int, product *models.Product) (*models.Product, error)
	Delete(id int) error
}
