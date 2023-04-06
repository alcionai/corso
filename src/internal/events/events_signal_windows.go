package events

import (
	"context"
	"syscall"

	"golang.org/x/sys/windows"
)

const (
	// On Windows, ctrl-break event appears as os.Interrupt(syscall.SIGINT)
	//https://pkg.go.dev/os/signal#hdr-Windows
	sentSignal = os.Interrupt
)

func signalDump(ctx context.Context) {
	err := windows.GenerateConsoleCtrlEvent(syscall.CTRL_BREAK_EVENT, uint32(syscall.Getpid()))
	if err != nil {
		logger.CtxErr(ctx, err).Error("metrics interval signal")
	}
}
