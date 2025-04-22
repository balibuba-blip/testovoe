package services

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"testovoe/internal/models"
	"testovoe/internal/repositories"
)

func CreateMeasure(measure *models.Measure) (int, error) {
	if measure.Name == "" {
		return 0, errors.New("measure name cannot be empty")
	}
	return repositories.CreateMeasure(measure)
}

func GetMeasureByID(id int) (*models.Measure, error) {
	if id <= 0 {
		return nil, errors.New("invalid measure ID")
	}
	return repositories.GetMeasureByID(id)
}

func GetAllMeasures() ([]models.Measure, error) {
	return repositories.GetAllMeasures()
}

func UpdateMeasure(id int, updateData *models.Measure) (*models.Measure, error) {
	if updateData.Name == "" {
		return nil, errors.New("measure name cannot be empty")
	}

	//проверка существует ли и обновление
	updatedMeasure, err := repositories.UpdateMeasure(id, updateData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("measure not found")
		}
		return nil, fmt.Errorf("failed to update measure: %w", err)
	}

	return updatedMeasure, nil
}

func DeleteMeasure(id int) error {
	//проверка существует ли measure
	if _, err := GetProductByID(strconv.Itoa(id)); err != nil {
		return fmt.Errorf("delete validation failed: %w", err)
	}

	_, err := repositories.DeleteMeasure(id)
	return err
}
