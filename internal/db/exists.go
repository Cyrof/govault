package db

import (
	"context"
	"database/sql"
)

func KeyExists(ctx context.Context, d *sql.DB, keyName string) (bool, error) {
	var x int
	err := d.QueryRowContext(ctx, `SELECT 1 FROM secrets WHERE key_name = ? LIMIT 1`, keyName).Scan(&x)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}
