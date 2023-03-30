package events

import (
	"context"

	"github.com/alcionai/corso/src/pkg/logger"
)

func signalDump(ctx context.Context) {
	logger.Ctx(ctx).Warn("cannot send signal on Windows")
}
