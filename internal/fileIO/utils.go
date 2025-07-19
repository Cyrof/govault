package fileIO

import (
	"fmt"
	"os"
)

func (f *FileIO) PrintPaths() {
	fmt.Println("Vault Directory: ", f.VaultDir)
	fmt.Println("Meta File: ", f.MetaPath)
	fmt.Println("Vault Path: ", f.VaultPath)
}

func (f *FileIO) CheckMetaFile() bool {
	_, err := os.Stat(f.MetaPath)
	return !os.IsNotExist(err)
}

func (f *FileIO) EnsureVaultDir() error {
	return os.MkdirAll(f.VaultDir, 0700)
}

func (f *FileIO) PurgeVault() error {
	return os.RemoveAll(f.VaultDir)
}
