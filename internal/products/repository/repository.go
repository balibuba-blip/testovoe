package repository

import (
	"database/sql"
	"fmt"
	"log"
	"testovoe/internal/products/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByID(id int) (*models.Product, error) {
	var product models.Product
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

func (r *Repository) GetAll() ([]models.Product, error) {
	rows, err := r.db.Query(getAllProductsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Quantity,
			&p.UnitCost,
			&p.MeasureID,
		); err != nil {
			log.Printf("Error scanning product row: %v", err)
			continue
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *Repository) Create(p *models.Product) (int, error) {
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

func (r *Repository) Update(id int, p *models.Product) (*models.Product, error) {
	var updated models.Product
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

func (r *Repository) Delete(id int) error {
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
