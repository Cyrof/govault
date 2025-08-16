package fileIO

import (
	"os"
	"path/filepath"
	"runtime"
)

func DataHome() string {
	home, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "windows":
		if appdata := os.Getenv("APPDATA"); appdata != "" {
			return filepath.Join(appdata, "GoVault")
		}
		return filepath.Join(home, "AppData", "Roaming", "GoVault")

	case "darwin":
		// ~/library/Application Support/GoVault
		return filepath.Join(home, "Library", "Application Support", "GoVault")

	default:
		// linux / unix
		if xdg := os.Getenv("XDG_DATA_HOME"); xdg != "" {
			return filepath.Join(xdg, "govault")
		}
		// ~/.local/share/govault
		return filepath.Join(home, ".local", "share", "govault")
	}
}
