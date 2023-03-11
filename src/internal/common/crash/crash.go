package crash

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/logger"
)

// Recovery provides a deferrable func that can be called
// to recover from, and log context about, crashes.
// If an error is returned, then a panic recovery occurred.
//
// Call it as follows:
//
//	defer func() {
//		if crErr := crash.Recovery(ctx, recover()); crErr != nil {
//			err = crErr // err needs to be a named return variable
//		}
//	}()
func Recovery(ctx context.Context, r any) error {
	var (
		err    error
		inFile string
	)

	if r != nil {
		if re, ok := r.(error); ok {
			err = re
		} else if re, ok := r.(string); ok {
			err = clues.New(re)
		} else {
			err = clues.New(fmt.Sprintf("%v", r))
		}

		_, file, _, ok := runtime.Caller(3)
		if ok {
			inFile = " in file: " + file
		}

		err = clues.Wrap(err, "panic recovery"+inFile).
			WithClues(ctx).
			With("stacktrace", string(debug.Stack()))
		logger.CtxErr(ctx, err).Error("backup panic")
	}

	return err
}
