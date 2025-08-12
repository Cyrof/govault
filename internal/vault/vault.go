package vault

import (
	"database/sql"

	"github.com/Cyrof/govault/internal/crypto"
	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/model"
)

type Vault struct {
	Secrets map[string]model.Secret
	FileIO  *fileIO.FileIO
	Crypto  *crypto.Crypto
	DB      *sql.DB
}

// constructor function
func NewVault(fileIO *fileIO.FileIO, crypto *crypto.Crypto, db *sql.DB) *Vault {
	return &Vault{
		Secrets: make(map[string]model.Secret),
		FileIO:  fileIO,
		Crypto:  crypto,
		DB:      db,
	}
}
