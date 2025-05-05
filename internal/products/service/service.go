package service

import (
	"errors"
	"fmt"
	"testovoe/internal/products/models"
	"testovoe/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) GetAll(limit, offset int) ([]models.Product, error) {
	if limit <= 0 || offset < 0 {
		return nil, errors.New("invalid pagination parameters")
	}
	return s.repo.GetAll(limit, offset)
}

func (s *ProductService) Create(p *models.Product) (int, error) {
	if p.MeasureID <= 0 {
		return 0, errors.New("invalid measure ID")
	}
	if p.Name == "" {
		return 0, errors.New("product name cannot be empty")
	}
	if p.Quantity < 0 {
		return 0, errors.New("quantity cannot be negative")
	}
	if p.UnitCost <= 0 {
		return 0, errors.New("unit cost must be positive")
	}
	return s.repo.Create(p)
}

func (s *ProductService) Update(id int, p *models.Product) (*models.Product, error) {
	if p.Name == "" {
		return nil, errors.New("product name cannot be empty")
	}
	if p.Quantity < 0 {
		return nil, errors.New("quantity cannot be negative")
	}
	if p.UnitCost <= 0 {
		return nil, errors.New("unit cost must be positive")
	}

	// Проверка существования продукта
	if _, err := s.repo.GetByID(id); err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	return s.repo.Update(id, p)
}

func (s *ProductService) Delete(id int) error {
	// Проверка существования продукта
	if _, err := s.repo.GetByID(id); err != nil {
		return fmt.Errorf("product not found: %w", err)
	}
	return s.repo.Delete(id)
}
