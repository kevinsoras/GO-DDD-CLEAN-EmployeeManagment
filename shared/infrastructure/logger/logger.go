package logger

import (
	"log/slog"
	"os"
	"strings"
)

// New creates a new slog.Logger based on environment variables.
// It reads LOG_LEVEL and LOG_FORMAT to configure the logger.
func New() *slog.Logger {
	logLevel := os.Getenv("LOG_LEVEL")
	logFormat := os.Getenv("LOG_FORMAT")

	var level slog.Level
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo // Default level
	}

	var handler slog.Handler
	switch strings.ToUpper(logFormat) {
	case "JSON":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	case "TEXT":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}) // Default format
	}

	logger := slog.New(handler)
	slog.SetDefault(logger) // Set this logger as the default for the whole application

	logger.Info("Logger initialized", "level", level.String(), "format", logFormat)

	return logger
}
