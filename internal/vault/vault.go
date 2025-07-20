package vault

import (
	"github.com/Cyrof/govault/internal/fileIO"
)

type Vault struct {
	secrets map[string]string
	fileIO  *fileIO.FileIO
	// should add crypto mod here
}

// constructor function
func NewVault(fileIO *fileIO.FileIO) *Vault {
	return &Vault{
		secrets: make(map[string]string),
		fileIO:  fileIO,
	}
}
