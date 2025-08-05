package cli

import (
	"errors"
	"fmt"
	"strings"
)

func PromptEdit(key string) (string, error) {
	for {
		Warn("Do you want to edit this secret? (yes/no): ")
		confirm, _ := Reader.ReadString('\n')

		confirm = strings.TrimSpace(confirm)

		if confirm == "" {
			Warn("Cannot be empty.\nPlease try again.\n\n")
			continue
		}

		if confirm == "yes" {
			return promptPassword(key), nil
		} else {
			return "", errors.New("Edit cancelled by user")
		}
	}
}

func promptPassword(key string) string {
	for {
		fmt.Printf("Enter the new password for %s: ", key)
		pass, _ := Reader.ReadString('\n')

		if pass == "" {
			Warn("Cannot be empty.\nPlease try again.\n\n")
			continue
		}

		return pass
	}
}
