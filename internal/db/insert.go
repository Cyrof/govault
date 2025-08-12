package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func InsertSecret(ctx context.Context, d *sql.DB, keyName string, vaultCT []byte) error {
	now := time.Now().Unix()
	_, err := d.ExecContext(ctx, `
		INSERT INTO secrets (key_name, value_ct, created_at, updated_at)
		VALUES (?, ?, ?, ?)
	`, keyName, vaultCT, now, now)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return ErrDuplicateKey
		}
		return fmt.Errorf("InsertSecret: %w", err)
	}

	return nil
}
