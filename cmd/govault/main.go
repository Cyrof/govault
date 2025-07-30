package main

import (
	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cobraCLI"
)

func main() {
	f := fileIO.NewFileIO()
	logger.InitLogger(f)

	defer func() {
		if err := logger.Logger.Sync(); err != nil {
			logger.Logger.Warnw("Logger sync failed", "error", err)
		}
	}()

	cobraCLI.Execute()
}
