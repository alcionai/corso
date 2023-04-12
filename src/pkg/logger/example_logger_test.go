package logger_test

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// mock helpers
// ---------------------------------------------------------------------------

const (
	loglevel = "info"
	logfile  = logger.Stderr
	itemID   = "item_id"
)

var (
	err         error
	itemPath, _ = path.Build("tid", "own", path.ExchangeService, path.ContactsCategory, false, "foo")
)

// ---------------------------------------------------------------------------
// examples
// ---------------------------------------------------------------------------

// ExampleSeed showcases seeding a logger into the context.
func Example_seed() {
	// Before logging, a logger instance first needs to get seeded into
	// the context.  Seeding only needs to be done once.  For example
	// Corso's CLI layer seeds the logger in the cli initialization.
	ctx := context.Background()
	ctx, log := logger.Seed(ctx, loglevel, logfile)

	// SDK consumers who configure their own zap logger can Set their logger
	// into the context directly, instead of Seeding a new one.
	ctx = logger.Set(ctx, log)

	// logs should always be flushed before exiting whichever func
	// seeded the logger.
	defer func() {
		_ = log.Sync() // flush all logs in the buffer
	}()

	// downstream, the logger will retrieve its configuration from
	// the context.
	func(ctx context.Context) {
		log := logger.Ctx(ctx)
		log.Info("hello, world!")
	}(ctx)
}

// ExampleLoggerStandards reviews code standards around logging in Corso.
func Example_logger_standards() {
	log := logger.Ctx(context.Background())

	// 1. Keep messages short.  Lowercase text, no ending punctuation.
	// This ensures logs are easy to scan, and simple to grok.
	//
	// preferred
	log.Info("getting item")
	// avoid
	log.Info("Getting one item from the service so that we can send it through the item feed.")

	// 2. Do not fmt values into the message.  Use With() or -w() to add structured data.
	// By keeping dynamic data in a structured format, we maximize log readability,
	// and make logs very easy to search or filter in bulk, and very easy to control pii.
	//
	// preferred
	log.With("err", err).Error("getting item")
	log.Errorw("getting item", "err", err)
	// avoid
	log.Errorf("getting item %s: %v", itemID, err)

	// 3. Give data keys reasonable namespaces.  Use snake_case.
	// Overly generic keys can collide unexpectedly.
	//
	// preferred
	log.With("item_id", itemID).Info("getting item")
	// avoid
	log.With("id", itemID).Error("getting item")

	// 4. Avoid Warn-level logging.  Prefer Info or Error.
	// Minimize confusion/contention about what level a log
	// "should be".  Error during a failure, Info (or Debug)
	// otherwise.
	//
	// preferred
	log.With("err", err).Error("getting item")
	// avoid
	log.With("err", err).Warn("getting item")

	// 4. Avoid Panic/Fatal-level logging.  Prefer Error.
	// Panic and Fatal logging can crash the application without
	// flushing buffered logs and finishing out other telemetry.
	//
	// preferred
	log.With("err", err).Error("unable to connect")
	// avoid
	log.With("err", err).Panic("unable to connecct")
}

// ExampleLoggerCluesStandards reviews code standards around using the Clues package while logging.
func Example_logger_clues_standards() {
	ctx := clues.Add(context.Background(), "foo", "bar")
	log := logger.Ctx(ctx)

	// 1. Clues Ctx values are always added in .Ctx(); you don't
	// need to add them directly.
	//
	// preferred
	ctx = clues.Add(ctx, "item_id", itemID)
	logger.Ctx(ctx).Info("getting item")
	//
	// avoid
	ctx = clues.Add(ctx, "item_id", itemID)
	logger.Ctx(ctx).With(clues.In(ctx).Slice()...).Info("getting item")

	// 2. Always extract structured data from errors.
	//
	// preferred
	log.With("error", err).Errorw("getting item", clues.InErr(err).Slice()...)
	//
	// avoid
	log.Errorw("getting item", "err", err)
	//
	// you can use the logger helper CtxErr() for the same results.
	// This helps to ensure all error values get packed into the logs
	// in the expected format.
	logger.CtxErr(ctx, err).Error("getting item")

	// 3. Protect pii in logs.
	// When it comes to protecting sensitive information, we only want
	// to hand loggers (and, by extension, clues errors) using one of
	// three approaches to securing values.
	//
	// First: plain, unhidden data.  This can only be logged if we are
	// absolutely assured that this data does not expose sensitive
	// information for a user.  Eg: internal ids and enums are fine to
	// log in plain text.  Everything else must be considered wisely.
	//
	// Second: manually concealed values.  Strings containing sensitive
	// info, and structs from external pacakges containing sensitive info,
	// can be logged by manually wrapping them with a clues.Hide() call.
	// Ex: clues.Hide(userName).  This will hash the value according to
	// the user's hash algorithm configuration.
	//
	// Third: structs that comply with clues.Concealer.  The Concealer
	// interface requires a struct to comply with Conceal() (for cases
	// where the struct is handed to a clues aggregator directly), and
	// fmt's Format(state, verb), where the assumption is the standard
	// format writer will be replaced with a Conceal() call (for cases
	// where the struct is handed to some non-compliant formatter/printer).
	//
	// preferred
	log.With(
		// internal type, safe to log plainly
		"resource_type", connector.Users,
		// string containing sensitive info, wrap with Hide()
		"user_name", clues.Hide("your_user_name@microsoft.example"),
		// a concealer-compliant struct, safe to add plainly
		"storage_path", itemPath,
	)
}
