package utils

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/kevinsoras/employee-management/shared/domain"
	"github.com/kevinsoras/employee-management/shared/infrastructure"
)

// HandleHTTPError inspects an error and writes an appropriate HTTP response.
// It centralizes the logic for handling different error types.
func HandleHTTPError(w http.ResponseWriter, logger *slog.Logger, err error) {
	var domainErr *domain.DomainError
	var infraErr *infrastructure.InfrastructureError

	// 1. Check for a rich DomainError first.
	if errors.As(err, &domainErr) {
		logger.Warn("Domain error occurred", "code", domainErr.Code, "msg", domainErr.Message, "original_err", domainErr.Unwrap())
		WriteJSONError(w, domainErr.Message, domainErr.HTTPStatusCode)
		return
	}

	// 2. Check for an Application-level validation error.
	if errors.Is(err, ErrValidation) {
		logger.Info("Request validation failed", "error", err)
		WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Check for common Infrastructure errors.
	if errors.As(err, &infraErr) {
				logger.Error("Infrastructure error occurred", "code", infraErr.Code, "msg", infraErr.Error(), "original_err", infraErr.Unwrap())
		WriteJSONError(w, infraErr.Error(), http.StatusInternalServerError)

		return
	}

	// 4. Check for specific system errors like context timeout.
	if errors.Is(err, context.DeadlineExceeded) {
		logger.Error("Request timed out", "error", err)
		WriteJSONError(w, "La solicitud excedió el tiempo de espera.", http.StatusGatewayTimeout)
		return
	}

	// 5. Fallback for any other unexpected error.
	logger.Error("Unexpected system error", "error", err)
	WriteJSONError(w, "Ocurrió un error interno inesperado", http.StatusInternalServerError)
}