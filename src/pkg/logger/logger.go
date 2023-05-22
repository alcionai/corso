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
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/str"
)

// Default location for writing logs, initialized in platform specific files
var userLogsDir string

var (
	logCore   *zapcore.Core
	loggerton *zap.SugaredLogger
)

type logLevel string

const (
	LLDebug    logLevel = "debug"
	LLInfo     logLevel = "info"
	LLWarn     logLevel = "warn"
	LLError    logLevel = "error"
	LLDisabled logLevel = "disabled"
)

type logFormat string

const (
	// use for cli/terminal
	LFText logFormat = "text"
	// use for cloud logging
	LFJSON logFormat = "json"
)

type piiAlg string

const (
	PIIHash      piiAlg = "hash"
	PIIMask      piiAlg = "mask"
	PIIPlainText piiAlg = "plaintext"
)

// flag names
const (
	DebugAPIFN          = "debug-api-calls"
	LogFileFN           = "log-file"
	LogFormatFN         = "log-format"
	LogLevelFN          = "log-level"
	ReadableLogsFN      = "readable-logs"
	MaskSensitiveDataFN = "mask-sensitive-data"
)

// flag values
var (
	DebugAPIFV          bool
	logFileFV           string
	LogFormatFV         string
	LogLevelFV          string
	ReadableLogsFV      bool
	MaskSensitiveDataFV bool

	ResolvedLogFile string // logFileFV after processing
	piiHandling     string // piiHandling after MaskSensitiveDataFV processing
)

const (
	Stderr = "stderr"
	Stdout = "stdout"
)

// Returns the default location for writing logs
func defaultLogLocation() string {
	return filepath.Join(
		userLogsDir,
		"corso",
		"logs",
		time.Now().UTC().Format("2006-01-02T15-04-05Z")+".log")
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
		string(LLInfo),
		fmt.Sprintf("set the log level to %s|%s|%s|%s", LLDebug, LLInfo, LLWarn, LLError))

	fs.StringVar(
		&LogFormatFV,
		LogFormatFN,
		string(LFText),
		fmt.Sprintf("set the log format to %s|%s", LFText, LFJSON))

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

	set := Settings{
		File:        defaultLogLocation(),
		Format:      LFText,
		Level:       LLInfo,
		PIIHandling: PIIPlainText,
	}

	// parse the os args list to find the log level flag
	if err := fs.Parse(args); err != nil {
		return set
	}

	// retrieve the user's preferred log level
	// defaults to "info"
	levelString, err := fs.GetString(LogLevelFN)
	if err != nil {
		return set
	}

	set.Level = logLevel(levelString)

	// retrieve the user's preferred log format
	// defaults to "text"
	formatString, err := fs.GetString(LogFormatFN)
	if err != nil {
		return set
	}

	set.Format = logFormat(formatString)

	// retrieve the user's preferred log file location
	// defaults to default log location
	lffv, err := fs.GetString(LogFileFN)
	if err != nil {
		return set
	}

	set.File = GetLogFile(lffv)
	ResolvedLogFile = set.File

	// retrieve the user's preferred PII handling algorithm
	// defaults to "plaintext"
	maskPII, err := fs.GetBool(MaskSensitiveDataFN)
	if err != nil {
		return set
	}

	if maskPII {
		set.PIIHandling = PIIHash
	}

	return set
}

// GetLogFile parses the log file.  Uses the provided value, if populated,
// then falls back to the env var, and then defaults to stderr.
func GetLogFile(logFileFlagVal string) string {
	if len(ResolvedLogFile) > 0 {
		fmt.Printf("\n-----\nalready resolved to: %+v\n-----\n", ResolvedLogFile)
		return ResolvedLogFile
	}

	r := logFileFlagVal

	fmt.Printf("\n-----\nflag val [%+v]\n-----\n", r)

	// if not specified, attempt to fall back to env declaration.
	if len(r) == 0 {
		r = os.Getenv("CORSO_LOG_FILE")
		fmt.Printf("\n-----\nuse env [%+v]\n-----\n", r)
	}

	// if no flag or env is specified, fall back to the default
	if len(r) == 0 {
		r = defaultLogLocation()
		fmt.Printf("\n-----\nuse default [%+v]\n-----\n", r)
	}

	if r == "-" {
		r = Stdout
		fmt.Printf("\n-----\nuse stdout\n-----\n")
	}

	if r != Stdout && r != Stderr {
		logdir := filepath.Dir(r)

		err := os.MkdirAll(logdir, 0o755)
		if err != nil {
			return Stderr
		}

		fmt.Printf("\n-----\nlogdir [%+v]\n-----\n", logdir)
	}

	return r
}

// Settings records the user's preferred logging settings.
type Settings struct {
	File        string    // what file to log to (alt: stderr, stdout)
	Format      logFormat // whether to format as text (console) or json (cloud)
	Level       logLevel  // what level to log at
	PIIHandling piiAlg    // how to obscure pii
}

// EnsureDefaults sets any non-populated settings to their default value.
// exported for testing without circular dependencies.
func (s Settings) EnsureDefaults() Settings {
	set := s

	levels := []logLevel{LLDisabled, LLDebug, LLInfo, LLWarn, LLError}
	if len(set.Level) == 0 || !slices.Contains(levels, set.Level) {
		set.Level = LLInfo
	}

	formats := []logFormat{LFText, LFJSON}
	if len(set.Format) == 0 || !slices.Contains(formats, set.Format) {
		set.Format = LFText
	}

	algs := []piiAlg{PIIPlainText, PIIMask, PIIHash}
	if len(set.PIIHandling) == 0 || !slices.Contains(algs, set.PIIHandling) {
		set.PIIHandling = piiAlg(str.First(piiHandling, string(PIIPlainText)))
	}

	if len(set.File) == 0 {
		set.File = GetLogFile("")
		ResolvedLogFile = set.File
	}

	return set
}

// ---------------------------------------------------------------------------
// constructors
// ---------------------------------------------------------------------------

func genLogger(set Settings) (*zapcore.Core, *zap.SugaredLogger) {
	// when testing, ensure debug logging matches the test.v setting
	for _, arg := range os.Args {
		if arg == `--test.v=true` {
			set.Level = LLDebug
		}
	}

	fmt.Printf("\n-----\nLOG SETTINGS FILE [%+v]\n-----\n", set.File)

	var (
		lgr  *zap.Logger
		err  error
		opts = []zap.Option{zap.AddStacktrace(zapcore.PanicLevel)}

		// set up a logger core to use as a fallback
		levelFilter = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			switch set.Level {
			case LLInfo:
				return lvl >= zapcore.InfoLevel
			case LLWarn:
				return lvl >= zapcore.WarnLevel
			case LLError:
				return lvl >= zapcore.ErrorLevel
			case LLDisabled:
				return false
			default:
				return true
			}
		})

		out            = zapcore.Lock(os.Stderr)
		consoleEncoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core           = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, out, levelFilter),
		)

		cfg zap.Config
	)

	switch set.Format {
	case LFJSON:
		cfg = setLevel(zap.NewProductionConfig(), set.Level)
		cfg.OutputPaths = []string{set.File}
	default:
		cfg = setLevel(zap.NewDevelopmentConfig(), set.Level)

		if ReadableLogsFV {
			opts = append(opts, zap.WithCaller(false))
			cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05.00")

			if set.File == Stderr || set.File == Stdout {
				cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			}
		}

		cfg.OutputPaths = []string{set.File}
	}

	// fall back to the core config if the default creation fails
	lgr, err = cfg.Build(opts...)
	if err != nil {
		lgr = zap.New(core)
	}

	return &core, lgr.Sugar()
}

func setLevel(cfg zap.Config, level logLevel) zap.Config {
	switch level {
	case LLInfo:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case LLWarn:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case LLError:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case LLDisabled:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	}

	return cfg
}

func singleton(set Settings) *zap.SugaredLogger {
	if loggerton != nil {
		return loggerton
	}

	if logCore != nil {
		lgr := zap.New(*logCore)
		loggerton = lgr.Sugar()

		return loggerton
	}

	set = set.EnsureDefaults()
	setCluesSecretsHash(set.PIIHandling)

	logCore, loggerton = genLogger(set)

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
	zsl := singleton(set)
	return Set(ctx, zsl), zsl
}

func setCluesSecretsHash(alg piiAlg) {
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

// CtxOrSeed attempts to retrieve the logger from the ctx.  If not found, it
// generates a logger with the given settings and adds it to the context.
func CtxOrSeed(ctx context.Context, set Settings) (context.Context, *zap.SugaredLogger) {
	l := ctx.Value(ctxKey)
	if l == nil {
		zsl := singleton(set)
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
		l = singleton(Settings{}.EnsureDefaults())
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
