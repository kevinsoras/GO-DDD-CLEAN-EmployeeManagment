package infrastructure

import (
	"fmt"
)

// InfrastructureError represents an error that originates from the infrastructure layer.
type InfrastructureError struct {
	Msg        string
	Code       string // A machine-readable code for programmatic handling
	WrappedErr error  // The original error that was wrapped
}

// Error implements the error interface for InfrastructureError.
func (e *InfrastructureError) Error() string {
	if e.WrappedErr != nil {
		return fmt.Sprintf("%s (code: %s): %v", e.Msg, e.Code, e.WrappedErr)
	}
	return fmt.Sprintf("%s (code: %s)", e.Msg, e.Code)
}

// Unwrap returns the wrapped error, allowing errors.Is and errors.As to work.
func (e *InfrastructureError) Unwrap() error {
	return e.WrappedErr
}

// Helper functions to create specific infrastructure errors
func NewDBError(msg string, err error) error {
	return &InfrastructureError{Msg: msg, Code: "DB_ERROR", WrappedErr: err}
}

func NewNetworkError(msg string, err error) error {
	return &InfrastructureError{Msg: msg, Code: "NETWORK_ERROR", WrappedErr: err}
}

func NewExternalServiceError(msg string, err error) error {
	return &InfrastructureError{Msg: msg, Code: "EXTERNAL_SERVICE_ERROR", WrappedErr: err}
}

// You can define specific error instances if they are common
var (
	ErrDBConnectionFailed = NewDBError("database connection failed", nil)
	ErrRecordNotFound     = NewDBError("record not found in database", nil) // Specific DB error for not found
	ErrUniqueConstraint   = NewDBError("unique constraint violation", nil)  // Specific DB error for unique constraint
)
