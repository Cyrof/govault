package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func Update(ctx context.Context, d *sql.DB, keyName string, valueCT []byte) error {
	now := time.Now().Unix()
	res, err := d.ExecContext(ctx, `
		UPDATE secrets
			SET value_ct = ?, updated_at = ?
		WHERE key_name = ?
	`, valueCT, now, keyName)

	if err != nil {
		return fmt.Errorf("UpdateSecret: %w", err)
	}

	if n, _ := res.RowsAffected(); n == 0 {
		return ErrNotFound
	}
	return nil
}
