package repositories

import (
	"context"
	"sso-service/internal/domain/dto/user/commands"
	"sso-service/internal/domain/dto/user/results"
)

type User interface {
	BeginTx(ctx context.Context) (Tx, error)
	AddTx(ctx context.Context, tx Tx, command commands.Add) (*results.Add, error)
}
