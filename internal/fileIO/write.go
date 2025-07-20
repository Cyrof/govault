package fileIO

import "os"

// function to write to meta file
func (f *FileIO) WriteMeta(data []byte) error {
	return os.WriteFile(f.MetaPath, data, 0600)
}

// function to write to vault file
func (f *FileIO) WriteSecret(data []byte) error {
	return os.WriteFile(f.VaultPath, data, 0600)
}
