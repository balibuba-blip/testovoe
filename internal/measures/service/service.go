package service

import (
	"context"
	"errors"
	"fmt"
	"testovoe/internal/measures/models"
	"testovoe/internal/repository"
)

type MeasureService struct {
	repo *repository.MeasureRepository
}

func NewService(repo *repository.MeasureRepository) *MeasureService {
	return &MeasureService{repo: repo}
}

func (s *MeasureService) GetByID(ctx context.Context, id int) (*models.Measure, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MeasureService) GetAll(ctx context.Context, limit, offset int) ([]models.Measure, error) {
	if limit <= 0 || offset < 0 {
		return nil, errors.New("invalid pagination parameters")
	}
	return s.repo.GetAll(ctx, limit, offset) // Передаем параметры
}

func (s *MeasureService) Create(ctx context.Context, m *models.Measure) (int, error) {
	if m.Name == "" {
		return 0, errors.New("measure name cannot be empty")
	}
	return s.repo.Create(ctx, m)
}

func (s *MeasureService) Update(ctx context.Context, id int, measure *models.Measure) (*models.Measure, error) {
	if measure.Name == "" {
		return nil, fmt.Errorf("measure name cannot be empty")
	}

	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return nil, fmt.Errorf("measure not found: %w", err)
	}

	updatedMeasure, err := s.repo.Update(ctx, id, measure)
	if err != nil {
		return nil, fmt.Errorf("update failed: %w", err)
	}

	return updatedMeasure, nil
}

func (s *MeasureService) Delete(ctx context.Context, id int) error {
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		return fmt.Errorf("measure not found: %w", err)
	}
	return s.repo.Delete(ctx, id)
}
