package fileIO

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Cyrof/govault/internal/model"
)

// function to write to meta file
func (f *FileIO) WriteMeta(meta model.Meta) error {
	data, err := json.MarshalIndent(meta, "", "	")
	if err != nil {
		return fmt.Errorf("failed to marshal meta: %w", err)
	}
	fmt.Println(string(data))
	return os.WriteFile(f.MetaPath, data, 0600)
}

// function to write to vault file
func (f *FileIO) WriteSecret(data []byte) error {
	return os.WriteFile(f.VaultPath, data, 0600)
}
