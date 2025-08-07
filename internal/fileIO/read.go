package fileIO

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
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

// read and return raw data of both file
func (f *FileIO) ReadAll() ([]byte, []byte, error) {
	vaultData, err := os.ReadFile(f.VaultPath)
	if err != nil {
		return nil, nil, err
	}
	metaData, err := os.ReadFile(f.MetaPath)
	if err != nil {
		return nil, nil, err
	}
	return vaultData, metaData, nil
}

// function to unzip and read encrypted folder
func ReadEncryptedZip(zipPath string) (map[string][]byte, error) {
	result := make(map[string][]byte)

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open zip file: %w", err)
	}

	var retErr error
	defer func() {
		if err := r.Close(); err != nil {
			retErr = err
		}
	}()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			retErr = fmt.Errorf("failed to open zip entry %s: %w", f.Name, err)
			break
		}

		var buf bytes.Buffer
		if _, err := io.Copy(&buf, rc); err != nil {
			rc.Close()
			retErr = fmt.Errorf("failed to read zip entry %s: %w", f.Name, err)
			break
		}

		if err := rc.Close(); err != nil {
			retErr = fmt.Errorf("failed to close zip entry %s: %w", f.Name, err)
			break
		}

		result[f.Name] = buf.Bytes()
	}

	if retErr != nil {
		return nil, retErr
	}
	return result, nil
}
