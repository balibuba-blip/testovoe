package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testovoe/internal/models"
)

const (
	getAllProductsQuery = `SELECT id, name, quantity, unit_cost, measure_id FROM products`
	getProductByIDQuery = `SELECT id, name, quantity, unit_cost, measure_id FROM products WHERE id = $1`
	createProductQuery  = `INSERT INTO products (name, quantity, unit_cost, measure_id) VALUES ($1, $2, $3, $4) RETURNING id`
	updateProductQuery  = `
	UPDATE products 
	SET name = $1, quantity = $2, unit_cost = $3, measure_id = $4 
	WHERE id = $5
	RETURNING id, name, quantity, unit_cost, measure_id`
	deleteProductQuery = `DELETE FROM products WHERE id = $1 RETURNING id`
)

func GetAllProducts() ([]models.Product, error) {

	rows, err := db.Query(getAllProductsQuery)
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
			log.Printf("Executing query: %s", getAllProductsQuery)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
func GetProductByID(id string) (*models.Product, error) {

	var product models.Product
	err := db.QueryRow(getProductByIDQuery, id).Scan(
		&product.ID,
		&product.Name,
		&product.Quantity,
		&product.UnitCost,
		&product.MeasureID,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func CreateProduct(p *models.Product) (int, error) {
	var id int
	err := db.QueryRow(
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

func UpdateProduct(id int, p *models.Product) (*models.Product, error) {
	var updatedProduct models.Product
	err := db.QueryRow(
		updateProductQuery,
		p.Name,
		p.Quantity,
		p.UnitCost,
		p.MeasureID,
		id,
	).Scan(
		&updatedProduct.ID,
		&updatedProduct.Name,
		&updatedProduct.Quantity,
		&updatedProduct.UnitCost,
		&updatedProduct.MeasureID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w (id: %d)", err, id)
	}
	return &updatedProduct, nil
}

func DeleteProduct(id int) (int, error) {
	var deletedID int
	err := db.QueryRow(deleteProductQuery, id).Scan(&deletedID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("product not found: %w", err)
		}
		return 0, fmt.Errorf("failed to delete product: %w (id: %d)", err, id)
	}
	return deletedID, nil
}
