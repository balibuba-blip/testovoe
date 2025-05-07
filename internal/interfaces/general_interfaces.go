package interfaces

import (
	"context"
	measuresmodels "testovoe/internal/measures/models"
	productsmodels "testovoe/internal/products/models"
)

type EntityType string

const (
	EntityTypeProduct EntityType = "product"
	EntityTypeMeasure EntityType = "measure"
	EntityTypeAll     EntityType = "all"
)

type Repository interface {
	// Универсальный метод
	GetAllEntities(entityType EntityType, ctx context.Context, limit, offset int) (interface{}, error)

	// Measure methods
	GetMeasureByID(ctx context.Context, id int) (*measuresmodels.Measure, error)
	GetAllMeasures(ctx context.Context, limit, offset int) ([]measuresmodels.Measure, error)
	CreateMeasure(ctx context.Context, measure *measuresmodels.Measure) (int, error)
	UpdateMeasure(ctx context.Context, id int, m *measuresmodels.Measure) (*measuresmodels.Measure, error)
	DeleteMeasure(ctx context.Context, id int) error

	// Product methods
	GetAllProducts(limit, offset int) ([]productsmodels.Product, error)
	GetProductByID(id int) (*productsmodels.Product, error)
	CreateProduct(product *productsmodels.Product) (int, error)
	UpdateProduct(id int, product *productsmodels.Product) (*productsmodels.Product, error)
	DeleteProduct(id int) error
}
