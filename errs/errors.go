package errs

import "net/http"

// AppError is
type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

// AsMessage is
func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

// NewNotFoundError is
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

// NewUnexpectedError is
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
