package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"test_task/internal/models"
)

const (
	createMeasureQuery  = `INSERT INTO measures (name) VALUES ($1) RETURNING id`
	getMeasureQuery     = `SELECT id, name FROM measures WHERE id = $1`
	getAllMeasuresQuery = `SELECT id, name FROM measures`
	updateMeasureQuery  = `UPDATE measures SET name = $1 WHERE id = $2 RETURNING id`
	deleteMeasureQuery  = `DELETE FROM measures WHERE id = $1 RETURNING id`
)

func CreateMeasure(measure *models.Measure) (int, error) {
	var id int
	err := db.QueryRow(createMeasureQuery, measure.Name).Scan(&id)
	return id, fmt.Errorf("failed to create measure: %w", err)
}

func GetMeasureByID(id int) (*models.Measure, error) {
	var m models.Measure
	err := db.QueryRow(getMeasureQuery, id).Scan(&m.ID, &m.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get measure: %w (id: %d)", err, id)
	}
	return &m, nil
}

func GetAllMeasures() ([]models.Measure, error) {
	rows, err := db.Query(getAllMeasuresQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var measures []models.Measure
	for rows.Next() {
		var m models.Measure
		if err := rows.Scan(
			&m.ID,
			&m.Name,
		); err != nil {
			log.Printf("Executing query: %s", getAllMeasuresQuery)
			return nil, err
		}
		measures = append(measures, m)
	}

	return measures, nil
}

func UpdateMeasure(id int, m *models.Measure) (*models.Measure, error) {
	var updatedMeasure models.Measure
	err := db.QueryRow(
		updateMeasureQuery,
		m.Name,
		id,
	).Scan(
		&updatedMeasure.ID,
		&updatedMeasure.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update measure: %w (id: %d)", err, id)
	}
	return &updatedMeasure, nil
}

func DeleteMeasure(id int) (int, error) {
	var deletedID int
	err := db.QueryRow(deleteMeasureQuery, id).Scan(&deletedID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("measure not found: %w", err)
		}
		return 0, fmt.Errorf("failed to delete measure: %w (id: %d)", err, id)
	}
	return deletedID, nil
}
