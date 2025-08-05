package cobraCLI

import (
	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/internal/vault"
	"github.com/Cyrof/govault/pkg/cli"

	"github.com/spf13/cobra"
)

var (
	skipSetupCommands = map[string]bool{
		"help":       true,
		"purge":      true,
		"completion": true,
		"generate":   true,
	}

	v *vault.Vault

	rootCmd = &cobra.Command{
		Use:   "govault",
		Short: "A secure local password vault CLI tool",
		Long:  cli.LoadBanner(),
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
		logger.Logger.Errorw("Command execution failed", "error", err)

		cli.Error("%v\n", err)
	}
}
