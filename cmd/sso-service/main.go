package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso-service/internal/app"
	"sso-service/internal/config"
	myLog "sso-service/internal/lib/log"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	logger := myLog.MustLoad(cfg.Env)

	logger.Info("application sso-service is starting",
		slog.String("env", cfg.Env))

	application := app.New(*cfg, logger)

	application.MySql.MustConnect()
	go application.GRPC.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	application.GRPC.Stop()
	_ = application.MySql.Close()

	logger.Info("application sso-service stopped")
}
