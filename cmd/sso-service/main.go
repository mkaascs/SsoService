package main

import (
	"log/slog"
	"sso-service/internal/app"
	"sso-service/internal/config"
	myLog "sso-service/internal/lib/log"
)

func main() {
	cfg := config.MustLoad()
	logger := myLog.MustLoad(cfg.Env)

	logger.Info("application sso-service is starting",
		slog.String("env", cfg.Env))

	application := app.New(*cfg, logger)
	application.GRPC.MustRun()
}
