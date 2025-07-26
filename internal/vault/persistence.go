package vault

import (
	"encoding/json"
	"fmt"

	"github.com/Cyrof/govault/internal/model"
)

// function to save secrets
func (v *Vault) Save() error {
	// convert secrets to json
	data, err := json.MarshalIndent(v.Secrets, "", " ")
	if err != nil {
		return fmt.Errorf("Failed to marshal secrets: %w", err)
	}

	// encrypt data
	encData, err := v.Crypto.Encrypt(data)
	if err != nil {
		return fmt.Errorf("Failed to encrypt secrets: %w", err)
	}
	fmt.Printf("Encrypted data length: %d bytes\n", len(encData))
	// write to enc file
	if err := v.FileIO.WriteSecret(encData); err != nil {
		return fmt.Errorf("Failed to write encrypted secrets: %w", err)
	}

	return nil
}

// function to load secrets
func (v *Vault) Load() error {
	encData, err := v.FileIO.ReadVault()
	if err != nil {
		return fmt.Errorf("Failed to read vault: %w", err)
	}

	// decrypt data
	data, err := v.Crypto.Decrypt(encData)
	if err != nil {
		return fmt.Errorf("Failed to decrypt secrets: %w", err)
	}

	// unmarshal to json
	var secrets map[string]model.Secret
	if err := json.Unmarshal(data, &secrets); err != nil {
		return fmt.Errorf("Failed to unmarshal secrets: %w", err)
	}

	// load to Secrets
	v.Secrets = secrets

	return nil
}
