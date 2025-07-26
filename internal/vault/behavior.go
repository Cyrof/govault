// file to handle all behaviors of vault
package vault

// function to add secret
func (v *Vault) AddSecret(key, value string) {
	v.Secrets[key] = value
}

// function to get secret
func (v *Vault) GetSecret(key string) (string, bool) {
	val, ok := v.Secrets[key]
	return val, ok
}
