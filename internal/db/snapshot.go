package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Snapshot(ctx context.Context, d *sql.DB, dir string) (string, error) {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", fmt.Errorf("snapshot mkdir: %w", err)
	}
	tmp := filepath.Join(dir, fmt.Sprintf("govalt_export_%d.db", time.Now().UnixNano()))

	// if _, err := d.ExecContext(ctx, `PRAGMA wal_checkpoint(TRUNCATE);`); err != nil {
	//
	// }
	if _, err := d.ExecContext(ctx, `VACUUM INTO ?;`, tmp); err != nil {
		return "", fmt.Errorf("VACUUM INTO: %w", err)
	}
	return tmp, nil
}
