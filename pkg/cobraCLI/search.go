package cobraCLI

import (
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cli"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Fuzzy search stored keys",
	Long: `Search through stored keys using fuzzy logic.
	Use this when you are unsure of the exact key name.`,
	Example: `
	govault search git
	govault s gmail
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cli.Error("Please provide a search query.\n")
			logger.Logger.Error("No search query provided")
		}

		query := args[0]
		if err := v.FuzzyFind(query); err != nil {
			cli.Error("Error searching for %s: %v", query, err)
			logger.Logger.Errorw("Error searching for query", "query", query, "error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
