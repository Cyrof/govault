package fileIO

import (
	"os"
	"path/filepath"
)

type FileIO struct {
	VaultDir  string
	MetaPath  string
	VaultPath string
}

// constructor for fileIO
func NewFileIO() *FileIO {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Failed to get home dir")
	}
	vaultDir := filepath.Join(home, ".localvault")

	return &FileIO{
		VaultDir:  vaultDir,
		MetaPath:  filepath.Join(vaultDir, "meta.json"),
		VaultPath: filepath.Join(vaultDir, "vault.enc"),
	}
}
