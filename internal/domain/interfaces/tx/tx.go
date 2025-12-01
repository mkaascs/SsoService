package repositories

import (
	"context"
)

type Tx interface {
	Rollback() error
	Commit() error
}

type TxBeginner interface {
	BeginTx(ctx context.Context) (Tx, error)
}
