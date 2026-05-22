package utils

import "strings"

func NewValidationError(message []string) *AppError {
	return &AppError{
		Code:    "VALIDATION_ERROR",
		Message: strings.Join(message, ", "),
		Status:  400,
	}
}
