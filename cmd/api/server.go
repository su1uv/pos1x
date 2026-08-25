package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/su1uv/pos1x/internal/db"
)

type Server struct {
	httpServer *http.Server
	db         *db.Queries
	cancel     context.CancelFunc
}

func newServer(port int, cancel context.CancelFunc, queries *db.Queries) *Server {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	s := &Server{
		httpServer: srv,
		db:         queries,
		cancel:     cancel,
	}

	// TODO: register endpoints

	return s
}

func (s *Server) start() error {
	ln, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		return err
	}
	if err := s.httpServer.Serve(ln); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
