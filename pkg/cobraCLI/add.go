package cobraCLI

import (
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cli"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new secret",
	Long:    `The 'add' command allows you to securely store a key-value pair in your local encrypted vault.`,
	Example: `	
	govault add --key github --value myGitHubToken123
	govault a -k email -v mypassword
	`,
	Run: func(cmd *cobra.Command, args []string) {
		v.AddSecret(key, value)
		cli.Success("Secret added.")
		logger.Logger.Infow("Secret added", "key", key)
		if err := v.Save(); err != nil {
			cli.Error("Error saving vault: %v\n", err)
			logger.Logger.Errorw("Error saving vault", "error", err)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&key, "key", "k", "", "The name/identifier for the service")
	addCmd.Flags().StringVarP(&value, "value", "v", "", "The value to store")
	if err := addCmd.MarkFlagRequired("key"); err != nil {
		cli.Error("Failed to mark flag as required\n\n")
		logger.Logger.Panicw("Failed to mark flag as required", "flag", "key", "error", err)
	}
	if err := addCmd.MarkFlagRequired("value"); err != nil {
		cli.Error("Failed to mark flag as required\n\n")
		logger.Logger.Panicw("Failed to mark flag as required", "flag", "key", "error", err)
	}

	rootCmd.AddCommand(addCmd)
}
