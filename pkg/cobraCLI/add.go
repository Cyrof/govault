package cobraCLI

import (
	"github.com/Cyrof/govault/internal/generator"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cli"
	"github.com/spf13/cobra"
)

var gen bool

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new secret or generate one using default settings",
	Long: `The 'add' command allows you to securely store a key-value pair in your local encrypted vault.
	You must specify either a manual value using --value, or use --gen to generate a random password.
	`,
	Example: `	
	govault add --key github --value myGitHubToken123
	govault a -k email -v mypassword
	govault add --key github --gen
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if value != "" && gen {
			cli.Error("You cannot use both --value and --gen. Please choose one.")
			return
		}

		var secretValue string

		if gen {
			opts := generator.DefaultOptions
			pass, err := generator.GeneratePassword(opts)
			if err != nil {
				cli.Error("Failed to generate password: %v\n", err)
				logger.Logger.Errorw("Password generation failed", "error", err)
				return
			}
			secretValue = pass
			cli.Success("Generated password: %s\n", pass)
			logger.Logger.Infow("Password generated successfully", "key", key)
		} else if value != "" {
			secretValue = value
			logger.Logger.Infow("Secret added manually", "key", key)
		} else {
			cli.Error("You must provide either --value or --gen.\n")
			return
		}

		err := v.AddSecret(key, secretValue)
		if err != nil {
			logger.Logger.Errorw("Error adding secret to vault", "error", err)
			cli.Error("Error adding secret to vault: %v", err)
			return
		}

		cli.Success("Secret added.\n")

	},
}

func init() {
	addCmd.Flags().StringVarP(&key, "key", "k", "", "The name/identifier for the service")
	addCmd.Flags().StringVarP(&value, "value", "v", "", "The value to store")
	addCmd.Flags().BoolVar(&gen, "gen", false, "Generate a random password using default settings")
	if err := addCmd.MarkFlagRequired("key"); err != nil {
		cli.Error("Failed to mark flag as required\n\n")
		logger.Logger.Panicw("Failed to mark flag as required", "flag", "key", "error", err)
	}

	rootCmd.AddCommand(addCmd)
}
