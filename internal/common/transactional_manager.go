package common

import (
	"context"
	"database/sql"
)

// WithExecuteTransactionalManager manages a insert or update transaction
func WithExecuteTransactionalManager(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}

// WithReadOnlyTransactionManager manages a read-only transaction
func WithReadOnlyTransactionManager(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	// Begin a read-only transaction
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: true,
	})
	if err != nil {
		return err
	}

	// Ensure transaction rollback in case of panic or error
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Execute the provided function with the transaction
	err = fn(tx)
	return err
}
