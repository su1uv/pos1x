package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/su1uv/pos1x/internal/db"
	"github.com/su1uv/pos1x/internal/handlers"
	"github.com/su1uv/pos1x/internal/services"
)

type server struct {
	httpServer *http.Server
	cancel     context.CancelFunc
	logger     *slog.Logger
}

func newServer(port int, cancel context.CancelFunc, logger *slog.Logger, queries *db.Queries) *server {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	s := &server{
		httpServer: srv,
		cancel:     cancel,
		logger:     logger,
	}

	services := services.NewServices(queries)
	h := handlers.NewHandlers(services)

	mux.HandleFunc("GET /health", h.Health)

	return s
}

func (s *server) start() error {
	ln, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		return err
	}
	if err := s.httpServer.Serve(ln); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *server) shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
