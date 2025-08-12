package db

import (
	"context"
	"database/sql"
	"fmt"
)

func Delete(ctx context.Context, d *sql.DB, keyName string) error {
	res, err := d.ExecContext(ctx, `DELETE FROM secrets WHERE key_name = ?`, keyName)
	if err != nil {
		return fmt.Errorf("DeleteSecret: %w", err)
	}
	// check if rows exist
	if n, _ := res.RowsAffected(); n == 0 {
		return ErrNotFound
	}
	return nil
}
