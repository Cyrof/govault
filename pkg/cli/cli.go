package cli

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Cyrof/govault/internal/vault"
)

func Setup(v *vault.Vault) {
	if v.FileIO.CheckMetaFile() {
		// setup crypto salt and hash
		salt, hash, err := v.FileIO.ReadMeta()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// call prompt pw cli to get master password from user
		password, _ := PromptPassword()

		// setup crypto salt and hash
		err = v.Crypto.SetupFromMeta(password, salt, hash)
		if err != nil {
			log.Fatal("Invalid password:", err)
			os.Exit(1)
		}
		v.Load()
		fmt.Println("Login successful.")
		time.Sleep(1 * time.Second) // 1 second pause
		ClearScreen()
	} else {
		// prompt welcome message
		PrintWelcome()

		// prompt to generate new password
		password, _ := PromptNewPassword()

		// generate salt and hash using password
		v.Crypto.SetupNewPassword(password)

		// save meta data
		metaData := v.Crypto.ToMeta()
		v.FileIO.EnsureVaultDir()
		v.FileIO.WriteMeta(metaData)
		fmt.Println("Password created successfully.")
		time.Sleep(1 * time.Second)
		ClearScreen()
	}
}
