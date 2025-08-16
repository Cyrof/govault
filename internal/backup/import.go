package backup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/db"
	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/internal/model"
	"github.com/Cyrof/govault/internal/vault"
)

func Import(password string, encZip string, keyFile string, v *vault.Vault) error {
	if v.DB == nil {
		return errors.New("database not initialise")
	}

	// read key file
	keyData, err := os.ReadFile(keyFile)
	if err != nil {
		return fmt.Errorf("read key file: %w", err)
	}

	if len(keyData) < 32 {
		return errors.New("key file is corrupted or incomplete")
	}

	salt := keyData[:32]
	cipherKey := keyData[32:]
	derivedKey := crypto.KDF(password, salt)

	archiveAES, err := v.Crypto.Decrypt(cipherKey, &crypto.DecryptOptions{Key: derivedKey})
	if err != nil {
		return fmt.Errorf("decrypt archive AES: %w", err)
	}

	// read encrypted zip
	files, err := fileIO.ReadEncryptedZip(encZip)
	if err != nil {
		return fmt.Errorf("read encrypted zip: %w", err)
	}

	var encDB, encMeta []byte
	var okDB, okMeta bool

	if encDB, okDB = files["vault.db.aes"]; !okDB {
		encDB, okDB = files["vault.enc.aes"] // legacy fallback
	}
	if encMeta, okMeta = files["meta.json.aes"]; !okMeta {
		return errors.New("archive missing meta.json.aes")
	}
	if !okDB {
		return errors.New("archive missing vault.db.aes (or legacy vault.enc.aes)")
	}

	// decrypt db snapshot to temp file
	dbBytes, err := v.Crypto.Decrypt(encDB, &crypto.DecryptOptions{Key: archiveAES})
	if err != nil {
		return fmt.Errorf("decrypt db snapshot: %w", err)
	}
	tmpDir := os.TempDir()
	tmpDB := filepath.Join(tmpDir, fmt.Sprintf("govault_import_%d.db", time.Now().UnixNano()))
	if err := os.WriteFile(tmpDB, dbBytes, 0o600); err != nil {
		return fmt.Errorf("write temp db: %w", err)
	}

	defer func() {
		if err := os.Remove(tmpDB); err != nil && !errors.Is(err, fs.ErrNotExist) {
			logger.Logger.Warnw("failed to remove temporary database", "path", tmpDB, "error", err)
		}
	}()

	// merge snapshot
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	conn, err := v.DB.Conn(ctx)
	if err != nil {
		return fmt.Errorf("get db conn: %w", err)
	}

	defer func() {
		if cerr := conn.Close(); cerr != nil {
			logger.Logger.Warnw("failed to close DB connection", "error", cerr)
		}
	}()

	// ensure live schema exists
	if err := db.SetupDatabase(v.DB); err != nil {
		return fmt.Errorf("ensure schema: %w", err)
	}

	// attach on connection
	if _, err := conn.ExecContext(ctx, `ATTACH DATABASE ? AS src`, tmpDB); err != nil {
		return fmt.Errorf("attach src db: %w", err)
	}

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		// detach before returning
		_, _ = conn.ExecContext(ctx, `DETACH DATABASE src`)
		return fmt.Errorf("begin tx: %w", err)
	}
	rollback := func(e error) error {
		_ = tx.Rollback()
		_, _ = conn.ExecContext(ctx, `DETACH DATABASE src`)
		return e
	}

	// replace contents
	if _, err := tx.ExecContext(ctx, `DELETE FROM main.secrets`); err != nil {
		return rollback(fmt.Errorf("clear main.secrets: %w", err))
	}
	if _, err := tx.ExecContext(ctx, `
		INSERT INTO main.secrets (id, key_name, value_ct, created_at, updated_at)
		SELECT id, key_name, value_ct, created_at, updated_at FROM src.secrets
	`); err != nil {
		return rollback(fmt.Errorf("copy secrets: %w", err))
	}

	// commit first then detach
	if err := tx.Commit(); err != nil {
		_, _ = conn.ExecContext(ctx, `DETACH DATABASE src`)
		return fmt.Errorf("commit import: %w", err)
	}
	if _, err := conn.ExecContext(ctx, `DETACH DATABASE src`); err != nil {
		return fmt.Errorf("detach src: %w", err)
	}

	// decrypt and write meta.json
	metaBytes, err := v.Crypto.Decrypt(encMeta, &crypto.DecryptOptions{Key: archiveAES})
	if err != nil {
		return fmt.Errorf("descryp meta: %w", err)
	}
	if err := v.FileIO.EnsureVaultDir(); err != nil {
		return fmt.Errorf("ensure vault dir: %w", err)
	}

	var meta model.Meta
	if err := json.Unmarshal(metaBytes, &meta); err != nil {
		return fmt.Errorf("unmarshal meta: %w", err)
	}
	if err := v.FileIO.WriteMeta(meta); err != nil {
		return fmt.Errorf("write meta: %w", err)
	}

	logger.Logger.Info("Import completed")
	return nil
}
