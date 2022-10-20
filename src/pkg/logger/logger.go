package logger

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
// This is a hack for help displays.  Due to seeding the context, we also
// need to parse the log level before we execute the command.
func AddLogLevelFlag(parent *cobra.Command) {
	fs := parent.PersistentFlags()
	fs.StringVar(&llFlag, "log-level", "info", "set the log level to debug|info|warn|error")
}

// Due to races between the lazy evaluation of flags in cobra and the need to init logging
// behavior in a ctx, log-level gets pre-processed manually here using pflags.  The canonical
// AddLogLevelFlag() ensures the flag is displayed as part of the help/usage output.
func PreloadLogLevel() string {
	fs := pflag.NewFlagSet("seed-logger", pflag.ContinueOnError)
	fs.ParseErrorsWhitelist.UnknownFlags = true
	fs.String("log-level", "info", "set the log level to debug|info|warn|error")
	// prevents overriding the corso/cobra help processor
	fs.BoolP("help", "h", false, "")

	// parse the os args list to find the log level flag
	if err := fs.Parse(os.Args[1:]); err != nil {
		return "info"
	}

	// retrieve the user's preferred log level
	// automatically defaults to "info"
	levelString, err := fs.GetString("log-level")
	if err != nil {
		return "info"
	}

	return levelString
}

func genLogger(level logLevel) (*zapcore.Core, *zap.SugaredLogger) {
	// when testing, ensure debug logging matches the test.v setting
	for _, arg := range os.Args {
		if arg == `--test.v=true` {
			level = Development
		}
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

	return &core, lgr.Sugar()
}

func singleton(level logLevel) *zap.SugaredLogger {
	if loggerton != nil {
		return loggerton
	}

	if logCore != nil {
		lgr := zap.New(*logCore)
		loggerton = lgr.Sugar()

		return loggerton
	}

	logCore, loggerton = genLogger(level)

	return loggerton
}

type loggingKey string

const ctxKey loggingKey = "corsoLogger"

// Seed embeds a logger into the context for later retrieval.
// It also parses the command line for flag values prior to executing
// cobra.  This early parsing is necessary since logging depends on
// a seeded context prior to cobra evaluating flags.
func Seed(ctx context.Context, lvl string) (context.Context, *zap.SugaredLogger) {
	if len(lvl) == 0 {
		lvl = "info"
	}

	zsl := singleton(levelOf(lvl))
	ctxOut := context.WithValue(ctx, ctxKey, zsl)

	return ctxOut, zsl
}

// SeedLevel embeds a logger into the context with the given log-level.
func SeedLevel(ctx context.Context, level logLevel) (context.Context, *zap.SugaredLogger) {
	l := ctx.Value(ctxKey)
	if l == nil {
		zsl := singleton(level)
		ctxWV := context.WithValue(ctx, ctxKey, zsl)

		return ctxWV, zsl
	}

	return ctx, l.(*zap.SugaredLogger)
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

// Flush writes out all buffered logs.
func Flush(ctx context.Context) {
	_ = Ctx(ctx).Sync()
}
