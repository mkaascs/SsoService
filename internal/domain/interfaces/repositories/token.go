package repositories

import (
	"context"
	"sso-service/internal/domain/dto/token/commands"
	"sso-service/internal/domain/dto/token/results"
	"sso-service/internal/domain/interfaces/tx"
)

type RefreshToken interface {
	tx.Beginner
	AddTx(ctx context.Context, tx tx.Tx, command commands.Add) error
	UpdateByTokenTx(ctx context.Context, tx tx.Tx, command commands.UpdateByToken) (*results.Update, error)
	UpdateByUserIDTx(ctx context.Context, tx tx.Tx, command commands.UpdateByUserID) (*results.Update, error)
}
