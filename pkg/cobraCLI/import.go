package cobraCLI

import (
	"github.com/Cyrof/govault/internal/backup"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	in    string
	keyIn string
)

var importCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"im"},
	Short:   "Import an encrypted vault archive and ites associated key file",
	Long:    `The 'import' command securely restores your vault and metadata from an encrypted ZIP archive, by decrypting it with a password-protected key file that contains the AES key used during export, ensuring both security and usability while maintaining data integrity.`,
	Example: `
	govault import --in vault_export.zip --key-in key.enc
	govault im --in vault_export2.zip --key-in key2.enc
	`,
	Run: func(cmd *cobra.Command, args []string) {
		password, _ := cli.PromptPassword()
		backup.Import(password, in, keyIn, v)
	},
}

func init() {
	importCmd.Flags().StringVar(&in, "in", "", "Input path for encrypted archive")
	importCmd.Flags().StringVar(&keyIn, "key-in", "", "Input path for encrypted key file")

	if err := importCmd.MarkFlagRequired("in"); err != nil {
		cli.Error("Failed to mark flag as required\n\n")
		logger.Logger.Panicw("Failed to mark flag as required", "flag", "in", "error", err)
	}

	if err := importCmd.MarkFlagRequired("key-in"); err != nil {
		cli.Error("Failed to mark flag as required\n\n")
		logger.Logger.Panicw("Failed to mark flag as required", "flag", "key-in", "error", err)
	}

	rootCmd.AddCommand(importCmd)
}
