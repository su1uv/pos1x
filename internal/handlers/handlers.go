package handlers

import (
	"net/http"

	"github.com/su1uv/pos1x/internal/services"
)

type Handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{services: services}
}

func (h *Handlers) Health(w http.ResponseWriter, _ *http.Request) {
	respondWithJSON(w, 200, map[string]string{"status": "ok"})
}
