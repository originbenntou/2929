package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DBManager interface {
	SQLManager
	Beginner
}

// TxManager is the manager of Tx.
type TxManager interface {
	SQLManager
	Commit() error
	Rollback() error
}

// SQLManager is the manager of DB.
type SQLManager interface {
	Querier
	Preparer
	Executor
}

type (
	// Executor is interface of Execute.
	Executor interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	}

	// Preparer is interface of Prepare.
	Preparer interface {
		Prepare(query string) (*sql.Stmt, error)
		PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	}

	// Querier is interface of Query.
	Querier interface {
		Query(query string, args ...interface{}) (*sql.Rows, error)
		QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	}

	// Beginner is interface of Begin.
	Beginner interface {
		Begin() (TxManager, error)
	}
)

// dbManager is the manager of SQL.
type dbManager struct {
	Conn *sql.DB
}

// NewDBManager generates and returns DBManager.
func NewDBManager(conn *sql.DB) DBManager {
	return &dbManager{conn}
}

// Exec executes SQL.
func (s dbManager) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.Conn.Exec(query, args...)
}

// ExecContext executes SQL with context.
func (s *dbManager) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.Conn.ExecContext(ctx, query, args...)
}

// Query executes query which return row.
func (s *dbManager) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.Conn.Query(query, args...)
	if err != nil {
		//err = &model.SQLError{
		//	BaseErr:                   err,
		//	InvalidReasonForDeveloper: "failed to execute query",
		//}
		return nil, err
	}

	return rows, nil
}

// QueryContext executes query which return row with context.
func (s *dbManager) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.Conn.Query(query, args...)
	if err != nil {
		//err = &model.SQLError{
		//	BaseErr:                   err,
		//	InvalidReasonForDeveloper: "failed to execute query with context",
		//}
		return nil, err
	}
	return rows, nil
}

// Prepare prepares statement for Query and Exec later.
func (s *dbManager) Prepare(query string) (*sql.Stmt, error) {
	return s.Conn.Prepare(query)
}

// Prepare prepares statement for Query and Exec later with context.
func (s *dbManager) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return s.Conn.PrepareContext(ctx, query)
}

// Begin begins tx.
func (s *dbManager) Begin() (TxManager, error) {
	return s.Conn.Begin()
}
