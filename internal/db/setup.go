package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"sort"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	d, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	// apply runtime settings
	pragmas := []string{
		`PRAGMA journal_mode=WAL;`,
		`PRAGMA synchronous=NORMAL;`,
		`PRAGMA foreign_keys=ON;`,
		`PRAGMA busy_timeout=5000;`,
		`PRAGMA temp_store=MEMORY;`,
	}
	for _, p := range pragmas {
		if _, perr := d.Exec(p); perr != nil {
			_ = d.Close()
			return nil, fmt.Errorf("apply %s: %w", p, perr)
		}
	}

	if err := d.Ping(); err != nil {
		_ = d.Close()
		return nil, err
	}
	return d, nil
}

//go:embed schema/*.sql
var schemaFS embed.FS

func SetupDatabase(d *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var current int
	if err := d.QueryRowContext(ctx, `PRAGMA user_version;`).Scan(&current); err != nil {
		return fmt.Errorf("read user_version: %w", err)
	}

	entries, err := schemaFS.ReadDir("schema")
	if err != nil {
		return err
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			files = append(files, e.Name())
		}
	}
	sort.Strings(files)

	for _, name := range files {
		v, err := versionFromName(name)
		if err != nil {
			return err
		}

		if v <= current {
			continue
		}

		sqlBytes, err := schemaFS.ReadFile("schema/" + name)
		if err != nil {
			return err
		}

		tx, err := d.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, string(sqlBytes)); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("schema step %s failed: %w", name, err)
		}

		// ensure version bump even if file forget to set
		if _, err := tx.ExecContext(ctx, fmt.Sprintf("PRAGMA user_version = %d", v)); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("set user_version %d failed: %w", v, err)
		}

		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}

func versionFromName(name string) (int, error) {
	if len(name) < 4 || name[0] < '0' || name[0] > '9' {
		return 0, nil
	}

	var v int
	if _, err := fmt.Sscanf(name[:4], "%d", &v); err != nil {
		return 0, fmt.Errorf("parse %q as int: %w", name[:4], err)
	}
	return v, nil
}
