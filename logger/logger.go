package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() Logger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	topicDebugging := zapcore.AddSync(io.Discard)
	topicErrors := zapcore.AddSync(io.Discard)

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	applicationEncoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(applicationEncoder, topicErrors, highPriority),
		zapcore.NewCore(applicationEncoder, consoleErrors, highPriority),
		zapcore.NewCore(applicationEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(applicationEncoder, consoleDebugging, lowPriority),
	)

	zapLogger := zap.New(core).Sugar()
	defer zapLogger.Sync()

	return Logger{zapLogger}
}

func (l Logger) Error(message string, err error) {
	l.Errorw(message, "error", err.Error())
}
