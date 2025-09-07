package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kevinsoras/employee-management/shared/domain"
)

// txKey is an unexported type to be used as a key for storing the transaction in the context.
type txKey struct{}

// Querier defines the common methods for sql.DB and sql.Tx, allowing repositories
// to work with both transactions and regular connections.
type Querier interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// PostgresUoW implements the UnitOfWork for PostgreSQL.
type PostgresUoW struct {
	db *sql.DB
}

// NewPostgresUoW creates a new PostgresUoW.
func NewPostgresUoW(db *sql.DB) *PostgresUoW {
	return &PostgresUoW{db: db}
}

// Execute runs the given function within a single atomic transaction.
func (uow *PostgresUoW) Execute(ctx context.Context, fn domain.UowCallback) error {
	tx, err := uow.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Create a new context with the transaction object.
	txCtx := context.WithValue(ctx, txKey{}, tx)

	// Execute the callback with the transactional context.
	err = fn(txCtx)
	if err != nil {
		// If the callback returns an error, roll back the transaction.
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction failed: %v, and rollback failed: %w", err, rbErr)
		}
		return fmt.Errorf("transaction failed: %w", err)
	}

	// If the callback succeeds, commit the transaction.
	return tx.Commit()
}

// GetQuerier extracts a Querier (either a *sql.Tx or *sql.DB) from the context.
// If a transaction (*sql.Tx) is present in the context, it is returned.
// Otherwise, it returns the provided database connection pool (*sql.DB).
func GetQuerier(ctx context.Context, db *sql.DB) Querier {
	tx, ok := ctx.Value(txKey{}).(*sql.Tx)
	if ok {
		return tx
	}
	return db
}
