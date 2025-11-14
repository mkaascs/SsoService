package app

import (
	"log/slog"
	grpcapp "sso-service/internal/app/grpc"
	"sso-service/internal/config"
)

type App struct {
	GRPC *grpcapp.App
}

func New(cfg config.Config, logger *slog.Logger) *App {
	grpcApp := grpcapp.New(logger, cfg.Port)

	return &App{
		GRPC: grpcApp,
	}
}
