package repositories

type Tx interface {
	Rollback() error
	Commit() error
}
