package backup

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/db"
	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/internal/vault"
)

func Export(password string, v *vault.Vault, outPath string, keyOutPath string) error {
	// generate aes to encrypt vault and meta data
	archiveAES, err := crypto.GenerateAES()
	if err != nil {
		return fmt.Errorf("generate AES cipher: %w", err)
	}

	// generate new kdf to encrypt key.enc
	newSalt, err := crypto.GenerateSalt(32)
	if err != nil {
		return fmt.Errorf("generate salt: %w", err)
	}
	derivedKey := crypto.KDF(password, newSalt)

	// snapshot db to a temp file
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	tmpDB, err := db.Snapshot(ctx, v.DB, os.TempDir())
	if err != nil {
		return fmt.Errorf("snapshot database: %w", err)
	}
	defer os.Remove(tmpDB)

	// read snapshot + meta.json
	dbBytes, err := os.ReadFile(tmpDB)
	if err != nil {
		return fmt.Errorf("read DB snapshot: %w", err)
	}
	metaBytes, err := os.ReadFile(v.FileIO.MetaPath)
	if err != nil {
		return fmt.Errorf("read metadata: %w", err)
	}

	// encrypt both with archiveAES
	encDB, err := v.Crypto.Encrypt(dbBytes, &crypto.EncryptOptions{Key: archiveAES})
	if err != nil {
		return fmt.Errorf("encrypt DB: %w", err)
	}
	encMeta, err := v.Crypto.Encrypt(metaBytes, &crypto.EncryptOptions{Key: archiveAES})
	if err != nil {
		return fmt.Errorf("encrypt metadata: %w", err)
	}

	// write encrypted zip
	files := map[string][]byte{
		"vault.db.aes":  encDB,
		"meta.json.aes": encMeta,
	}

	if err := fileIO.WriteEncryptedZip(outPath, files); err != nil {
		return fmt.Errorf("create encrypted zip: %w", err)
	}

	// encrypt aes key using master password and export it
	encryptedAESKey, err := v.Crypto.Encrypt(archiveAES, &crypto.EncryptOptions{Key: derivedKey})
	if err != nil {
		return fmt.Errorf("encrypt AES key: %w", err)
	}
	keyFileContent := append(newSalt, encryptedAESKey...)

	if err := fileIO.WriteKeyFile(keyFileContent, keyOutPath); err != nil {
		return fmt.Errorf("export key file: %w", err)
	}

	logger.Logger.Info("Vault exported successfully")
	logger.Logger.Info("Key file exported successfully")
	return nil
}
