package fileIO

import (
	"os"
)

func (f *FileIO) ReadMeta() ([]byte, error) {
	return os.ReadFile(f.MetaPath)
}
