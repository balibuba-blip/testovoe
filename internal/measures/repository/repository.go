package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testovoe/internal/measures/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context, limit, offset int) ([]models.Measure, error) {

	rows, err := r.db.QueryContext(ctx, getAllMeasuresQuery, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query measures: %w", err)
	}
	defer rows.Close()

	var measures []models.Measure
	for rows.Next() {
		var m models.Measure
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

func (r *Repository) GetByID(ctx context.Context, id int) (*models.Measure, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}
	var m models.Measure
	err := r.db.QueryRowContext(ctx, getMeasureByIDQuery, id).
		Scan(&m.ID, &m.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("measure not found: %w", err)
		}
		return nil, fmt.Errorf("repository error: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get measure: %w", err)
	}

	return &m, nil
}

func (r *Repository) Create(ctx context.Context, m *models.Measure) (int, error) {
	var id int
	err := r.db.QueryRowContext(ctx, createMeasureQuery, m.Name).
		Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to create measure: %w", err)
	}

	return id, nil
}

func (r *Repository) Update(ctx context.Context, id int, m *models.Measure) (*models.Measure, error) {
	var updated models.Measure
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

func (r *Repository) Delete(ctx context.Context, id int) error {
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
