package services

import (
	"errors"
	"regexp"
)

var (
	ErrEmptyProductID   = errors.New("product ID cannot be empty")
	ErrInvalidProductID = errors.New("invalid product ID format")
)

func validateProductID(id string) error {
	if id == "" {
		return errors.New("ID cannot be empty")
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9-]+$`).MatchString(id) {
		return ErrInvalidProductID
	}

	return nil
}
