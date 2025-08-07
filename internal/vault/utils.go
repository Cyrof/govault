package vault

import "errors"

// function to check if key exist
func (v *Vault) CheckKey(key string) bool {
	if _, exist := v.Secrets[key]; exist {
		return true
	}
	return false
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
