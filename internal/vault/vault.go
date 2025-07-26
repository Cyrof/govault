package vault

import (
	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/model"
)

type Vault struct {
	Secrets map[string]model.Secret
	FileIO  *fileIO.FileIO
	Crypto  *crypto.Crypto
}

// constructor function
func NewVault(fileIO *fileIO.FileIO, crypto *crypto.Crypto) *Vault {
	return &Vault{
		Secrets: make(map[string]model.Secret),
		FileIO:  fileIO,
		Crypto:  crypto,
	}
}
