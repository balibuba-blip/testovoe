package services

import (
	"errors"
	"fmt"
	"strconv"
	"test_task/internal/models"
	"test_task/internal/repositories"
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

	//проверка существует ли measure
	if _, err := GetMeasureByID(strconv.Itoa(id)); err != nil {
		return nil, fmt.Errorf("measure not found: %w", err)
	}

	return repositories.UpdateMeasure(id, updateData)
}

func DeleteMeasure(id int) error {
	//проверка существует ли measure
	if _, err := GetProductByID(strconv.Itoa(id)); err != nil {
		return fmt.Errorf("delete validation failed: %w", err)
	}

	_, err := repositories.DeleteMeasure(id)
	return err
}
