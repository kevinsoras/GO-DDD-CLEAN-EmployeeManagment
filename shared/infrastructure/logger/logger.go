package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// New creates a new slog.Logger based on environment variables.
func New() *slog.Logger {
	logLevel := os.Getenv("LOG_LEVEL")
	logFormat := os.Getenv("LOG_FORMAT")
	logOutputs := os.Getenv("LOG_OUTPUTS")
	logFilePath := os.Getenv("LOG_FILE_PATH")

	// --- Setup Log Level ---
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
		level = slog.LevelInfo
	}

	// --- Setup Log Outputs ---
	writers := []io.Writer{}
	outputSlice := strings.Split(logOutputs, ",")
	for _, output := range outputSlice {
		switch strings.TrimSpace(strings.ToLower(output)) {
		case "stdout":
			writers = append(writers, os.Stdout)
		case "file":
			if logFilePath != "" {
				if err := os.MkdirAll(filepath.Dir(logFilePath), 0755); err != nil {
					slog.Error("Failed to create log directory", "path", logFilePath, "error", err)
					continue
				}
				file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					slog.Error("Failed to open log file", "path", logFilePath, "error", err)
					continue
				}
				writers = append(writers, file)
			}
		}
	}
	if len(writers) == 0 {
		writers = append(writers, os.Stdout) // Default to stdout if no valid outputs are configured
	}
	multiWriter := io.MultiWriter(writers...)

	// --- Setup Log Handler ---
	var handler slog.Handler
	switch strings.ToUpper(logFormat) {
	case "JSON":
		handler = slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{Level: level})
	default:
		handler = slog.NewTextHandler(multiWriter, &slog.HandlerOptions{Level: level})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	logger.Info("Logger initialized", "level", level.String(), "format", logFormat, "outputs", logOutputs)

	return logger
}