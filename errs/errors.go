package errs

import "net/http"

// Error is
type Error struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

// AsMessage is
func (e Error) AsMessage() *Error {
	return &Error{
		Message: e.Message,
	}
}

// NewNotFoundError is
func NewNotFoundError(message string) *Error {
	return &Error{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

// NewUnexpectedError is
func NewUnexpectedError(message string) *Error {
	return &Error{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
