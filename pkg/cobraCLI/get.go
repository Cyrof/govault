package cobraCLI

import (
	"fmt"

	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cli"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get value of secret",
	Long:    `The 'get' command allows you to securely retrieve a secret value from the encrypted vault by specifying the associated key.`,
	Example: `
	govault get --key github
	govault g -k email
	`,
	Run: func(cmd *cobra.Command, args []string) {
		secret, ok := v.GetSecret(key)
		if ok {
			fmt.Println("Value:", secret.Value)
			logger.Logger.Debugw("Found value", "key", key)
		} else {
			cli.Error("Key not found.")
			logger.Logger.Error("Key not found.")
		}
	},
}

func init() {
	getCmd.Flags().StringVarP(&key, "key", "k", "", "The name/identifier of the sevice")
	if err := getCmd.MarkFlagRequired("key"); err != nil {
		cli.Error("Failed to mark required flag: %v\n", err)
		logger.Logger.Panicw("Failed to mark flag as required", "flag", "key", "error", err)
	}

	rootCmd.AddCommand(getCmd)
}
