package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func GetSecretCT(ctx context.Context, d *sql.DB, keyName string) ([]byte, error) {
	var ct []byte
	err := d.QueryRowContext(ctx, `SELECT value_ct FROM secrets WHERE key_name = ?`, keyName).Scan(&ct)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("GetSecretCT: %w", err)
	}
	return ct, nil
}
