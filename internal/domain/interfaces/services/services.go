package services

import (
	"context"
	"sso-service/internal/domain/dto/commands"
	"sso-service/internal/domain/dto/results"
)

type Auth interface {
	Register(ctx context.Context, command commands.Register) (*results.Register, error)
	Login(ctx context.Context, command commands.Login) (*results.Login, error)
	Refresh(ctx context.Context, command commands.Refresh) (*results.Refresh, error)
}

type JWT interface {
	Generate(command commands.Generate) (*results.Generate, error)
	Parse(command commands.Parse) (*results.Parse, error)
}
