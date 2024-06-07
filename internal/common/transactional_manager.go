package common

import (
	"context"
	"database/sql"
	"fmt"
)

func WithTransactionalManager(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
			panic(p)
		} else if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		} else {
			fmt.Printf("MASUK SINI")
			err = tx.Commit()
		}
	}()

	return fn(tx)
}
