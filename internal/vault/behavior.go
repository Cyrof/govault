// file to handle all behaviors of vault
package vault

import (
	"errors"
	"fmt"

	"github.com/Cyrof/govault/internal/model"
	"github.com/lithammer/fuzzysearch/fuzzy"
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

// function to updated password
func (v *Vault) EditPassword(key string, newPass string) error {
	if !v.CheckKey(key) {
		return errors.New("key not found in vault")
	}
	v.Secrets[key] = model.Secret{
		Key:   key,
		Value: newPass,
	}
	return nil
}

// function to use fuzzy search to find key
func (v *Vault) FuzzyFind(query string) error {
	keys, err := v.GetKeys()
	if err != nil {
		return err
	}
	matches := fuzzy.Find(query, keys)

	if len(matches) == 0 {
		return errors.New("no matches found")
	}

	fmt.Printf("Matches for %s:\n", query)
	for _, match := range matches {
		fmt.Printf("- %s\n", match)
	}
	return nil
}
