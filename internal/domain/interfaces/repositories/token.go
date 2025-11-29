package repositories

import (
	"context"
	"sso-service/internal/domain/dto/token/commands"
	"sso-service/internal/domain/dto/token/results"
)

type RefreshToken interface {
	AddTx(ctx context.Context, tx Tx, command commands.Add) error
	Update(ctx context.Context, command commands.Update) (*results.Update, error)
}
