package repository

const (
	getAllProductsQuery = `
        SELECT id, name, quantity, unit_cost, measure_id 
        FROM products
        ORDER BY id
        LIMIT $1 OFFSET $2`

	createProductQuery = `
        INSERT INTO products (name, quantity, unit_cost, measure_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	updateProductQuery = `
        UPDATE products
        SET name = $1, quantity = $2, unit_cost = $3, measure_id = $4
        WHERE id = $5
        RETURNING id, name, quantity, unit_cost, measure_id`

	deleteProductQuery = `
        DELETE FROM products
        WHERE id = $1
        RETURNING id`
)
