package fileIO

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Cyrof/govault/internal/model"
)

// function to read meta file
func (f *FileIO) ReadMeta() ([]byte, []byte, error) {
	data, err := os.ReadFile(f.MetaPath)
	if err != nil {
		return nil, nil, err
	}

	var meta model.Meta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, nil, err
	}

	saltBytes, err := base64.StdEncoding.DecodeString(meta.Salt)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode salt: %w", err)
	}

	hashBytes, err := base64.StdEncoding.DecodeString(meta.Hash)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode hash: %w", err)
	}

	return saltBytes, hashBytes, nil
}

// function to read vault file
func (f *FileIO) ReadVault() ([]byte, error) {
	return os.ReadFile(f.VaultPath)
}
