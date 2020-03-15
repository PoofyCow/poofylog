package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var GlobalLogger *zap.Logger

func init() {
	var (
		outFilter = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.InfoLevel || lvl == zapcore.DebugLevel
		})
		errFilter = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl > zapcore.InfoLevel
		})

		stdOut = zapcore.Lock(os.Stdout)
		stdErr = zapcore.Lock(os.Stderr)

		encoderConfig = zap.NewProductionEncoderConfig()
	)

	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.LevelKey = "level"
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder

	console := zapcore.NewConsoleEncoder(encoderConfig)

	GlobalLogger = zap.New(zapcore.NewTee(
		zapcore.NewCore(console, stdOut, outFilter),
		zapcore.NewCore(console, stdErr, errFilter),
	))
}
