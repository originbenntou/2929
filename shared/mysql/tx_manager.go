package mysql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// TxManager is the manager of Tx.
type TxManager interface {
	Executor
	Commit() error
	Rollback() error
}

// txManager is the manager of Tx.
type txManager struct {
	*sqlx.Tx
}

// Executor implement
// ExecContext executes SQL with context on sqlx.
func (s *txManager) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return s.Tx.NamedExecContext(ctx, query, arg)
}

// Tx Commit
func (s *txManager) Commit() error {
	return s.Tx.Commit()
}

// Tx Rollback
func (s *txManager) Rollback() error {
	return s.Tx.Rollback()
}

// Close Tx
// Don't write panic() not to stop server
func CloseTransaction(tx TxManager, err error) error {
	var txErr error
	if recover() != nil {
		txErr = tx.Rollback()
	} else if err != nil {
		txErr = tx.Rollback()
	} else {
		txErr = tx.Commit()
	}

	if txErr != nil {
		err = txErr
	}

	return err
}
