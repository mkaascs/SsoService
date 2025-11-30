package repositories

import (
	"context"
	"sso-service/internal/domain/dto/token/commands"
	"sso-service/internal/domain/dto/token/results"
)

type RefreshToken interface {
	TxBeginner
	AddTx(ctx context.Context, tx Tx, command commands.Add) error
	UpdateByTokenTx(ctx context.Context, tx Tx, command commands.UpdateByToken) (*results.Update, error)
	UpdateByUserIDTx(ctx context.Context, tx Tx, command commands.UpdateByUserID) (*results.Update, error)
}
