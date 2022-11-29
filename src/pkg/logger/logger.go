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
	Disabled
)

const logLevelFN = "log-level"

// adds the persistent flag --log-level to the provided command.
// defaults to "info".
// This is a hack for help displays.  Due to seeding the context, we also
// need to parse the log level before we execute the command.
func AddLogLevelFlag(parent *cobra.Command) {
	fs := parent.PersistentFlags()
	fs.StringVar(&llFlag, logLevelFN, "info", "set the log level to debug|info|warn|error")
}

// Due to races between the lazy evaluation of flags in cobra and the need to init logging
// behavior in a ctx, log-level gets pre-processed manually here using pflags.  The canonical
// AddLogLevelFlag() ensures the flag is displayed as part of the help/usage output.
func PreloadLogLevel() string {
	fs := pflag.NewFlagSet("seed-logger", pflag.ContinueOnError)
	fs.ParseErrorsWhitelist.UnknownFlags = true
	fs.String(logLevelFN, "info", "set the log level to debug|info|warn|error")
	// prevents overriding the corso/cobra help processor
	fs.BoolP("help", "h", false, "")

	// parse the os args list to find the log level flag
	if err := fs.Parse(os.Args[1:]); err != nil {
		return "info"
	}

	// retrieve the user's preferred log level
	// automatically defaults to "info"
	levelString, err := fs.GetString(logLevelFN)
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
		case Disabled:
			return false
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
		case Disabled:
			cfg.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
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

// ------------------------------------------------------------------------------------------------
// context management
// ------------------------------------------------------------------------------------------------

type loggingKey string

const ctxKey loggingKey = "corsoLogger"

// Seed generates a logger within the context for later retrieval.
// It also parses the command line for flag values prior to executing
// cobra.  This early parsing is necessary since logging depends on
// a seeded context prior to cobra evaluating flags.
func Seed(ctx context.Context, lvl string) (context.Context, *zap.SugaredLogger) {
	if len(lvl) == 0 {
		lvl = "info"
	}

	zsl := singleton(levelOf(lvl))

	return Set(ctx, zsl), zsl
}

// SeedLevel generates a logger within the context with the given log-level.
func SeedLevel(ctx context.Context, level logLevel) (context.Context, *zap.SugaredLogger) {
	l := ctx.Value(ctxKey)
	if l == nil {
		zsl := singleton(level)
		return Set(ctx, zsl), zsl
	}

	return ctx, l.(*zap.SugaredLogger)
}

// Set allows users to embed their own zap.SugaredLogger within the context.
func Set(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	if logger == nil {
		return ctx
	}

	return context.WithValue(ctx, ctxKey, logger)
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
	case "disabled":
		return Disabled
	}

	return Info
}

// Flush writes out all buffered logs.
func Flush(ctx context.Context) {
	_ = Ctx(ctx).Sync()
}

// ------------------------------------------------------------------------------------------------
// log wrapper for downstream api compliance
// ------------------------------------------------------------------------------------------------

type wrapper struct {
	zap.SugaredLogger

	forceDebugLogLevel bool
}

func (w *wrapper) process(opts ...option) {
	for _, opt := range opts {
		opt(w)
	}
}

type option func(*wrapper)

// ForceDebugLogLevel reduces all logs emitted in the wrapper to
// debug level, independent of their original log level.  Useful
// for silencing noisy dependency packages without losing the info
// altogether.
func ForceDebugLogLevel() option {
	return func(w *wrapper) {
		w.forceDebugLogLevel = true
	}
}

// Wrap returns the logger in the package with an extended api used for
// dependency package interface compliance.
func WrapCtx(ctx context.Context, opts ...option) *wrapper {
	return Wrap(Ctx(ctx), opts...)
}

// Wrap returns the sugaredLogger with an extended api used for
// dependency package interface compliance.
func Wrap(zsl *zap.SugaredLogger, opts ...option) *wrapper {
	w := &wrapper{SugaredLogger: *zsl}
	w.process(opts...)

	return w
}

func (w *wrapper) Logf(tmpl string, args ...any) {
	if w.forceDebugLogLevel {
		w.SugaredLogger.Debugf(tmpl, args...)
		return
	}

	w.SugaredLogger.Infof(tmpl, args...)
}

func (w *wrapper) Errorf(tmpl string, args ...any) {
	if w.forceDebugLogLevel {
		w.SugaredLogger.Debugf(tmpl, args...)
		return
	}

	w.SugaredLogger.Errorf(tmpl, args...)
}
