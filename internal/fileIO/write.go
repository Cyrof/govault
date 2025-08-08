package fileIO

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Cyrof/govault/internal/model"
)

// function to write to meta file
func (f *FileIO) WriteMeta(meta model.Meta) error {
	data, err := json.MarshalIndent(meta, "", "	")
	if err != nil {
		return fmt.Errorf("failed to marshal meta: %w", err)
	}
	return os.WriteFile(f.MetaPath, data, 0600)
}

// function to write to vault file
func (f *FileIO) WriteSecret(data []byte) error {
	return os.WriteFile(f.VaultPath, data, 0600)
}

func WriteEncryptedZip(outPath string, files map[string][]byte) (retErr error) {
	if outPath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		outPath = filepath.Join(cwd, "vault_export.zip")
	}

	if err := os.MkdirAll(filepath.Dir(outPath), os.ModePerm); err != nil {
		return err
	}

	// prevent overwrite
	if _, err := os.Stat(outPath); err == nil {
		return errors.New("file already exists as " + outPath)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("unable to check file: %w", err)
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := outFile.Close(); cerr != nil && retErr == nil {
			retErr = cerr
		}
	}()

	zipWriter := zip.NewWriter(outFile)
	defer func() {
		if cerr := zipWriter.Close(); cerr != nil && retErr == nil {
			retErr = cerr
		}
	}()

	for name, data := range files {
		writer, err := zipWriter.Create(name)
		if err != nil {
			return err
		}
		if _, err := writer.Write(data); err != nil {
			return err
		}
	}

	return nil
}

func WriteKeyFile(data []byte, outPath string) error {
	if outPath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		outPath = filepath.Join(cwd, "key.enc")
	}

	if err := os.WriteFile(outPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write key to file: %w", err)
	}

	return nil
}
