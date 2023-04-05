package events

import (
	"context"
	"syscall"

	"golang.org/x/sys/windows"
)

func signalDump(ctx context.Context) {
	//logger.Ctx(ctx).Warn("cannot send signal on Windows")
	windows.GenerateConsoleCtrlEvent(syscall.CTRL_BREAK_EVENT, uint32(syscall.Getpid()))
}
