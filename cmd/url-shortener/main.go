package main

import (
	"URL-shortener/internal/config"
	"log/slog"
	_ "log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting server")
	//TODO init logger: slog

	//TODO: init storage: postgress

	//TODO: init router: chi, chi-render

	//TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
