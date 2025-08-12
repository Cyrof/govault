package cobraCLI

import (
	"fmt"

	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cli"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete a secret from the vault",
	Long:    "Removes a stored secret from the vault using its key. This action is irreversible and the deleted secret cannot be recovered.",
	Example: `
	govault delete --key github
	govault d -k email
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if !cli.PromptDel() {
			fmt.Println("Deletion aborted by user.")
			return
		}

		if err := v.DeleteSecret(key); err != nil {
			cli.Error("Failed to delete secret: %v\n", err)
			logger.Logger.Errorw("Failed to delete secret", "key", key, "error", err)
			return
		}

		cli.Success("Secret '%s' deleted successfully.\n", key)
		logger.Logger.Infow("Secret deleted", "key", key)
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&key, "key", "k", "", "The name/identifier of the service")
	if err := deleteCmd.MarkFlagRequired("key"); err != nil {
		cli.Error("Failed to mark required flag: %v\n", err)
		logger.Logger.Panicw("Failed to mark flag as required", "flag", "key", "error", err)
	}

	rootCmd.AddCommand(deleteCmd)
}
