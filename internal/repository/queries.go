package repository

const (
	//products
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

	//Measures
	getAllMeasuresQuery = `
		SELECT id, name 
    	FROM measures
    	ORDER BY id
    	LIMIT $1 OFFSET $2`

	getMeasureByIDQuery = `
		SELECT id, name 
		FROM measures 
		WHERE id = $1`

	createMeasureQuery = `
		INSERT INTO measures (name)
		VALUES ($1)
		RETURNING id`

	updateMeasureQuery = `
		UPDATE measures
		SET name = $1
		WHERE id = $2
		RETURNING id, name`

	deleteMeasureQuery = `
		DELETE FROM measures
		WHERE id = $1`
)
