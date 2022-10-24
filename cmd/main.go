package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"task/rest_service/api"
	"task/rest_service/api/handlers"
	"task/rest_service/config"
	"task/rest_service/storage/postgres"

	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {
	var loggerLevel string
	cfg := config.Load()

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)

	store, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer store.CloseDB()

	h := handlers.NewHandler(cfg, log, store)

	r := api.SetUpRouter(h, cfg)

	go func() {
		if r.Listen(cfg.HTTPPort) != nil {
			panic(err)
		}
	}()

	shutdownChan := make(chan os.Signal, 1)
	defer close(shutdownChan)
	signal.Notify(shutdownChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-shutdownChan

	log.Info("received os signal", logger.Any("signal", sig))
	if err := r.Shutdown(); err != nil {
		log.Error("could not shutdown http server", logger.Error(err))
		return
	}

	log.Info("server shutdown successfully")
}
