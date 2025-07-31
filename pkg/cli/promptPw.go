package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// prompt for new password creation
func PromptNewPassword() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a new master password: ")
		pass1, _ := reader.ReadString('\n')

		fmt.Print("Confirm your master password: ")
		pass2, _ := reader.ReadString('\n')

		// format inputs
		pass1 = strings.TrimSpace(pass1)
		pass2 = strings.TrimSpace(pass2)

		// check if pass 1 or 2 is empty
		if pass1 == "" || pass2 == "" {
			Warn("Password cannot be empty.\nPlease try again.\n\n")
		}

		// check if pass matches
		if pass1 != pass2 {
			Error("Password do not match.\nPlease try again.\n\n")
			continue
		}

		return pass1, nil
	}
}

func PromptPassword() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your master password: ")
		pass1, _ := reader.ReadString('\n')
		pass1 = strings.TrimSpace(pass1)

		if pass1 == "" {
			Warn("Password cannot be empty.\nPlease try again.\n\n")
			continue
		}

		return pass1, nil
	}
}
