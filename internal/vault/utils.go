package vault

// function to check if key exist
func (v *Vault) CheckKey(key string) bool {
	if _, exist := v.Secrets[key]; exist {
		return true
	}
	return false
}
