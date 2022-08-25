package logger

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/alcionai/corso/cli/print"
)

var (
	logCore   *zapcore.Core
	loggerton *zap.SugaredLogger
	// logging level flag
	// TODO: infer default based on environment.
	llFlag = "info"
)

type logLevel int

const (
	Development logLevel = iota
	Info
	Warn
	Production
)

// adds the persistent flag --log-level to the provided command.
// defaults to "info".
// This is a hack for help displays.  Due to seeding the context, we
// need to parse the log level before we execute the command.
func AddLogLevelFlag(parent *cobra.Command) {
	fs := parent.PersistentFlags()
	fs.StringVar(&llFlag, "log-level", "info", "set the log level to debug|info|warn|error")
}

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
// It also parses the command line for flag values prior to executing
// cobra.  This early parsing is necessary since logging depends on
// a seeded context prior to cobra evaluating flags.
func Seed(ctx context.Context) (ctxOut context.Context, zsl *zap.SugaredLogger) {
	level := Info

	// this func handles composing the return values whether or not an error occurs
	defer func() {
		zsl = singleton(level)
		ctxOut = context.WithValue(ctx, ctxKey, zsl)
	}()

	fs := pflag.NewFlagSet("seed-logger", pflag.ContinueOnError)
	fs.ParseErrorsWhitelist.UnknownFlags = true
	fs.String("log-level", "info", "set the log level to debug|info|warn|error")

	// parse the os args list to find the log level flag
	if err := fs.Parse(os.Args[1:]); err != nil {
		print.Err(ctx, err.Error())
		return
	}

	// retrieve the user's preferred log level
	// automatically defaults to "info"
	levelString, err := fs.GetString("log-level")
	if err != nil {
		print.Err(ctx, err.Error())
		return
	}

	level = levelOf(levelString)
	return // return values handled in defer
}

// Ctx retrieves the logger embedded in the context.
func Ctx(ctx context.Context) *zap.SugaredLogger {
	l := ctx.Value(ctxKey)
	if l == nil {
		return singleton(levelOf(llFlag))
	}

	return l.(*zap.SugaredLogger)
}

// transforms the llevel flag value to a logLevel enum
func levelOf(lvl string) logLevel {
	switch lvl {
	case "debug":
		return Development
	case "warn":
		return Warn
	case "error":
		return Production
	}
	return Info
}
