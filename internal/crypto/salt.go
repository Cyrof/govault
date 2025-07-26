package crypto

import (
	"crypto/rand"
	"fmt"
)

func GenerateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate salt: %w", err)
	}
	return salt, nil
}
