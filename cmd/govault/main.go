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
		if err := logger.Logger.Sync(); err != nil && err.Error() != "sync /dev/stderr: invalid argument" { // this is ignore since it is a known issue and does not affect functionality
			logger.Logger.Warnw("Logger sync failed", "error", err)
		}
	}()

	cobraCLI.Execute()
}
