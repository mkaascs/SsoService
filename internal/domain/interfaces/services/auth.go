package services

import (
	"context"
	"sso-service/internal/domain/dto/auth/commands"
	"sso-service/internal/domain/dto/auth/results"
)

type Auth interface {
	Register(ctx context.Context, command commands.Register) (*results.Register, error)
	Login(ctx context.Context, command commands.Login) (*results.Login, error)
	Refresh(ctx context.Context, command commands.Refresh) (*results.Refresh, error)
}
