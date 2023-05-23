package tester

import (
	"context"
	"os"
	"testing"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/logger"
)

func NewContext(t *testing.T) (context.Context, func()) {
	level := logger.LLInfo
	format := logger.LFText

	for _, a := range os.Args {
		if a == "-test.v=true" {
			level = logger.LLDebug
		}
	}

	ls := logger.Settings{
		Level:  level,
		Format: format,
	}

	//nolint:forbidigo
	ctx, _ := logger.CtxOrSeed(context.Background(), ls)

	// ensure logs can be easily associated with each test
	if t != nil {
		ctx = clues.Add(ctx, "test_name", t.Name())
	}

	return ctx, func() { logger.Flush(ctx) }
}

func WithContext(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
) (context.Context, func()) {
	ls := logger.Settings{
		Level:  logger.LLDebug,
		Format: logger.LFText,
	}
	ctx, _ = logger.CtxOrSeed(ctx, ls)

	// ensure logs can be easily associated with each test
	// todo: replace with test name.  starting off with
	// uuid to avoid million-line PR change.
	ctx = clues.Add(ctx, "test_name", t.Name())

	return ctx, func() { logger.Flush(ctx) }
}
