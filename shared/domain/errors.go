package domain

import (
	"fmt"
	"net/http" // Using net/http for status codes as an example
)

// ClientError defines an interface for errors that can be safely exposed to clients.
type ClientError interface {
	error
	ClientMessage() string
	ClientStatusCode() int
}

// DomainError represents an error that originates from the domain layer.
type DomainError struct {
	Msg        string
	Code       string // A machine-readable code for programmatic handling
	WrappedErr error  // The original error that was wrapped
}

// Error implements the error interface for DomainError.
func (e *DomainError) Error() string {
	if e.WrappedErr != nil {
		return fmt.Sprintf("%s (code: %s): %v", e.Msg, e.Code, e.WrappedErr)
	}
	return fmt.Sprintf("%s (code: %s)", e.Msg, e.Code)
}

// Unwrap returns the wrapped error, allowing errors.Is and errors.As to work.
func (e *DomainError) Unwrap() error {
	return e.WrappedErr
}

// ClientMessage returns a message suitable for client consumption.
func (e *DomainError) ClientMessage() string {
	// For domain errors, the message is often suitable for the client.
	return e.Msg
}

// ClientStatusCode maps domain error codes to HTTP status codes.
func (e *DomainError) ClientStatusCode() int {
	switch e.Code {
	case "NOT_FOUND":
		return http.StatusNotFound // 404
	case "ALREADY_EXISTS":
		return http.StatusConflict // 409
	case "INVALID_INPUT":
		return http.StatusBadRequest // 400
	case "UNAUTHORIZED":
		return http.StatusUnauthorized // 401
	case "FORBIDDEN": // Example of another domain error
		return http.StatusForbidden // 403
	default:
		return http.StatusInternalServerError // 500 for unhandled domain errors
	}
}

// Helper functions to create specific domain errors
func NewNotFoundError(msg string, err error) error {
	return &DomainError{Msg: msg, Code: "NOT_FOUND", WrappedErr: err}
}

func NewAlreadyExistsError(msg string, err error) error {
	return &DomainError{Msg: msg, Code: "ALREADY_EXISTS", WrappedErr: err}
}

func NewInvalidInputError(msg string, err error) error {
	return &DomainError{Msg: msg, Code: "INVALID_INPUT", WrappedErr: err}
}

func NewUnauthorizedError(msg string, err error) error {
	return &DomainError{Msg: msg, Code: "UNAUTHORIZED", WrappedErr: err}
}

func NewForbiddenError(msg string, err error) error {
	return &DomainError{Msg: msg, Code: "FORBIDDEN", WrappedErr: err}
}

// Generic internal server error for domain layer, if a domain rule leads to an unexpected internal issue
func NewInternalServerError(msg string, err error) error {
	return &DomainError{Msg: msg, Code: "INTERNAL_SERVER_ERROR", WrappedErr: err}
}

// You can still define specific error instances if they are common and don't need dynamic messages
var (
	ErrEmployeeNotFound = NewNotFoundError("employee not found", nil)
	ErrPersonNotFound   = NewNotFoundError("person not found", nil)
	ErrAlreadyExists    = NewAlreadyExistsError("resource already exists", nil)
)
