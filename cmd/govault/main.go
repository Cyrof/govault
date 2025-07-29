package main

import (
	// "github.com/Cyrof/govault/internal/crypto"
	// "github.com/Cyrof/govault/internal/fileIO"
	// "github.com/Cyrof/govault/internal/vault"
	// cli.Execute(store)
	"github.com/Cyrof/govault/pkg/cobraCLI"
)

func main() {
	// crypto := crypto.NewCrypto()
	// io := fileIO.NewFileIO()
	// store := vault.NewVault(io, crypto)
	// cli.Execute(store)

	cobraCLI.Execute()
}
