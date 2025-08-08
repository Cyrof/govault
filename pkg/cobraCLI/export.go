package cobraCLI

import (
	"github.com/Cyrof/govault/internal/backup"
	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	out    string
	keyOut string
)

var exportCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"ep"},
	Short:   "Export you encrypted vault and metadata securely",
	Long: `The 'export' command securely exports your vault data and metadata by encrypting them into an archive file (ZIP format).
	It also generates a seperate key file (key.enc) that contains the encrypted AES key used for securing the archive.
	This key file is protected using your master password, ensuring only you can decrypt it later for import.`,
	Example: `
	govault export --out vault-backup.zip --key-out key.enc
	govault export	# Exports to default current directory 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		salt, hash, err := v.FileIO.ReadMeta()
		if err != nil {
			cli.Error("Failed to read metadata.\nExiting...\n\n")
			logger.Logger.Fatalw("Failed to load meta", "path", v.FileIO.MetaPath, "error", err)
		}

		password, _ := cli.PromptPassword()
		_, ok := crypto.VerifyHash(password, salt, hash)
		if !ok {
			cli.Error("Incorrect master password.\nAborting export.\n\n")
			logger.Logger.Warn("Export aborted due to incorrect password")
			return
		}

		backup.Export(password, v, out, keyOut)
	},
}

func init() {
	exportCmd.Flags().StringVar(&out, "out", "", "Output path for encrypted ZIP archive")
	exportCmd.Flags().StringVar(&keyOut, "key-out", "", "Output path for encrypted AES key file")

	rootCmd.AddCommand(exportCmd)

}
