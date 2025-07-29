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

	if err := v.Load(); err != nil {
		fmt.Println("Error loading vault:", err)
	}

	fmt.Println("Login successful.")
}

func handleFirstTime(v *vault.Vault) {
	password, _ := PromptNewPassword()

	if err := v.Crypto.SetupNewPassword(password); err != nil {
		fmt.Println("Error creating master password:", err)
	}

	metaData := v.Crypto.ToMeta()
	if err := v.FileIO.EnsureVaultDir(); err != nil {
		fmt.Println("Error initialising directory:", err)
	}

	if err := v.FileIO.WriteMeta(metaData); err != nil {
		fmt.Println("Error writing meta data:", err)
	}

	fmt.Println("Password created successfully.")
}
