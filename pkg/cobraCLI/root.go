package cobraCLI

import (
	"fmt"
	"os"

	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/vault"
	"github.com/Cyrof/govault/pkg/cli"

	"github.com/spf13/cobra"
)

var (
	skipSetupCommands = map[string]bool{
		"help":       true,
		"completion": true,
	}

	v *vault.Vault

	rootCmd = &cobra.Command{
		Use:   "govault",
		Short: "A secure local password vault CLI tool",
		Long:  `GoVault is a lightweight CLI tool for storing secrets using AES encryption`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// skip if its just the root command
			if cmd.Parent() == nil || skipSetupCommands[cmd.Name()] {
				return
			}

			// skip if command is not runnable
			if !cmd.Runnable() {
				return
			}

			// initialise dependencies
			crypto := crypto.NewCrypto()
			io := fileIO.NewFileIO()
			v = vault.NewVault(io, crypto)

			// run login/setup
			cli.Setup(v)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
