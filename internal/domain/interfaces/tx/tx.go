package tx

import (
	"context"
)

type Tx interface {
	Rollback() error
	Commit() error
}

type Beginner interface {
	BeginTx(ctx context.Context) (Tx, error)
}
