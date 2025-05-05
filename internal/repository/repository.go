package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	measuresmodels "testovoe/internal/measures/models"
	productsmodels "testovoe/internal/products/models"
)

type UnifiedRepository struct {
	Product *ProductRepository
	Measure *MeasureRepository
}

func NewRepository(db *sql.DB) *UnifiedRepository {
	return &UnifiedRepository{
		Product: &ProductRepository{db: db},
		Measure: &MeasureRepository{db: db},
	}
}

// ProductRepository
type ProductRepository struct {
	db *sql.DB
}

func (r *ProductRepository) GetByID(id int) (*productsmodels.Product, error) {
	var product productsmodels.Product
	err := r.db.QueryRow(
		`SELECT id, name, quantity, unit_cost, measure_id FROM products WHERE id = $1`,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Quantity,
		&product.UnitCost,
		&product.MeasureID,
	)
	return &product, err
}

func (r *ProductRepository) GetAll(limit, offset int) ([]productsmodels.Product, error) {
	rows, err := r.db.Query(getAllProductsQuery, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []productsmodels.Product
	for rows.Next() {
		var p productsmodels.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.UnitCost, &p.MeasureID); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func (r *ProductRepository) Create(p *productsmodels.Product) (int, error) {
	var id int
	err := r.db.QueryRow(
		createProductQuery,
		p.Name,
		p.Quantity,
		p.UnitCost,
		p.MeasureID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create product: %w", err)
	}
	return id, nil
}

func (r *ProductRepository) Update(id int, p *productsmodels.Product) (*productsmodels.Product, error) {
	var updated productsmodels.Product
	err := r.db.QueryRow(
		updateProductQuery,
		p.Name,
		p.Quantity,
		p.UnitCost,
		p.MeasureID,
		id,
	).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Quantity,
		&updated.UnitCost,
		&updated.MeasureID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found with id: %d", id)
		}
		return nil, fmt.Errorf("failed to update product: %w", err)
	}
	return &updated, nil
}

func (r *ProductRepository) Delete(id int) error {
	var deletedID int
	err := r.db.QueryRow(deleteProductQuery, id).Scan(&deletedID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("product with id %d not found", id)
		}
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}

type MeasureRepository struct {
	db *sql.DB
}

// Measures methods
func (r *MeasureRepository) GetByID(ctx context.Context, id int) (*measuresmodels.Measure, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}
	var m measuresmodels.Measure
	err := r.db.QueryRowContext(ctx, getMeasureByIDQuery, id).
		Scan(&m.ID, &m.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("measure not found: %w", err)
		}
		return nil, fmt.Errorf("repository error: %w", err)
	}

	return &m, nil
}

func (r *MeasureRepository) GetAll(ctx context.Context, limit, offset int) ([]measuresmodels.Measure, error) {
	rows, err := r.db.QueryContext(ctx, getAllMeasuresQuery, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query measures: %w", err)
	}
	defer rows.Close()

	var measures []measuresmodels.Measure
	for rows.Next() {
		var m measuresmodels.Measure
		if err := rows.Scan(&m.ID, &m.Name); err != nil {
			return nil, fmt.Errorf("failed to scan measure: %w", err)
		}
		measures = append(measures, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return measures, nil
}

func (r *MeasureRepository) Create(ctx context.Context, m *measuresmodels.Measure) (int, error) {
	var id int
	err := r.db.QueryRowContext(ctx, createMeasureQuery, m.Name).
		Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to create measure: %w", err)
	}

	return id, nil
}

func (r *MeasureRepository) Update(ctx context.Context, id int, m *measuresmodels.Measure) (*measuresmodels.Measure, error) {
	var updated measuresmodels.Measure
	err := r.db.QueryRowContext(ctx, updateMeasureQuery, m.Name, id).
		Scan(&updated.ID, &updated.Name)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("measure not found: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update measure: %w", err)
	}

	return &updated, nil
}

func (r *MeasureRepository) Delete(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(ctx, deleteMeasureQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete measure: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("measure not found")
	}

	return nil
}
