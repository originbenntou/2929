package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is Log object.
var Logger *zap.Logger

func init() {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.InfoLevel)

	myConfig := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "Stack",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		Development:      true,
		OutputPaths:      []string{"stdout", "/var/log/ap/ap.log"},
		ErrorOutputPaths: []string{"stderr", "/var/log/ap/ap.log"},
	}
	var err error
	Logger, err = myConfig.Build()
	if err != nil {
		panic(err)
	}
}
