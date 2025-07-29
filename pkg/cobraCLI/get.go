package cobraCLI

import (
	"fmt"

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
		} else {
			fmt.Println("Key not found.")
		}
	},
}

func init() {
	getCmd.Flags().StringVarP(&key, "key", "k", "", "The name/identifier of the sevice")
	getCmd.MarkFlagRequired("key")

	rootCmd.AddCommand(getCmd)
}
