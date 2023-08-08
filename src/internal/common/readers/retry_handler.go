package readers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"syscall"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/logger"
)

var _ io.ReadCloser = &resetRetryHandler{}

const (
	minSleepTime   = 3
	numMaxRetries  = 3
	rangeHeaderKey = "Range"
	// One-sided range like this is defined as starting at the given byte and
	// extending to the end of the item.
	rangeHeaderOneSidedValueTmpl = "bytes=%d-"
)

// Could make this per wrapper instance if we need additional flexibility
// between callers.
var retryErrs = []error{
	syscall.ECONNRESET,
}

type Getter interface {
	// SupportsRange returns true if this Getter supports adding Range headers to
	// the Get call. Otherwise returns false.
	SupportsRange() bool
	// Get attempts to get another reader for the data this reader is returning.
	// headers denotes any additional headers that should be added to the request,
	// like a Range header.
	//
	// Don't allow passing a URL to Get so that we can hide the fact that some
	// components may need to dynamically refresh the fetch URL (i.e. OneDrive)
	// from this wrapper.
	//
	// Get should encapsulate all error handling and status code checking required
	// for the component. This function is called both during NewResetRetryHandler
	// and Read so it's possible to discover errors with the item prior to
	// informing other components about it if desired.
	Get(ctx context.Context, headers map[string]string) (io.ReadCloser, error)
}

// NewResetRetryHandler returns an io.ReadCloser with the reader initialized to
// the result of getter. The reader is eagerly initialized during this call so
// if callers of this function want to delay initialization they should wrap
// this reader in a lazy initializer.
//
// Selected errors that the reader hits during Read calls (e.x.
// syscall.ECONNRESET) will be automatically retried by the returned reader.
func NewResetRetryHandler(
	ctx context.Context,
	getter Getter,
) (*resetRetryHandler, error) {
	rrh := &resetRetryHandler{
		ctx:    ctx,
		getter: getter,
	}

	// Retry logic encapsulated in reconnect so no need for it here.
	_, err := rrh.reconnect(numMaxRetries)

	return rrh, clues.Wrap(err, "initializing reader").OrNil()
}

//nolint:unused
type resetRetryHandler struct {
	ctx         context.Context
	getter      Getter
	innerReader io.ReadCloser
	offset      int64
}

func isRetriable(err error) bool {
	if err == nil {
		return false
	}

	for _, e := range retryErrs {
		if errors.Is(err, e) {
			return true
		}
	}

	return false
}

func (rrh *resetRetryHandler) Read(p []byte) (int, error) {
	if rrh.innerReader == nil {
		return 0, clues.New("not initialized")
	}

	var (
		// Use separate error variable just to make other assignments in the loop a
		// bit cleaner.
		finalErr   error
		read       int
		numRetries int
	)

	// Still need to check retry count in loop header so we don't go through one
	// last time after failing to reconnect due to exhausting retries.
	for numRetries < numMaxRetries {
		n, err := rrh.innerReader.Read(p[read:])
		rrh.offset = rrh.offset + int64(n)
		read = read + n

		// Catch short reads with no error and errors we don't know how to retry.
		if !isRetriable(err) {
			// Not everything knows how to handle a wrapped version of EOF (including
			// io.ReadAll) so return the error itself here.
			if errors.Is(err, io.EOF) {
				// Log info about the error, but only if it's not directly an EOF.
				// Otherwise this can be rather chatty and annoying to filter out.
				if err != io.EOF {
					logger.CtxErr(rrh.ctx, err).Debug("dropping wrapped io.EOF")
				}

				return read, io.EOF
			}

			return read, clues.Stack(err).WithClues(rrh.ctx).OrNil()
		}

		logger.Ctx(rrh.ctx).Infow(
			"restarting reader",
			"supports_range", rrh.getter.SupportsRange(),
			"restart_at_offset", rrh.offset,
			"retries_remaining", numMaxRetries-numRetries,
			"retriable_error", err)

		attempts, err := rrh.reconnect(numMaxRetries - numRetries)
		numRetries = numRetries + attempts
		finalErr = err
	}

	// We couln't read anything through all the retries but never had an error
	// getting another reader. Report this as an error so we don't get stuck in an
	// infinite loop.
	if read == 0 && finalErr == nil && numRetries >= numMaxRetries {
		finalErr = clues.Wrap(io.ErrNoProgress, "unable to read data")
	}

	return read, clues.Stack(finalErr).OrNil()
}

// reconnect attempts to get another instance of the underlying reader and set
// the reader to pickup where the previous reader left off.
//
// Since this function can be called by functions that also implement retries on
// read errors pass an int in to denote how many times to attempt to reconnect.
// This avoids mulplicative retries when called from other functions.
func (rrh *resetRetryHandler) reconnect(maxRetries int) (int, error) {
	var (
		attempts int
		skip     = rrh.offset
		headers  = map[string]string{}
		// This is annoying but we want the equivalent of a do-while loop.
		err = retryErrs[0]
	)

	if rrh.getter.SupportsRange() {
		headers[rangeHeaderKey] = fmt.Sprintf(
			rangeHeaderOneSidedValueTmpl,
			rrh.offset)
		skip = 0
	}

	ctx := clues.Add(
		rrh.ctx,
		"supports_range", rrh.getter.SupportsRange(),
		"restart_at_offset", rrh.offset)

	for attempts < maxRetries && isRetriable(err) {
		// Attempts will be 0 the first time through so it won't sleep then.
		time.Sleep(time.Duration(attempts*minSleepTime) * time.Second)

		attempts++

		var r io.ReadCloser

		r, err = rrh.getter.Get(ctx, headers)
		if err != nil {
			err = clues.Wrap(err, "retrying connection").
				WithClues(ctx).
				With("attempt_num", attempts)

			continue
		}

		if rrh.innerReader != nil {
			rrh.innerReader.Close()
		}

		rrh.innerReader = r

		// If we can't request a specific range of content then read as many bytes
		// as we've already processed into the equivalent of /dev/null so that the
		// next read will get content we haven't seen before.
		if skip > 0 {
			_, err = io.CopyN(io.Discard, rrh.innerReader, skip)
			if err != nil {
				err = clues.Wrap(err, "seeking to correct offset").
					WithClues(ctx).
					With("attempt_num", attempts)
			}
		}
	}

	return attempts, err
}

func (rrh *resetRetryHandler) Close() error {
	err := rrh.innerReader.Close()
	rrh.innerReader = nil

	return clues.Stack(err).OrNil()
}
