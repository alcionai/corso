package crash

import (
	"context"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

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
func Recovery(ctx context.Context, r any, namespace string) error {
	var (
		err    error
		inFile string
		j      int
	)

	if r == nil {
		return nil
	}

	if re, ok := r.(error); ok {
		err = re
	} else if re, ok := r.(string); ok {
		err = clues.New(re)
	} else {
		err = clues.New(fmt.Sprintf("%v", r))
	}

	for i := 1; i < 10; i++ {
		_, file, line, ok := runtime.Caller(i)
		if j > 0 {
			if strings.Contains(file, "panic.go") {
				j = 0
			} else {
				inFile = fmt.Sprintf(": file %s - line %d", file, line)
				break
			}
		}

		// skip the location where Recovery() gets called.
		if j == 0 && ok && !strings.Contains(file, "panic.go") && !strings.Contains(file, "crash.go") {
			j++
		}
	}

	err = clues.Wrap(err, "panic recovery"+inFile).
		WithClues(ctx).
		With("stacktrace", string(debug.Stack()))
	logger.CtxErr(ctx, err).Error(namespace + " panic")

	return err
}
