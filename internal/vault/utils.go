package vault

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Cyrof/govault/internal/db"
)

// function to check if key exist
func (v *Vault) CheckKey(key string) (bool, error) {
	if v.DB == nil {
		return false, fmt.Errorf("database not initialised")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return db.KeyExists(ctx, v.DB, key)
}

// function to return all keys
func (v *Vault) GetKeys() ([]string, error) {
	keys := make([]string, 0, len(v.Secrets))
	for key := range v.Secrets {
		keys = append(keys, key)
	}

	if len(keys) > 0 {
		return keys, nil
	} else {
		return nil, errors.New("no keys in vault")
	}
}
