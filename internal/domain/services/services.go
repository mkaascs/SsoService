package services

import (
	"context"
	"sso-service/internal/domain/models/commands"
	"sso-service/internal/domain/models/results"
)

type Auth interface {
	Register(ctx context.Context, register commands.Register) (results.Register, error)
}
