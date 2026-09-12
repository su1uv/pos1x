package logger

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

func RequestLogger() *slog.Logger {
	withoutColor := true
	if isatty.IsCygwinTerminal(os.Stderr.Fd()) || isatty.IsTerminal(os.Stderr.Fd()) {
		withoutColor = false
	}

	debugHandler := tint.NewTextHandler(os.Stderr, &tint.Options{
		Level: slog.LevelDebug,
		NoColor: withoutColor,
	})


	l := slog.New(debugHandler)

	return l
}

func LoggerMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(n http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
		
			n.ServeHTTP(w, r)

			logAttrs := []any{
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Duration("duration", time.Since(start)),
			}

			logger.Info("Server request", logAttrs...)
		})
	}
}
