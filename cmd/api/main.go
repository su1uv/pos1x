package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/su1uv/pos1x/internal/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env vars: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("error converting port to int: %v", err)
	}

	httpPort := flag.Int("port", port, "port listen on")
	flag.Parse()

	status := run(ctx, cancel, *httpPort)
	cancel()
	os.Exit(status)
}

func run(ctx context.Context, cancel context.CancelFunc, httpPort int) int {
	conn, err := pgx.Connect(ctx, fmt.Sprintf(
		"postgres://%v:%v@%v:5432/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("HOST"), os.Getenv("POSTGRES_DB"),
	))
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	s := newServer(httpPort, cancel, queries)
	var serverErr error
	go func() {
		serverErr = s.start()
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer func() {
		cancel()
	}()

	if err := s.shutdown(shutdownCtx); err != nil {
		return 1
	}
	if serverErr != nil {
		return 1
	}
	return 0
}
