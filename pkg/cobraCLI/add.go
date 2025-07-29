package cobraCLI

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	key   string
	value string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new secret",
	Run: func(cmd *cobra.Command, args []string) {
		v.AddSecret(key, value)
		fmt.Println("Secret added.")
		v.Save()
	},
}

func init() {
	addCmd.Flags().StringVarP(&key, "key", "k", "", "The name/identifier for the service")
	addCmd.Flags().StringVarP(&value, "value", "v", "", "The value to store")
	addCmd.MarkFlagRequired("key")
	addCmd.MarkFlagRequired("value")

	rootCmd.AddCommand(addCmd)
}
