package service

import (
	"errors"
	"fmt"
	"testovoe/internal/interfaces"
	"testovoe/internal/products/models"
)

type Service struct {
	repo interfaces.ProductRepository
}

func NewService(repo interfaces.ProductRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *Service) Create(p *models.Product) (int, error) {
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

func (s *Service) Update(id int, p *models.Product) (*models.Product, error) {
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

func (s *Service) Delete(id int) error {
	// Проверка существования продукта
	if _, err := s.repo.GetByID(id); err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	return s.repo.Delete(id)
}
