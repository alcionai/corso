package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logCore *zapcore.Core

func coreSingleton() *zapcore.Core {
	if logCore == nil {
		// level handling
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.ErrorLevel
		})
		// level-based output
		consoleDebugging := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)
		// encoder type
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		// combine into a logger core
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		)
		logCore = &core
	}
	return logCore
}

type loggingKey string

const ctxKey loggingKey = "corsoLogger"

// Seed embeds a logger into the context for later retrieval.
func Seed(ctx context.Context) (context.Context, *zap.SugaredLogger) {
	l := zap.New(*coreSingleton())
	s := l.Sugar()
	return context.WithValue(ctx, ctxKey, s), s
}

// Ctx retrieves the logger embedded in the context.
func Ctx(ctx context.Context) *zap.SugaredLogger {
	l := ctx.Value(ctxKey)
	if l == nil {
		return zap.New(*coreSingleton()).Sugar()
	}
	return l.(*zap.SugaredLogger)
}
