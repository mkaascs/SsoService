package log

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

var (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func MustLoad(env string) *slog.Logger {
	logger, err := Load(env)
	if err != nil {
		log.Fatal(err)
	}

	return logger
}

func Load(env string) (*slog.Logger, error) {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		return nil, fmt.Errorf("unknown environment: %s", env)
	}

	return logger, nil
}
