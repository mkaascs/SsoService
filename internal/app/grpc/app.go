package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"sso-service/internal/delivery/grpc/auth"
	sloglib "sso-service/internal/lib/log/slog"
)

type App struct {
	server *grpc.Server
	port   int
	logger *slog.Logger
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		a.logger.Error("failed to run grpc server", sloglib.Error(err))
		os.Exit(1)
	}
}

func (a *App) Run() error {
	const fn = "app.grpc.App.Run"

	a.logger.With(slog.String("fn", fn))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		a.logger.Error("failed to listen tcp", sloglib.Error(err))
		return fmt.Errorf("%s: failed to listen tcp: %w", fn, err)
	}

	a.logger.Info("grpc server is running", slog.Int("port", a.port))

	if err := a.server.Serve(listener); err != nil {
		a.logger.Error("failed to serve", sloglib.Error(err))
		return fmt.Errorf("%s: failed to serve: %w", fn, err)
	}

	return nil
}

func (a *App) Stop() {
	const fn = "app.grpc.App.Stop"

	a.logger.With(slog.String("fn", fn))
	a.logger.Info("grpc server is stopping", slog.Int("port", a.port))

	a.server.GracefulStop()
}

func New(logger *slog.Logger, port int) *App {
	server := grpc.NewServer()

	auth.Register(server)

	return &App{
		server: server,
		port:   port,
		logger: logger,
	}
}
