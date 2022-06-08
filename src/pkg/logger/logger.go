package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func subLogger() zerolog.Logger {
	return log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

type loggingKey string

const ctxKey loggingKey = "corsoLogger"

// Seed embeds a logger into the context for later retrieval.
func Seed(ctx context.Context) context.Context {
	return subLogger().WithContext(ctx)
}

// Ctx retrieves the logger embedded in the context.
func Ctx(ctx context.Context) *zerolog.Logger {
	return log.Ctx(ctx)
}
