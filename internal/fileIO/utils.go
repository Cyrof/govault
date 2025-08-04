package fileIO

import (
	"fmt"
	"os"
)

// function to print all constant path
func (f *FileIO) PrintPaths() {
	fmt.Println("Vault Directory: ", f.VaultDir)
	fmt.Println("Meta File: ", f.MetaPath)
	fmt.Println("Vault Path: ", f.VaultPath)
}

// function to check if meta file exists
func (f *FileIO) CheckMetaFile() bool {
	_, err := os.Stat(f.MetaPath)
	return !os.IsNotExist(err)
}

// function to make sure dir exists
func (f *FileIO) EnsureVaultDir() error {
	return os.MkdirAll(f.VaultDir, 0700)
}

// function to purge all files and directory
func (f *FileIO) PurgeVault() error {
	return os.RemoveAll(f.VaultDir)
}

// function to check if vault exists
func (f *FileIO) IsEmpty() (bool, error) {
	_, err := os.Stat(f.VaultPath)
	if err == nil {
		// file exists
		return false, nil
	}
	if os.IsNotExist(err) {
		// file does not exist
		return true, nil
	}
	// other errors (e.g., permission)
	return false, err
}
