package cli

import (
	"bufio"
	"os"
	"strings"
)

func PromptPurge() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		Error("WARNING: This action will permanently delete all stored secrets and reset the vault.\n")
		Error("This operation cannot be undone.\n\n")
		Warn("Are you sure you want to continue? (yes/no): ")
		confirm, _ := reader.ReadString('\n')

		// format input
		confirm = strings.TrimSpace(confirm)

		if confirm == "" {
			Warn("Cannot be empty.\nPlease try again.\n\n")
			continue
		}

		if confirm == "yes" {
			return true
		} else {
			return false
		}

	}
}
