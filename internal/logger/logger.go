package logger

import (
	"log/slog"
	"os"
)

func RequestLogger() *slog.Logger {
	l := slog.New(slog.NewTextHandler(os.Stderr, nil))

	return l
}
