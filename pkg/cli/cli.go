package cli

import (
	// "flag"
	"fmt"
	"os"
	"time"

	"github.com/Cyrof/govault/internal/vault"
)

func Setup(v *vault.Vault) {
	if v.FileIO.CheckMetaFile() {
		// setup crypto salt and hash
		metaData, err := v.FileIO.ReadMeta()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// call prompt pw cli to get master password from user
		password, _ := PromptPassword()

		// setup crypto salt and hash
		v.Crypto.SetupFromMeta(password, []byte(metaData.Salt), []byte(metaData.Hash))
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

// /* func Execute(v *vault.Vault) {
// 	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
// 	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
//
// 	// sub args for add arg
// 	addKey := addCmd.String("key", "", "Key to store")
// 	addVal := addCmd.String("value", "", "Value to store")
//
// 	// sub args for get arg
// 	getKey := getCmd.String("key", "", "Key to retrieve")
//
// 	if len(os.Args) < 2 {
// 		PrintUsage()
// 		os.Exit(1)
// 	}
//
// 	switch os.Args[1] {
// 	case "add":
// 		addCmd.Parse(os.Args[2:])
// 		v.AddSecret(*addKey, *addVal)
// 		fmt.Println("Secret added.")
// 	}
// } */

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println(" govault add -key <key> -value <value>")
	fmt.Println(" govault get -key <key>")
}
