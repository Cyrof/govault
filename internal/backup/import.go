package backup

import (
	"encoding/json"
	"os"

	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/internal/model"
	"github.com/Cyrof/govault/internal/vault"
	"github.com/Cyrof/govault/pkg/cli"
)

func Import(password string, encZip string, keyFile string, v *vault.Vault) {
	// read key file
	keyData, err := os.ReadFile(keyFile)
	if err != nil {
		cli.Error("Failed to read key file: %v", err)
		logger.Logger.Errorw("Failed to read key file", "error", err)
	}

	if len(keyData) < 32 {
		cli.Error("Key file is corrupted or incomplete")
		logger.Logger.Error("Key file is corrupted or incomplete")
	}

	salt := keyData[:32]
	cipherKey := keyData[32:]
	derivedKey := crypto.KDF(password, salt)

	archiveAES, err := v.Crypto.Decrypt(cipherKey, &crypto.DecryptOptions{Key: derivedKey})
	if err != nil {
		cli.Error("Failed to decrypt AES key: %v", err)
		logger.Logger.Errorw("Failed to decrypt AES key", "error", err)
	}

	// read encrypted zip
	files, err := fileIO.ReadEncryptedZip(encZip)
	if err != nil {
		cli.Error("Failed to read zip archive: %v", err)
		logger.Logger.Errorw("Failed to read zip archive", "error", err)
	}

	vaultEnc, ok1 := files["vault.enc.aes"]
	metaEnc, ok2 := files["meta.json.aes"]
	if !ok1 || !ok2 {
		cli.Error("Missing required files in archive")
		logger.Logger.Error("Missing required files in archive")
	}

	// decrypt vault + meta
	vaultData, err := v.Crypto.Decrypt(vaultEnc, &crypto.DecryptOptions{Key: archiveAES})
	if err != nil {
		cli.Error("Failed to decrypt vault data: %v", err)
		logger.Logger.Errorw("Failed to decrypt vault data", "error", err)
	}

	metaData, err := v.Crypto.Decrypt(metaEnc, &crypto.DecryptOptions{Key: archiveAES})
	if err != nil {
		cli.Error("Failed to decrypt meta data: %v", err)
		logger.Logger.Errorw("Failed to decrypt meta data", "error", err)
	}

	// write to file
	if err := v.FileIO.EnsureVaultDir(); err != nil {
		cli.Error("Failed to initialise vault directory: %v\n", err)
		logger.Logger.Panicw("Failed to initialise vault directory", "error", err)
	}

	if err := v.FileIO.WriteSecret(vaultData); err != nil {
		cli.Error("Failed to write vault: %v", err)
		logger.Logger.Errorw("Failed to write to vault", "error", err)
	}
	cli.Success("Vault restored successfully.\n")
	logger.Logger.Info("Vault restored successfully")

	var meta model.Meta
	if err := json.Unmarshal(metaData, &meta); err != nil {
		cli.Error("Failed to unmarshal meta: %v", err)
		logger.Logger.Errorw("Failed to unmarshal meta", "error", err)
	}
	if err := v.FileIO.WriteMeta(meta); err != nil {
		cli.Error("Failed to write meta: %v", err)
		logger.Logger.Errorw("Failed to write meta", "error", err)
	}
	cli.Success("Meta data restored successfully.\n\n")
	logger.Logger.Info("Meta data restored successfully.")

	cli.Success("You may now resume using your vault with your original master password.")
}
