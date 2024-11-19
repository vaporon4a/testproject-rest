package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"testproject-rest/internal/config"
	"testproject-rest/internal/http-server/handlers/wallet/balance"
	"testproject-rest/internal/http-server/handlers/wallet/operation"
	midlogger "testproject-rest/internal/http-server/middleware/logger"
	"testproject-rest/internal/lib/logger/slhelper"
	"testproject-rest/internal/storage/pgsql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := initLogger(cfg.Env)

	log.Info("starting API", slog.String("env", cfg.Env))
	log.Debug("debug messages enabled")

	storage, err := pgsql.New(fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DB.PostgresHost, cfg.DB.PostgresPort,
		cfg.DB.PostgresUser, cfg.DB.PostgresPass,
		cfg.DB.PostgresDBName))
	if err != nil {
		log.Error("failed to init storage", slhelper.Err(err))
		os.Exit(1)
	}
	defer storage.Close()

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(midlogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/api/v1/wallet", operation.UseWallet(log, storage))
	router.Get("/api/v1/wallets/{walletId}", balance.ShowBalance(log, storage))

	log.Info("starting server", slog.String("address", cfg.ApiAddres))

	srv := &http.Server{
		Addr:         cfg.ApiAddres,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.ApiTimeout,
		WriteTimeout: cfg.HTTPServer.ApiTimeout,
		IdleTimeout:  cfg.HTTPServer.ApiIdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server", slhelper.Err(err))
		os.Exit(1)
	}

	log.Error("server stopped")
}

func initLogger(env string) *slog.Logger {
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
