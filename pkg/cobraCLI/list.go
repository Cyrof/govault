package cobraCLI

import (
	"github.com/Cyrof/govault/internal/logger"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all stored keys in the vault",
	Long:    `Displays all keys stored in the vault without revealing their values.`,
	Example: `
	govault list
	govault l
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := v.DisplayKeys(); err != nil {
			logger.Logger.Errorw("Error displaying Keys", "error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
