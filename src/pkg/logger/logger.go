package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Default location for writing logs, initialized in platform specific files
var userLogsDir string

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
	Disabled
)

// flag names
const (
	DebugAPIFN          = "debug-api-calls"
	LogFileFN           = "log-file"
	LogLevelFN          = "log-level"
	ReadableLogsFN      = "readable-logs"
	MaskSensitiveDataFN = "mask-sensitive-data"
)

// flag values
var (
	DebugAPIFV          bool
	logFileFV           = ""
	LogLevelFV          = "info"
	ReadableLogsFV      bool
	MaskSensitiveDataFV bool
	SensitiveDataCfg    = PIIPlainText

	LogFile string // logFileFV after processing
)

const (
	Stderr = "stderr"
	Stdout = "stdout"

	PIIHash      = "hash"
	PIIMask      = "mask"
	PIIPlainText = "plaintext"

	LLDebug    = "debug"
	LLInfo     = "info"
	LLWarn     = "warn"
	LLError    = "error"
	LLDisabled = "disabled"
)

// Returns the default location for writing logs
func defaultLogLocation() string {
	return filepath.Join(userLogsDir, "corso", "logs", time.Now().UTC().Format("2006-01-02T15-04-05Z")+".log")
}

// adds the persistent flag --log-level and --log-file to the provided command.
// defaults to "info" and the default log location.
// This is a hack for help displays.  Due to seeding the context, we also
// need to parse the log level before we execute the command.
func AddLoggingFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()

	addFlags(fs, "corso-<timestamp>.log")

	//nolint:errcheck
	fs.MarkHidden(ReadableLogsFN)
}

// internal deduplication for adding flags
func addFlags(fs *pflag.FlagSet, defaultFile string) {
	fs.StringVar(
		&LogLevelFV,
		LogLevelFN,
		LLInfo,
		fmt.Sprintf("set the log level to %s|%s|%s|%s", LLDebug, LLInfo, LLWarn, LLError))

	// The default provided here is only for help info
	fs.StringVar(&logFileFV, LogFileFN, defaultFile, "location for writing logs, use '-' for stdout")
	fs.BoolVar(&DebugAPIFV, DebugAPIFN, false, "add non-2xx request/response errors to logging")

	fs.BoolVar(
		&ReadableLogsFV,
		ReadableLogsFN,
		false,
		"minimizes log output for console readability: removes the file and date, colors the level")

	fs.BoolVar(
		&MaskSensitiveDataFV,
		MaskSensitiveDataFN,
		false,
		"anonymize personal data in log output")
}

// Settings records the user's preferred logging settings.
type Settings struct {
	File        string // what file to log to (alt: stderr, stdout)
	Level       string // what level to log at
	PIIHandling string // how to obscure pii
}

// Due to races between the lazy evaluation of flags in cobra and the
// need to init logging behavior in a ctx, log-level and log-file gets
// pre-processed manually here using pflags.  The canonical
// AddLogLevelFlag() and AddLogFileFlag() ensures the flags are
// displayed as part of the help/usage output.
func PreloadLoggingFlags(args []string) Settings {
	fs := pflag.NewFlagSet("seed-logger", pflag.ContinueOnError)
	fs.ParseErrorsWhitelist.UnknownFlags = true
	addFlags(fs, "")

	// prevents overriding the corso/cobra help processor
	fs.BoolP("help", "h", false, "")

	if MaskSensitiveDataFV {
		SensitiveDataCfg = PIIHash
	}

	ls := Settings{
		File:        "",
		Level:       LogLevelFV,
		PIIHandling: SensitiveDataCfg,
	}

	// parse the os args list to find the log level flag
	if err := fs.Parse(args); err != nil {
		return ls
	}

	// retrieve the user's preferred log level
	// automatically defaults to "info"
	levelString, err := fs.GetString(LogLevelFN)
	if err != nil {
		return ls
	}

	ls.Level = levelString

	// retrieve the user's preferred log file location
	// automatically defaults to default log location
	lffv, err := fs.GetString(LogFileFN)
	if err != nil {
		return ls
	}

	ls.File = GetLogFile(lffv)
	LogFile = ls.File

	// retrieve the user's preferred PII handling algorithm
	// automatically defaults to default log location
	pii, err := fs.GetString(MaskSensitiveDataFN)
	if err != nil {
		return ls
	}

	ls.PIIHandling = pii

	return ls
}

// GetLogFile parses the log file.  Uses the provided value, if populated,
// then falls back to the env var, and then defaults to stderr.
func GetLogFile(logFileFlagVal string) string {
	r := logFileFlagVal

	// if not specified, attempt to fall back to env declaration.
	if len(r) == 0 {
		r = os.Getenv("CORSO_LOG_FILE")
	}

	// if no flag or env is specified, fall back to the default
	if len(r) == 0 {
		r = defaultLogLocation()
	}

	if r == "-" {
		r = Stdout
	}

	if r != Stdout && r != Stderr {
		logdir := filepath.Dir(r)

		err := os.MkdirAll(logdir, 0o755)
		if err != nil {
			return Stderr
		}
	}

	return r
}

func genLogger(level logLevel, logfile string) (*zapcore.Core, *zap.SugaredLogger) {
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
		lgr  *zap.Logger
		err  error
		opts = []zap.Option{zap.AddStacktrace(zapcore.PanicLevel)}
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

		if ReadableLogsFV {
			opts = append(opts, zap.WithCaller(false))
			cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05.00")

			if logfile == Stderr || logfile == Stdout {
				cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			}
		}

		cfg.OutputPaths = []string{logfile}
		lgr, err = cfg.Build(opts...)
	} else {
		cfg := zap.NewProductionConfig()
		cfg.OutputPaths = []string{logfile}
		lgr, err = cfg.Build(opts...)
	}

	// fall back to the core config if the default creation fails
	if err != nil {
		lgr = zap.New(core)
	}

	return &core, lgr.Sugar()
}

func singleton(level logLevel, logfile string) *zap.SugaredLogger {
	if loggerton != nil {
		return loggerton
	}

	if logCore != nil {
		lgr := zap.New(*logCore)
		loggerton = lgr.Sugar()

		return loggerton
	}

	logCore, loggerton = genLogger(level, logfile)

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
func Seed(ctx context.Context, set Settings) (context.Context, *zap.SugaredLogger) {
	if len(set.Level) == 0 {
		set.Level = LLInfo
	}

	setCluesSecretsHash(set.PIIHandling)

	zsl := singleton(levelOf(set.Level), set.File)

	return Set(ctx, zsl), zsl
}

func setCluesSecretsHash(alg string) {
	switch alg {
	case PIIHash:
		// TODO: a persistent hmac key for each tenant would be nice
		// as a way to correlate logs across runs.
		clues.SetHasher(clues.DefaultHash())
	case PIIMask:
		clues.SetHasher(clues.HashCfg{HashAlg: clues.Flatmask})
	case PIIPlainText:
		clues.SetHasher(clues.NoHash())
	}
}

// SeedLevel generates a logger within the context with the given log-level.
func SeedLevel(ctx context.Context, level logLevel) (context.Context, *zap.SugaredLogger) {
	l := ctx.Value(ctxKey)
	if l == nil {
		logfile := os.Getenv("CORSO_LOG_FILE")

		if len(logfile) == 0 {
			logfile = defaultLogLocation()
		}

		zsl := singleton(level, logfile)

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
		return singleton(levelOf(LogLevelFV), defaultLogLocation())
	}

	return l.(*zap.SugaredLogger).With(clues.In(ctx).Slice()...)
}

// CtxErr retrieves the logger embedded in the context
// and packs all of the structured data in the error inside it.
func CtxErr(ctx context.Context, err error) *zap.SugaredLogger {
	return Ctx(ctx).
		With(
			"error", err,
			"error_labels", clues.Labels(err)).
		With(clues.InErr(err).Slice()...)
}

// transforms the llevel flag value to a logLevel enum
func levelOf(lvl string) logLevel {
	switch lvl {
	case LLDebug:
		return Development
	case LLWarn:
		return Warn
	case LLError:
		return Production
	case LLDisabled:
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

// ------------------------------------------------------------------------------------------------
// io.writer that writes values to the logger
// ------------------------------------------------------------------------------------------------

// Writer is a wrapper that turns the logger embedded in
// the given ctx into an io.Writer.  All logs are currently
// info-level.
type Writer struct {
	Ctx context.Context
}

func (w Writer) Write(p []byte) (int, error) {
	Ctx(w.Ctx).Info(string(p))
	return len(p), nil
}
