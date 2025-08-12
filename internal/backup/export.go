package backup

import (
	"context"
	"os"
	"time"

	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/db"
	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/internal/vault"
	"github.com/Cyrof/govault/pkg/cli"
)

func Export(password string, v *vault.Vault, outPath string, keyOutPath string) {
	// generate aes to encrypt vault and meta data
	archiveAES, err := crypto.GenerateAES()
	if err != nil {
		cli.Error("Failed to generate AES cipher: %v\n", err)
		logger.Logger.Errorw("Failed to generate AES cipher", "error", err)
	}

	// generate new kdf to encrypt key.enc
	newSalt, err := crypto.GenerateSalt(32)
	if err != nil {
		cli.Error("Failed to generate salt")
		logger.Logger.Errorw("Failed to generate salt", "error", err)
	}
	derivedKey := crypto.KDF(password, newSalt)

	// snapshot db to a temp file
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	tmpDB, err := db.Snapshot(ctx, v.DB, os.TempDir())
	if err != nil {
		cli.Error("Failed to snapshot database: %v\n", err)
		logger.Logger.Errorw("snapshot db", "error", err)
	}
	defer os.Remove(tmpDB)

	// read snapshot + meta.json
	dbBytes, err := os.ReadFile(tmpDB)
	if err != nil {
		cli.Error("Failed to read DB snapshot: %v\n", err)
		logger.Logger.Errorw("read snapshot", "error", err)
	}
	metaBytes, err := os.ReadFile(v.FileIO.MetaPath)
	if err != nil {
		cli.Error("Failed to read metadata: %v\n", err)
		logger.Logger.Errorw("read meta", "error", err)
	}

	// encrypt both with archiveAES
	encDB, err := v.Crypto.Encrypt(dbBytes, &crypto.EncryptOptions{Key: archiveAES})
	if err != nil {
		cli.Error("Failed to encrypt DB: %v\n", err)
		logger.Logger.Errorw("encrypt meta", "error", err)
	}
	encMeta, err := v.Crypto.Encrypt(metaBytes, &crypto.EncryptOptions{Key: archiveAES})
	if err != nil {
		cli.Error("Failed to encrypt metadata: %v\n", err)
		logger.Logger.Errorw("encrypt meta", "error", err)
	}

	// write encrypted zip
	files := map[string][]byte{
		"vault.db.aes":  encDB,
		"meta.json.aes": encMeta,
	}

	if err := fileIO.WriteEncryptedZip(outPath, files); err != nil {
		cli.Error("Failed to created zip: %v\n", err)
		logger.Logger.Panicw("Failed to create zip", "error", err)
	}

	// encrypt aes key using master password and export it
	encryptedAESKey, err := v.Crypto.Encrypt(archiveAES, &crypto.EncryptOptions{Key: derivedKey})
	if err != nil {
		cli.Error("Failed to encrypt data: %v\n", err)
		logger.Logger.Errorw("Failed to encrypt data", "error", err)
	}
	keyFileContent := append(newSalt, encryptedAESKey...)

	if err := fileIO.WriteKeyFile(keyFileContent, keyOutPath); err != nil {
		cli.Error("Failed to export key file")
		logger.Logger.Errorw("Failed to export key file", "error", err)
	}

	cli.Success("Vault exported successfully")
	cli.Success("Key file exported successfully")
	logger.Logger.Info("Vault exported successfully")
	logger.Logger.Info("Key file exported successfully")
}
