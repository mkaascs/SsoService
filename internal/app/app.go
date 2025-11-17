package app

import (
	"log/slog"
	grpcapp "sso-service/internal/app/grpc"
	"sso-service/internal/app/mysql"
	"sso-service/internal/config"
)

type App struct {
	GRPC  *grpcapp.App
	MySql *mysql.App
}

func New(cfg config.Config, logger *slog.Logger) *App {
	grpcApp := grpcapp.New(logger, cfg.Port)
	mysqlApp, _ := mysql.New(logger, cfg.ConnectionString)

	return &App{
		GRPC:  grpcApp,
		MySql: mysqlApp,
	}
}
