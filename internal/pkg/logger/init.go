package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	options := slog.HandlerOptions{Level: slog.LevelInfo}

	h := slog.NewJSONHandler(os.Stdout, &options)

	logger := slog.New(h)

	return logger
}
