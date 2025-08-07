package backup

import (
	"github.com/Cyrof/govault/internal/crypto"
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

	// read and encrypt both file
	vaultData, metaData, err := v.FileIO.ReadAll()
	if err != nil {
		cli.Error("Error reading file: %v\n", err)
		logger.Logger.Errorw("Error reading file", "error", err)
	}

	encryptVaultData, err := v.Crypto.Encrypt(vaultData, &crypto.EncryptOptions{Key: archiveAES})
	if err != nil {
		cli.Error("Failed to encrypt data: %v\n", err)
		logger.Logger.Errorw("Failed to encrypt data", "error", err)
	}
	encryptMetaData, err := v.Crypto.Encrypt(metaData, &crypto.EncryptOptions{Key: archiveAES})
	if err != nil {
		cli.Error("Failed to encrypt data: %v\n", err)
		logger.Logger.Errorw("Failed to encrypt data", "error", err)
	}

	files := map[string][]byte{
		"vault.enc.aes": encryptVaultData,
		"meta.json.aes": encryptMetaData,
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
