package events

import (
	"context"
	"github.com/alcionai/corso/src/pkg/logger"
	"golang.org/x/sys/windows"
	"signal"
	"syscall"
)

func signalDump(ctx context.Context) {
	//logger.Ctx(ctx).Warn("cannot send signal on Windows")
	windows.GenerateConsoleCtrlEvent(syscall.CTRL_BREAK_EVENT, syscall.Getpid())
}
