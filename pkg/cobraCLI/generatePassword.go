package cobraCLI

import (
	"github.com/Cyrof/govault/internal/generator"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cli"
	"github.com/spf13/cobra"
)

var (
	length      int
	noLowercase bool
	noUppercase bool
	noNumbers   bool
	withSymbols bool
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate a secure random password",
	Long:    "Generate a secure, random password with configurable length and character sets.",
	Example: `
	govault generate 
	govault generate --length 24 --symbols
	govault generate -l 24 -s
	govault generate --no-lowercase --no-uppercase
	`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := generator.DefaultOptions

		opts.Length = length
		opts.UseLowercase = !noLowercase
		opts.UseUppercase = !noUppercase
		opts.UseNumbers = !noNumbers
		opts.UseSymbols = withSymbols

		password, err := generator.GeneratePassword(opts)
		if err != nil {
			cli.Error("Error generating password: %v\n", err)
			logger.Logger.Errorw("Error generating password", "error", err)
		}

		cli.Success("Generated password: %v", password)
		logger.Logger.Infow("Password generated successfully")
	},
}

func init() {
	generateCmd.Flags().IntVarP(&length, "length", "l", generator.DefaultOptions.Length, "Length of the password")
	generateCmd.Flags().BoolVar(&noLowercase, "no-lowercase", !generator.DefaultOptions.UseLowercase, "Exclude lowercase letters")
	generateCmd.Flags().BoolVar(&noUppercase, "no-uppercase", !generator.DefaultOptions.UseUppercase, "Exclude uppercase letters")
	generateCmd.Flags().BoolVar(&noNumbers, "no-numbers", !generator.DefaultOptions.UseNumbers, "Exclude numeric characters")
	generateCmd.Flags().BoolVarP(&withSymbols, "symbols", "s", generator.DefaultOptions.UseSymbols, "Include symbols")

	rootCmd.AddCommand(generateCmd)
}
