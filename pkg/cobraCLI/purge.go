package cobraCLI

import (
	"fmt"
	"os"

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
			if err := v.FileIO.PurgeVault(); err != nil {
				fmt.Println("Failed to purge vault:", err)
			}
			fmt.Println("All vault data has been successfully purged. The system has been reset.")
		} else {
			fmt.Println("Purge operation cancelled. No changes were made.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}
