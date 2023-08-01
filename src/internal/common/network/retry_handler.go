package network

import (
	"context"
	"errors"
	"fmt"
	"io"
	"syscall"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ io.ReadCloser = &resetRetryHandler{}

const (
	rangeHeaderKey               = "Range"
	rangeHeaderOneSidedValueTmpl = "bytes=%d-"
)

func NewRetryResetHandler(
	ctx context.Context,
	getter api.Getter,
	url string,
	innerReader io.ReadCloser,
	supportsRangeReq bool,
) *resetRetryHandler {
	return &resetRetryHandler{
		ctx:              ctx,
		getter:           getter,
		url:              url,
		innerReader:      innerReader,
		supportsRangeReq: supportsRangeReq,
	}
}

//nolint:unused
type resetRetryHandler struct {
	ctx              context.Context
	getter           api.Getter
	url              string
	offset           int64
	innerReader      io.ReadCloser
	supportsRangeReq bool
	// TODO(ashmrtn): Add some way to get an updated URL since they can expire?
	// Unclear what the data consistency for that situation would be (i.e. could
	// the item's content change between the old URL and the new one such that
	// resuming the reads at the same offset with a new URL yields different data
	// than we would have gotten at the same offset from the old URL).
}

func (rrh *resetRetryHandler) Read(p []byte) (int, error) {
	if rrh.innerReader == nil {
		return 0, clues.New("not initialized")
	}

	n, err := rrh.innerReader.Read(p)
	rrh.offset = rrh.offset + int64(n)

	if err != nil {
		// Some other error we're not handling here. Will handle EOF errors in the
		// way we want as well.
		if !errors.Is(err, syscall.ECONNRESET) {
			return n, clues.Stack(err)
		}

		// Allow retry to reread as well just to simplify things.
		var n2 int

		n2, err = rrh.retry(p[n:])
		rrh.offset = rrh.offset + int64(n2)
		n = n + n2
	}

	return n, clues.Stack(err).OrNil()
}

func (rrh *resetRetryHandler) retry(p []byte) (int, error) {
	var (
		headers = map[string]string{}
		skip    = rrh.offset
	)

	if rrh.supportsRangeReq {
		headers[rangeHeaderKey] = fmt.Sprintf(
			rangeHeaderOneSidedValueTmpl,
			rrh.offset)
		skip = 0
	}

	resp, err := rrh.getter.Get(rrh.ctx, rrh.url, headers)
	if err != nil {
		return 0, clues.Wrap(err, "retrying connection")
	}

	if resp != nil && (resp.StatusCode/100) != 2 {
		// TODO(ashmrtn): Labeling should really be done in a different layer so
		// this can be service-agnostic.
		return 0, clues.Wrap(clues.New(resp.Status), "retrying connection").
			Label(graph.LabelStatus(resp.StatusCode))
	}

	rrh.innerReader = resp.Body

	// If we can't request a specific range of content then read as many bytes as
	// we've already processed into the equivalent of /dev/null so that the next
	// read will get content we haven't seen before.
	if skip > 0 {
		if _, err := io.CopyN(io.Discard, rrh.innerReader, skip); err != nil {
			return 0, clues.Wrap(err, "seeking to correct offset")
		}
	}

	n, err := rrh.innerReader.Read(p)

	return n, clues.Stack(err).OrNil()
}

func (rrh *resetRetryHandler) Close() error {
	err := rrh.innerReader.Close()
	rrh.innerReader = nil

	return clues.Stack(err).OrNil()
}
