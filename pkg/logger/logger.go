package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func MustInit(name string, production bool) *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	atom := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	loggerConfig.Level = atom
	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatalf("failed to create logger: %s\n", err)
	}

	logger = logger.Named(name)

	// DEBUG < INFO < WARN < ERROR < DPanic < PANIC < FATAL
	if production {
		atom.SetLevel(zap.InfoLevel)
	}

	return logger
}
