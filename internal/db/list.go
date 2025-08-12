package db

import (
	"context"
	"database/sql"
	"fmt"
)

func ListKeys(ctx context.Context, d *sql.DB) ([]string, error) {
	rows, err := d.QueryContext(ctx, `SELECT key_name FROM secrets ORDER BY key_name COLLATE NOCASE`)
	if err != nil {
		return nil, fmt.Errorf("ListKeys: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	var keys []string
	for rows.Next() {
		var k string
		if err := rows.Scan(&k); err != nil {
			return nil, fmt.Errorf("ListKey scan: %w", err)
		}
		keys = append(keys, k)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ListKeys rows: %w", err)
	}
	return keys, nil
}
