package cobraCLI

import (
	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/db"
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
		"import":     true,
	}

	v *vault.Vault

	rootCmd = &cobra.Command{
		Use:   "govault",
		Short: "A secure local password vault CLI tool",
		Long:  cli.LoadBanner(),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// initialise dependencies
			crypto := crypto.NewCrypto()
			io := fileIO.NewFileIO()

			// open / create sqlite db
			dbConn, err := db.Open(io.DBPath)
			if err != nil {
				logger.Logger.Fatalw("Failed to open database", "path", io.DBPath, "error", err)
				cli.Error("Failed to open database: %v", err)
			}

			// setup schema from embedded sql files
			if err := db.SetupDatabase(dbConn); err != nil {
				logger.Logger.Fatalw("Database setup failed", "error", err)
				cli.Error("Database setup failed: %v", err)
			}

			v = vault.NewVault(io, crypto, dbConn)

			// skip if its just the root command
			if cmd.Parent() == nil || skipSetupCommands[cmd.Name()] {
				return
			}

			// skip if command is not runnable
			if !cmd.Runnable() {
				return
			}

			// run login/setup
			cli.Setup(v)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if v != nil && v.DB != nil {
				if err := v.DB.Close(); err != nil {
					logger.Logger.Warnw("Failed to close DB", "error", err)
					cli.Warn("Failed to close DB: %v", err)
				}
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Logger.Errorw("Command execution failed", "error", err)

		cli.Error("%v\n", err)
	}
}
