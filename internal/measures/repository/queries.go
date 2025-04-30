package repository

const (
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
