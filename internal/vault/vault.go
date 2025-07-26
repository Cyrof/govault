package vault

import (
	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/fileIO"
)

type Vault struct {
	Secrets map[string]string
	FileIO  *fileIO.FileIO
	Crypto  *crypto.Crypto
}

// constructor function
func NewVault(fileIO *fileIO.FileIO, crypto *crypto.Crypto) *Vault {
	return &Vault{
		Secrets: make(map[string]string),
		FileIO:  fileIO,
		Crypto:  crypto,
	}
}
