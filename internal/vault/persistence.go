package vault

import (
	"encoding/json"
)

// function to save secrets
func (v *Vault) Save() error {
	// convert secrets to json
	data, err := json.Marshal(v.Secrets)
	if err != nil {
		return err
	}

	return v.FileIO.WriteSecret(data)
}
