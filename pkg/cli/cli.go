package cli

import (
	"fmt"
	"os"

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
		fmt.Println("Failed to read meta:", err)
		os.Exit(1)
	}

	password, _ := PromptPassword()

	err = v.Crypto.SetupFromMeta(password, salt, hash)
	if err != nil {
		fmt.Println("Password Invalid:", err)
		os.Exit(1)
	}

	v.Load()
	fmt.Println("Login successful.")
}

func handleFirstTime(v *vault.Vault) {
	password, _ := PromptNewPassword()

	v.Crypto.SetupNewPassword(password)
	metaData := v.Crypto.ToMeta()
	v.FileIO.EnsureVaultDir()
	v.FileIO.WriteMeta(metaData)
	fmt.Println("Password created successfully.")
}
