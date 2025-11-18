package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"numbsort/internal/config"
	"numbsort/internal/db"
	"numbsort/internal/handler"
	"numbsort/internal/repository"
	"numbsort/internal/routes"
	"numbsort/internal/service"
)

func main() {
	cfg := config.MustLoad()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer pool.Close()

	if err := db.EnsureSchema(ctx, pool); err != nil {
		log.Fatalf("failed to ensure schema: %v", err)
	}

	svc := service.NewNumberService(repository.NewNumberRepository(pool))
	h := handler.NewNumberHandler(svc)
	router := routes.SetupRouter(h)

	server := routes.NewServer(cfg.HTTPAddr, router)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("server start failed: %v", err)
		}
	}()

	log.Printf("server listening on %s", cfg.HTTPAddr)

	<-ctx.Done()
	log.Println("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Stop(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}
