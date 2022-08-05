package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logCore   *zapcore.Core
	loggerton *zap.SugaredLogger
)

type logLevel int

const (
	Development logLevel = iota
	Info
	Warn
	Production
)

func singleton(level logLevel) *zap.SugaredLogger {
	if loggerton != nil {
		return loggerton
	}

	// set up a logger core to use as a fallback
	levelFilter := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		switch level {
		case Info:
			return lvl >= zapcore.InfoLevel
		case Warn:
			return lvl >= zapcore.WarnLevel
		case Production:
			return lvl >= zapcore.ErrorLevel
		default:
			return true
		}
	})
	out := zapcore.Lock(os.Stderr)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, out, levelFilter),
	)
	logCore = &core

	// then try to set up a logger directly
	var (
		lgr *zap.Logger
		err error
	)
	if level != Production {
		cfg := zap.NewDevelopmentConfig()
		switch level {
		case Info:
			cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		case Warn:
			cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
		}
		lgr, err = cfg.Build()
	} else {
		lgr, err = zap.NewProduction()
	}

	// fall back to the core config if the default creation fails
	if err != nil {
		lgr = zap.New(*logCore)
	}

	loggerton = lgr.Sugar()
	return loggerton
}

type loggingKey string

const ctxKey loggingKey = "corsoLogger"

// Seed embeds a logger into the context for later retrieval.
func Seed(ctx context.Context) (context.Context, *zap.SugaredLogger) {
	l := singleton(0)
	return context.WithValue(ctx, ctxKey, l), l
}

// Ctx retrieves the logger embedded in the context.
func Ctx(ctx context.Context) *zap.SugaredLogger {
	l := ctx.Value(ctxKey)
	if l == nil {
		return singleton(0)
	}
	return l.(*zap.SugaredLogger)
}
