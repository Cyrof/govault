package cobraCLI

import (
	"fmt"

	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cli"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "Edit a stored secret in the vault",
	Long:    "Edit a existing secret in the vault by specifying its key and updating its value securely.",
	Example: `
	govault edit --key github
	govault e -k email
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// check if key exist
		exists, err := v.CheckKey(key)
		if err != nil {
			cli.Error("Failed to check key: %v\n", err)
			logger.Logger.Errorw("CheckKey failed", "key", key, "error", err)
			return
		}
		if !exists {
			cli.Error("Key '%s' not found in the vault.", key)
			logger.Logger.Warnw("Key not found", "key", key)
			return
		}

		// prompt user for confirmation
		newPass, err := cli.PromptEdit(key)
		if err != nil {
			fmt.Println(err)
		}

		// change existing password logic
		if err := v.EditPassword(key, newPass); err != nil {
			cli.Error("Failed to update password: %v\n", err)
			logger.Logger.Errorw("Failed to update password", "key", key, "error", err)
			return
		}

		cli.Success("'%s' password successfully updated.\n", key)
		logger.Logger.Infow("Secret updated", "key", key)
	},
}

func init() {
	editCmd.Flags().StringVarP(&key, "key", "k", "", "The name/identifier of the service")
	if err := editCmd.MarkFlagRequired("key"); err != nil {
		cli.Error("Failed to mark required flag: %v\n", err)
		logger.Logger.Panicw("Failed to mark flag as required", "flag", "key", "error", err)
	}

	rootCmd.AddCommand(editCmd)
}
