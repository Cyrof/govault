// file to handle all behaviors of vault
package vault

import (
	"errors"
	"fmt"

	"github.com/Cyrof/govault/internal/model"
)

// function to add secret
func (v *Vault) AddSecret(key, value string) {
	s := model.Secret{
		Key:   key,
		Value: value,
	}
	v.Secrets[key] = s
}

// function to get secret
func (v *Vault) GetSecret(key string) (model.Secret, bool) {
	val, ok := v.Secrets[key]
	return val, ok
}

// function to display all keys
func (v *Vault) DisplayKeys() {
	fmt.Println("Stored keys in the vault:")
	for key := range v.Secrets {
		fmt.Println(" -", key)
	}
}

// function to delete
func (v *Vault) DeleteSecret(key string) error {
	if !v.CheckKey(key) {
		return errors.New("key not found in vault")
	}
	delete(v.Secrets, key)
	return nil
}
