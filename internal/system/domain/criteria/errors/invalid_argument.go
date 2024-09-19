package errors

import "errors"

type InvalidArgumentError struct {
	error
}

func NewInvalidArgumentError(message string) *InvalidArgumentError {
	return &InvalidArgumentError{
		error: errors.New(message),
	}
}
