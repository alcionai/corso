package logger_test

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// mock helpers
// ---------------------------------------------------------------------------

const (
	loglevel = "info"
	logfile  = "stderr"
	itemID   = "item_id"
)

var err error

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

	// 1. Keep messsages short.  Lowercase text, no ending punctuation.
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
	log := logger.Ctx(context.Background())

	// 1. Clues Ctx values are always added in .Ctx(); you don't
	// need to add them directly.
	//
	// preferred
	ctx := clues.Add(context.Background(), "item_id", itemID)
	logger.Ctx(ctx).Info("getting item")
	// avoid
	ctx = clues.Add(context.Background(), "item_id", itemID)
	logger.Ctx(ctx).With(clues.In(ctx).Slice()...).Info("getting item")

	// 2. Always extract structured data from errors.
	//
	// preferred
	log.With("err", err).Errorw("getting item", clues.InErr(err).Slice()...)
	// avoid
	log.Errorw("getting item", "err", err)

	// TODO(keepers): PII
	// 3. Protect pii in logs.
}
