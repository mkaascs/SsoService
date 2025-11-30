package repositories

import (
	"context"
	"sso-service/internal/domain/dto/user/commands"
	"sso-service/internal/domain/dto/user/results"
)

type User interface {
	TxBeginner
	AddTx(ctx context.Context, tx Tx, command commands.Add) (*results.Add, error)
	GetByID(ctx context.Context, userID int64) (*results.Get, error)
	GetByLogin(ctx context.Context, command commands.GetByLogin) (*results.Get, error)
}
