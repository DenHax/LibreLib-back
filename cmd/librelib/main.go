package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DenHax/LibreLib-back/internal/api"
	"github.com/DenHax/LibreLib-back/internal/config"
	"github.com/DenHax/LibreLib-back/internal/repo"
	"github.com/DenHax/LibreLib-back/internal/server"
	"github.com/DenHax/LibreLib-back/internal/service"
	"github.com/DenHax/LibreLib-back/internal/storage/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func main() {
	fmt.Println("LibreLib Backend")

	cfg := config.MustLoad()

	log := setupLogger(cfg.Logger.Env)
	log.Info(
		"start logger in",
		slog.String("env", cfg.Logger.Env),
	)

	log.Debug("storage", cfg.Storage)
	log.Debug("server", cfg.Server)

	storage, err := postgres.New(cfg.Storage.URL)
	if err != nil {
		log.Error("failed to init storage", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer storage.Close()

	repos := repo.NewRepository(storage)

	service := service.NewService(repos)
	handlers := api.NewHandler(service)

	// r := routers.Api(storage)
	// http.ListenAndServe(":"+cfg.Server.Port, r)
	// TODO: auth
	// TODO: auth-session
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("starting server", slog.String("address", cfg.Server.Address))

	srv := server.New(cfg.Server, handlers.Init())

	go func() {
		if err := srv.Run(); err != nil {
			log.Error("failed to stop server", slog.String("error", err.Error()))
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", slog.String("error", err.Error()))
		return
	}
}
