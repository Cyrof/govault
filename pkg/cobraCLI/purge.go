package cobraCLI

import (
	"os"

	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cli"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:     "purge",
	Aliases: []string{"p"},
	Short:   "Delete all vault data and reset",
	Long:    `Permanently deltes the vault and metadata files, resetting GoVault`,
	Example: `
	govault purge
	govault p
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if cli.PromptPurge() {
			fileIO := fileIO.NewFileIO()
			if err := fileIO.PurgeVault(); err != nil {
				cli.Error("Failed to purge vault: %v\n", err)
				logger.Logger.Debugw("Failed to purge vault", "error", err)
			}
			cli.Success("All vault data has been successfully purged. The system has been reset.")
			logger.Logger.Info("All vault data purged")
		} else {
			cli.Warn("Purge operation cancelled. No changes were made.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}
