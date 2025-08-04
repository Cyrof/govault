package cli

import (
	"bufio"
	"os"
	"strings"
)

func PromptDel() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		Error("Are you sure you want to delete this secret?\n")
		Warn("This action is irreversible and the secret cannot be recoved. (yes/no): ")
		confirm, _ := reader.ReadString('\n')

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
