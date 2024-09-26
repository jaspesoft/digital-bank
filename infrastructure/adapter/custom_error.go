package adapter

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationErrors(err error) []ErrorResponse {
	var validationErrors []ErrorResponse

	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		for _, fieldErr := range validationErrs {
			validationErrors = append(validationErrors, ErrorResponse{
				Field:   fieldErr.StructField(),
				Message: fmt.Sprintf("%s", fieldErr.Tag()),
			})
		}
	}

	return validationErrors
}
