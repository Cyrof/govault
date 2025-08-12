// file to handle all behaviors of vault
package vault

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Cyrof/govault/internal/db"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

// function to add secret
func (v *Vault) AddSecret(key, value string) error {
	ct, err := v.Crypto.Encrypt([]byte(value), nil)
	if err != nil {
		return fmt.Errorf("encrypt secret: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.InsertSecret(ctx, v.DB, key, ct); err != nil {
		if err == db.ErrDuplicateKey {
			return fmt.Errorf("key %q already exists", key)
		}
		return err
	}
	return nil
}

// function to get secret
func (v *Vault) GetSecret(key string) (string, error) {
	if v.DB == nil {
		return "", errors.New("database not initialised")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ct, err := db.GetSecretCT(ctx, v.DB, key)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return "", err
		}
		return "", err
	}

	pt, err := v.Crypto.Decrypt(ct, nil)
	if err != nil {
		return "", err
	}

	return string(pt), nil
}

// function to display all keys
func (v *Vault) DisplayKeys() error {
	if v.DB == nil {
		return fmt.Errorf("database not initialised")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	keys, err := db.ListKeys(ctx, v.DB)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		fmt.Println("No keys stored yet.")
		return nil
	}

	fmt.Println("Stored keys in the vault:")
	for _, k := range keys {
		fmt.Println(" -", k)
	}
	return nil
}

// function to delete
func (v *Vault) DeleteSecret(key string) error {
	if v.DB == nil {
		return fmt.Errorf("database not initialised")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Delete(ctx, v.DB, key); err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return fmt.Errorf("key %q not found", key)
		}
		return err
	}
	return nil
}

// function to updated password
func (v *Vault) EditPassword(key string, newPass string) error {
	if v.DB == nil {
		return fmt.Errorf("database not initialised")
	}

	ct, err := v.Crypto.Encrypt([]byte(newPass), nil)
	if err != nil {
		return fmt.Errorf("encrypt: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// update row
	if err := db.Update(ctx, v.DB, key, ct); err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return fmt.Errorf("key %q not found", key)
		}
		return err
	}
	return nil
}

// function to use fuzzy search to find key
func (v *Vault) FuzzyFind(query string) error {
	if v.DB == nil {
		return fmt.Errorf("database not initialised")
	}

	keys, err := v.GetKeys()
	if err != nil {
		return err
	}

	matches := fuzzy.FindNormalizedFold(query, keys)

	if len(matches) == 0 {
		return errors.New("no matches found")
	}

	fmt.Printf("Matches for %q:\n", query)
	for _, m := range matches {
		fmt.Printf("- %s\n", m)
	}
	return nil
}
