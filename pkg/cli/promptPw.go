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
			fmt.Println("Password cannot be empty. Please try again.")
			continue
		}

		// check if pass matches
		if pass1 != pass2 {
			fmt.Println("Password do not match. Please try again.")
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

		if pass1 == "" {
			fmt.Println("Password cannot be empty. Please try again.")
			continue
		}

		return pass1, nil
	}
}
