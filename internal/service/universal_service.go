package service

import (
	"context"
	"testovoe/internal/interfaces"
)

type UniversalService struct {
	repo interfaces.Repository
}

func NewUniversalService(repo interfaces.Repository) *UniversalService {
	return &UniversalService{repo: repo}
}

func (s *UniversalService) GetAllEntities(entityType interfaces.EntityType, ctx context.Context, limit, offset int) (interface{}, error) {
	if entityType == "GetAllEntity" {
		// Возвращаем комбинированный результат
		products, err := s.repo.GetAllEntities("product", ctx, limit, offset)
		if err != nil {
			return nil, err
		}
		measures, err := s.repo.GetAllEntities("measure", ctx, limit, offset)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"products": products,
			"measures": measures,
		}, nil
	}
	return s.repo.GetAllEntities(entityType, ctx, limit, offset)
}
