package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptPurge() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("WARNING: This action will permanently delete all stored secrets and reset the vault.")
		fmt.Println("This operation cannot be undone.")
		fmt.Println()
		fmt.Print("Are you sure you want to continue? (yes/no): ")
		confirm, _ := reader.ReadString('\n')

		// format input
		confirm = strings.TrimSpace(confirm)

		if confirm == "" {
			fmt.Println("Cannot be empty. Please try again.")
			continue
		}

		if confirm == "yes" {
			return true
		} else {
			return false
		}

	}
}
