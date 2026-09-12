package handlers

import (
	"log/slog"
	"net/http"

	"github.com/su1uv/pos1x/internal/services"
)

type Handlers struct {
	services *services.Services
	logger   *slog.Logger
}

func NewHandlers(services *services.Services, logger *slog.Logger) *Handlers {
	return &Handlers{services: services, logger: logger}
}

func (h *Handlers) Health(w http.ResponseWriter, _ *http.Request) {
	err := respondWithJSON(w, 200, map[string]string{"status": "ok"})
	if err != nil {
		h.logger.Error("something went wrong")
	}
}
