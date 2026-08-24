package telemetry

import (
	"log/slog"
	"os"
)

// NewLogger returns a structured JSON logger at the given level ("debug",
// "info", "warn", "error"), suitable for container log collection.
func NewLogger(level string) *slog.Logger {
	var l slog.Level
	if err := l.UnmarshalText([]byte(level)); err != nil {
		l = slog.LevelInfo
	}
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l})
	return slog.New(handler)
}
