package vault

import (
	"encoding/json"
)

// function to save secrets
func (v *Vault) Save() error {
	// convert secrets to json
	data, err := json.Marshal(v.secrets)
	if err != nil {
		return err
	}

	// encryption should be here

	return v.fileIO.WriteSecret(data)
}
