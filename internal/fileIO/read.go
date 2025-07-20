package fileIO

import (
	"os"
)

// function to read meta file
func (f *FileIO) ReadMeta() ([]byte, error) {
	return os.ReadFile(f.MetaPath)
}

// function to read vault file
func (f *FileIO) ReadVault() ([]byte, error) {
	return os.ReadFile(f.MetaPath)
}
