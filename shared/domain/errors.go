package domain

import (
	"fmt"
	"net/http"
)

// DomainError represents a standard error structure for the business domain.
// It contains enough information for layers above to handle it appropriately.
type DomainError struct {
	HTTPStatusCode int    // The HTTP status code that corresponds to this error.
	Code           string // An internal, stable error code (e.g., "ALREADY_EXISTS").
	Message        string // A user-friendly message.
	cause          error  // The original underlying error, for internal logging.
}

// Error makes DomainError satisfy the standard error interface.
func (e *DomainError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap provides compatibility with errors.Is and errors.As.
func (e *DomainError) Unwrap() error {
	return e.cause
}

// --- Error Constructors ---

// NewAlreadyExistsError creates a new domain error for a resource that already exists.
func NewAlreadyExistsError(message string, cause error) *DomainError {
	return &DomainError{
		HTTPStatusCode: http.StatusConflict, // 409
		Code:           "ALREADY_EXISTS",
		Message:        message,
		cause:          cause,
	}
}

// NewInvalidInputError creates a new domain error for invalid user input.
func NewInvalidInputError(message string, cause error) *DomainError {
	return &DomainError{
		HTTPStatusCode: http.StatusBadRequest, // 400
		Code:           "INVALID_INPUT",
		Message:        message,
		cause:          cause,
	}
}

// NewNotFoundError creates a new domain error for a resource that cannot be found.
func NewNotFoundError(message string, cause error) *DomainError {
	return &DomainError{
		HTTPStatusCode: http.StatusNotFound, // 404
		Code:           "NOT_FOUND",
		Message:        message,
		cause:          cause,
	}
}