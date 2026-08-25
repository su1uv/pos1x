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
	godotenv.Load()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	httpPort := flag.Int("port", port, "port listen on")
	flag.Parse()

	status := run(ctx, cancel, *httpPort)
	cancel()
	os.Exit(status)
}

func run(ctx context.Context, cancel context.CancelFunc, httpPort int) int {
	conn, err := pgx.Connect(ctx, fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRESS_PASSWORD"), os.Getenv("POSTGRES_DB"),
	))
	if err != nil {
		log.Fatal()
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
