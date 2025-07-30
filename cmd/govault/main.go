package main

import (
	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/Cyrof/govault/internal/logger"
	"github.com/Cyrof/govault/pkg/cobraCLI"
)

func main() {
	f := fileIO.NewFileIO()
	logger.InitLogger(f)

	defer logger.Logger.Sync()

	cobraCLI.Execute()
}
