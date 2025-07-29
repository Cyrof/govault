package cli

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func LoadBanner() string {
	path := filepath.Join("assets", "banner.txt")

	data, err := os.ReadFile(path)
	if err != nil {
		return "GoVault - Secure CLI Tool\n"
	}

	return string(data)
}
