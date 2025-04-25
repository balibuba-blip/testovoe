package interfaces

import (
	"context"
	"testovoe/internal/measures/models"
)

type MeasureRepository interface {
	GetByID(ctx context.Context, id int) (*models.Measure, error)
	GetAll(ctx context.Context) ([]models.Measure, error)
	Create(ctx context.Context, measure *models.Measure) (int, error)
	Update(ctx context.Context, id int, m *models.Measure) (*models.Measure, error)
	Delete(ctx context.Context, id int) error
}
