package fileIO

import (
	"os"
	"path/filepath"
)

type FileIO struct {
	VaultDir  string
	MetaPath  string
	VaultPath string
	LogPath   string
	DBPath    string
}

// constructor for fileIO
func NewFileIO() *FileIO {
	base := DataHome()

	if err := os.MkdirAll(filepath.Join(base, "logs"), 0o700); err != nil {
		panic("error creating logs")
	}

	return &FileIO{
		VaultDir:  base,
		MetaPath:  filepath.Join(base, "meta.json"),
		VaultPath: filepath.Join(base, "vault.enc"),
		LogPath:   filepath.Join(base, "logs", "govault.log"),
		DBPath:    filepath.Join(base, "vault.db"),
	}
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	panic("Failed to get home dir")
	// }
	// vaultDir := filepath.Join(home, ".localvault")
	//
	// return &FileIO{
	// 	VaultDir:  vaultDir,
	// 	MetaPath:  filepath.Join(vaultDir, "meta.json"),
	// 	VaultPath: filepath.Join(vaultDir, "vault.enc"),
	// 	LogPath:   filepath.Join(vaultDir, "logs", "govault.log"),
	// }
}
