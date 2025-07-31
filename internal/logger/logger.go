package logger

import (
	"os"
	"sync"

	"github.com/Cyrof/govault/internal/fileIO"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	warn   string
	once   sync.Once
	Logger *zap.SugaredLogger
)

func InitLogger(f *fileIO.FileIO) {
	once.Do(func() {
		// try to access .env file (dev only)
		if err := godotenv.Load(); err != nil {
			warn = "No .env file found, falling back to system environment variables"
		}

		env := os.Getenv("GO_ENV")
		// setup dev logger
		if env == "dev" {
			config := zap.NewDevelopmentConfig()
			config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

			logger, err := config.Build()
			if err != nil {
				panic("failed to initialise development logger: " + err.Error())
			}

			Logger = logger.Sugar()
			if warn != "" {
				Logger.Warn(warn)
			}
			return
		}

		// make sure file path exist
		if err := f.EnsureVaultDir(); err != nil {
			panic("failed to initialise directory: " + err.Error())
		}

		// config lumberjack
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   f.LogPath,
			MaxSize:    5,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		})

		// zap logger config
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			writer,
			zap.InfoLevel,
		)

		logger := zap.New(core)

		Logger = logger.Sugar()

		if warn != "" {
			Logger.Warn(warn)
		}
	})
}
