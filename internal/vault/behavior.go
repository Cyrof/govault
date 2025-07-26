// file to handle all behaviors of vault
package vault

import (
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
