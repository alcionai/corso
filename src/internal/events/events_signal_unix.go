//go:build !windows
// +build !windows

package events

import (
	"context"
	"syscall"

	"github.com/armon/go-metrics"
	"golang.org/x/sys/unix"

	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	sentSignal = metrics.DefaultSignal
)

func signalDump(ctx context.Context) {
	if err := unix.Kill(syscall.Getpid(), metrics.DefaultSignal); err != nil {
		logger.CtxErr(ctx, err).Error("metrics interval signal")
	}
}
