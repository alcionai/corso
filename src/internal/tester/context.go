package tester

import (
	"context"
	"os"

	"github.com/alcionai/clues"
	"github.com/google/uuid"

	"github.com/alcionai/corso/src/pkg/logger"
)

func NewContext(t TestT) (context.Context, func()) {
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
	ctx = enrichTestCtx(t, ctx)

	return ctx, func() { logger.Flush(ctx) }
}

func WithContext(
	t TestT,
	ctx context.Context, //revive:disable-line:context-as-argument
) (context.Context, func()) {
	ls := logger.Settings{
		Level:  logger.LLDebug,
		Format: logger.LFText,
	}
	ctx, _ = logger.CtxOrSeed(ctx, ls)
	ctx = enrichTestCtx(t, ctx)

	return ctx, func() { logger.Flush(ctx) }
}

func enrichTestCtx(
	t TestT,
	ctx context.Context, //revive:disable-line:context-as-argument
) context.Context {
	if t == nil {
		return ctx
	}

	// ensure logs can be easily associated with each test
	LogTimeOfTest(t)

	ctx = clues.Add(
		ctx,
		// the actual test name, in case you want to look up
		// logs correlated to a certain test.
		"test_name", t.Name(),
		// an arbitrary uuid might be easier to match on when
		// looking up logs, in case of common log test names.
		"test_uuid", uuid.NewString())

	return ctx
}
