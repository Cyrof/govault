package cli

import (
	"flag"
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

func Execute(v *vault.Vault) {
	// main commands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	purgeCmd := flag.NewFlagSet("purge", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// sub args for add arg
	addKey := addCmd.String("key", "", "Key to store")
	addVal := addCmd.String("value", "", "Value to store")

	// sub args for get arg
	getKey := getCmd.String("key", "", "Key to retrieve")

	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		Setup(v)
		addCmd.Parse(os.Args[2:])
		// add secret to vault
		v.AddSecret(*addKey, *addVal)
		fmt.Println("Secret added.")
		v.Save()
	case "get":
		Setup(v)
		getCmd.Parse(os.Args[2:])
		val, ok := v.GetSecret(*getKey)
		if ok {
			fmt.Println("Value:", val.Value)
		} else {
			fmt.Println("Error retrieving secret")
		}
	case "purge":
		purgeCmd.Parse(os.Args[2:])
		confirm := PromptPurge()
		if confirm {
			v.FileIO.PurgeVault()
			fmt.Println("All vault data has been successfully purged. The system has been reset.")
		} else {
			fmt.Println("Purge operation cancelled. No changes were made.")
			os.Exit(1)
		}
	case "list":
		Setup(v)
		listCmd.Parse(os.Args[2:])
		v.DisplayKeys()
		os.Exit(1)
	default:
		PrintUsage()
		os.Exit(1)
	}
}

func PrintUsage() {
	fmt.Println("Usage:")
	// add key description
	fmt.Println("	govault add -key <key> -value <value>")
	fmt.Println("		Add a new secret to the vault.")

	// get key description
	fmt.Println("	\n\tgovault get -key <key>")
	fmt.Println("		Retrieve a stored secreted from the vault by its key.")

	// purge vault description
	fmt.Println("	\n\tgovault purge")
	fmt.Println("		Permanently delete all stored vault data and reset the application.")

	// example description
	fmt.Println("\nExample:")
	fmt.Println("	govault add -key email -value myemailpassword")
	fmt.Println("	govault get -key email")
	fmt.Println("	govault purge")
}
