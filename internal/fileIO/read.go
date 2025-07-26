package fileIO

import (
	"encoding/json"
	"os"

	"github.com/Cyrof/govault/internal/model"
)

// function to read meta file
func (f *FileIO) ReadMeta() (*model.Meta, error) {
	data, err := os.ReadFile(f.MetaPath)
	if err != nil {
		return nil, err
	}

	var meta model.Meta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

// function to read vault file
func (f *FileIO) ReadVault() ([]byte, error) {
	return os.ReadFile(f.VaultPath)
}
