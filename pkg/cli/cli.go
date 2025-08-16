package cli

import (
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/internal/vault"
)

func Setup(v *vault.Vault) {
	if v.FileIO.CheckMetaFile() {
		handleExistingUser(v)
	} else {
		handleFirstTime(v)
	}
}

func handleExistingUser(v *vault.Vault) {
	salt, hash, err := v.FileIO.ReadMeta()
	if err != nil {
		Error("Failed to read metadata.\nExiting...\n\n")
		logger.Logger.Fatalw("Failed to load meta", "path", v.FileIO.MetaPath, "error", err)
	}

	password, _ := PromptPassword()

	err = v.Crypto.SetupFromMeta(password, salt, hash)
	if err != nil {
		Error("Login Failed.\n\n")
		logger.Logger.Fatalw("Failed to login", "error", err)
	}

	logger.Logger.Info("Login successful.")
	Success("Login successful.\n\n")
}

func handleFirstTime(v *vault.Vault) {
	password, _ := PromptNewPassword()

	if err := v.Crypto.SetupNewPassword(password); err != nil {
		Error("Error creating password.\n\n")
		logger.Logger.Panicw("Error creating master password", "error", err)
		return
	}

	metaData := v.Crypto.ToMeta()
	if err := v.FileIO.EnsureVaultDir(); err != nil {
		Error("Error initialising directory.\n\n")
		logger.Logger.Warnw("Error initialising directory", "error", err)
	}

	if err := v.FileIO.WriteMeta(metaData); err != nil {
		Error("Error writing meta data.\n\n")
		logger.Logger.Panicw("Error writing meta data", "error", err)
		return
	}

	logger.Logger.Info("Password created successfully.")
	Success("Password created successfully.\n\n")
}
