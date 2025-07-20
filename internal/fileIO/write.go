package fileIO

import "os"

func (f *FileIO) WriteMeta(data []byte) error {
	return os.WriteFile(f.MetaPath, data, 0600)
}

func (f *FileIO) WriteSecret(data []byte) error {
	return os.WriteFile(f.VaultPath, data, 0600)
}
