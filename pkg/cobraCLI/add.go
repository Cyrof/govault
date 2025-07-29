package cobraCLI

import (
	"fmt"

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
