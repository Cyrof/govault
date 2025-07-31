package cli

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/Cyrof/govault/internal/logger"
	"github.com/fatih/color"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			logger.Logger.Debugw("Failed to clear screen", "error", err)
		}
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			logger.Logger.Debugw("Failed to clear screen", "error", err)
		}
	}
}

func LoadBanner() string {
	path := filepath.Join("assets", "banner.txt")

	data, err := os.ReadFile(path)
	if err != nil {
		return color.New(color.FgCyan, color.Bold).Sprintf("GoVault - Secure CLI Tool\n")
	}

	return color.New(color.FgCyan, color.Bold).Sprint(string(data))
}
