package cobraCLI

import (
	"os"

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
		v.DisplayKeys()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
