package readers

import (
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"syscall"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type readResp struct {
	read int
	// sticky denotes whether the error should continue to be returned until reset
	// is called.
	sticky bool
	err    error
}

type mockReader struct {
	r    io.Reader
	data []byte
	// Associate return values for Read with calls. Allows partial reads as well.
	// If a value for a particular read call is not in the map that means
	// completing the request completely with no errors (i.e. all bytes requested
	// are returned or as many as possible and EOF).
	resps     map[int]readResp
	callCount int
	stickyErr error
}

func (mr *mockReader) Read(p []byte) (int, error) {
	defer func() {
		mr.callCount++
	}()

	if mr.r == nil {
		mr.reset(0)
	}

	if mr.stickyErr != nil {
		return 0, clues.Wrap(mr.stickyErr, "sticky error")
	}

	resp, ok := mr.resps[mr.callCount]
	if !ok {
		n, err := mr.r.Read(p)
		return n, clues.Stack(err).OrNil()
	}

	n, err := mr.r.Read(p[:resp.read])

	if resp.err != nil {
		if resp.sticky {
			mr.stickyErr = resp.err
		}

		return n, clues.Stack(resp.err)
	}

	return n, clues.Stack(err).OrNil()
}

func (mr *mockReader) reset(n int) {
	mr.r = bytes.NewBuffer(mr.data[n:])
	mr.stickyErr = nil
}

type getterResp struct {
	offset int
	err    error
}

type mockGetter struct {
	t             *testing.T
	supportsRange bool
	reader        *mockReader
	resps         map[int]getterResp
	expectHeaders map[int]map[string]string
	callCount     int
}

func (mg *mockGetter) SupportsRange() bool {
	return mg.supportsRange
}

func (mg *mockGetter) Get(
	ctx context.Context,
	headers map[string]string,
) (io.ReadCloser, error) {
	defer func() {
		mg.callCount++
	}()

	expectHeaders := mg.expectHeaders[mg.callCount]
	if expectHeaders == nil {
		expectHeaders = map[string]string{}
	}

	assert.Equal(mg.t, expectHeaders, headers)

	resp := mg.resps[mg.callCount]

	if resp.offset >= 0 {
		mg.reader.reset(resp.offset)
	}

	return io.NopCloser(mg.reader), clues.Stack(resp.err).OrNil()
}

type ResetRetryHandlerUnitSuite struct {
	tester.Suite
}

func TestResetRetryHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &ResetRetryHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ResetRetryHandlerUnitSuite) TestResetRetryHandler() {
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	// Pick a smaller read size so we can see how things will act if we have a
	// "chunked" set of data.
	readSize := 4

	table := []struct {
		name          string
		supportsRange bool
		// 0th entry is the return data when trying to initialize the wrapper.
		getterResps map[int]getterResp
		// 0th entry is the return data when trying to initialize the wrapper.
		getterExpectHeaders map[int]map[string]string
		readerResps         map[int]readResp
		expectData          []byte
		expectErr           error
	}{
		{
			name: "OnlyFirstGetErrors ECONNRESET NoRangeSupport",
			getterResps: map[int]getterResp{
				0: {
					err: syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name: "OnlyFirstGetErrors ErrUnexpectedEOF NoRangeSupport",
			getterResps: map[int]getterResp{
				0: {
					err: io.ErrUnexpectedEOF,
				},
			},
			expectData: data,
		},
		{
			name:          "OnlyFirstReadErrors RangeSupport",
			supportsRange: true,
			getterResps: map[int]getterResp{
				0: {
					err: syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name: "ErrorInMiddle NoRangeSupport",
			readerResps: map[int]readResp{
				3: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name:          "ErrorInMiddle RangeSupport",
			supportsRange: true,
			getterResps: map[int]getterResp{
				1: {offset: 12},
			},
			getterExpectHeaders: map[int]map[string]string{
				1: {"Range": "bytes=12-"},
			},
			readerResps: map[int]readResp{
				3: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name: "MultipleErrorsInMiddle NoRangeSupport",
			readerResps: map[int]readResp{
				3: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
				7: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name:          "MultipleErrorsInMiddle RangeSupport",
			supportsRange: true,
			getterResps: map[int]getterResp{
				1: {offset: 12},
				2: {offset: 20},
			},
			getterExpectHeaders: map[int]map[string]string{
				1: {"Range": "bytes=12-"},
				2: {"Range": "bytes=20-"},
			},
			readerResps: map[int]readResp{
				3: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
				6: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name:          "MultipleRetriableErrorTypesInMiddle RangeSupport",
			supportsRange: true,
			getterResps: map[int]getterResp{
				1: {offset: 12},
				2: {offset: 20},
			},
			getterExpectHeaders: map[int]map[string]string{
				1: {"Range": "bytes=12-"},
				2: {"Range": "bytes=20-"},
			},
			readerResps: map[int]readResp{
				3: {
					read: 0,
					err:  io.ErrUnexpectedEOF,
				},
				6: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name: "ShortReadWithError NoRangeSupport",
			readerResps: map[int]readResp{
				3: {
					read: readSize / 2,
					err:  syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name:          "ShortReadWithError RangeSupport",
			supportsRange: true,
			getterResps: map[int]getterResp{
				1: {offset: 14},
			},
			getterExpectHeaders: map[int]map[string]string{
				1: {"Range": "bytes=14-"},
			},
			readerResps: map[int]readResp{
				3: {
					read: readSize / 2,
					err:  syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name: "ErrorAtEndOfRead NoRangeSupport",
			readerResps: map[int]readResp{
				3: {
					read:   readSize,
					sticky: true,
					err:    syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name:          "ErrorAtEndOfRead RangeSupport",
			supportsRange: true,
			getterResps: map[int]getterResp{
				1: {offset: 16},
			},
			getterExpectHeaders: map[int]map[string]string{
				1: {"Range": "bytes=16-"},
			},
			readerResps: map[int]readResp{
				3: {
					read:   readSize,
					sticky: true,
					err:    syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name: "UnexpectedError NoRangeSupport",
			readerResps: map[int]readResp{
				3: {
					read: 0,
					err:  assert.AnError,
				},
			},
			expectData: data[:12],
			expectErr:  assert.AnError,
		},
		{
			name:          "UnexpectedError RangeSupport",
			supportsRange: true,
			getterResps: map[int]getterResp{
				1: {offset: 12},
			},
			getterExpectHeaders: map[int]map[string]string{
				1: {"Range": "bytes=12-"},
			},
			readerResps: map[int]readResp{
				3: {
					read: 0,
					err:  assert.AnError,
				},
			},
			expectData: data[:12],
			expectErr:  assert.AnError,
		},
		{
			name: "ErrorWhileSeeking NoRangeSupport",
			readerResps: map[int]readResp{
				3: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
				4: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name: "ShortReadNoError NoRangeSupport",
			readerResps: map[int]readResp{
				3: {
					read: readSize / 2,
				},
			},
			expectData: data,
		},
		{
			name:          "ShortReadNoError RangeSupport",
			supportsRange: true,
			getterResps: map[int]getterResp{
				1: {offset: 14},
			},
			getterExpectHeaders: map[int]map[string]string{
				1: {"Range": "bytes=14-"},
			},
			readerResps: map[int]readResp{
				3: {
					read: readSize / 2,
				},
			},
			expectData: data,
		},
		{
			name: "TooManyRetriesDuringRead NoRangeSupport",
			// Fail the final reconnect attempt so we run out of retries. Otherwise we
			// exit with a short read and successful reconnect.
			getterResps: map[int]getterResp{
				3: {err: syscall.ECONNRESET},
			},
			// Even numbered read requests are seeks to the proper offset.
			readerResps: map[int]readResp{
				3: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
				5: {
					read: 1,
					err:  syscall.ECONNRESET,
				},
				7: {
					read: 1,
					err:  syscall.ECONNRESET,
				},
			},
			expectData: data[:14],
			expectErr:  syscall.ECONNRESET,
		},
		{
			name:          "TooManyRetriesDuringRead RangeSupport",
			supportsRange: true,
			getterResps: map[int]getterResp{
				1: {offset: 12},
				2: {offset: 12},
				3: {err: syscall.ECONNRESET},
			},
			getterExpectHeaders: map[int]map[string]string{
				1: {"Range": "bytes=12-"},
				2: {"Range": "bytes=13-"},
				3: {"Range": "bytes=14-"},
			},
			readerResps: map[int]readResp{
				3: {
					read: 0,
					err:  syscall.ECONNRESET,
				},
				4: {
					read: 1,
					err:  syscall.ECONNRESET,
				},
				5: {
					read: 1,
					err:  syscall.ECONNRESET,
				},
			},
			expectData: data[:14],
			expectErr:  syscall.ECONNRESET,
		},
		{
			name:          "TooManyRetriesDuringRead AlwaysReturnError RangeSupport",
			supportsRange: true,
			getterResps: map[int]getterResp{
				1: {offset: -1},
				2: {offset: -1},
				3: {offset: -1},
				4: {offset: -1},
				5: {offset: -1},
			},
			readerResps: map[int]readResp{
				0: {
					sticky: true,
					err:    syscall.ECONNRESET,
				},
			},
			expectData: []byte{},
			expectErr:  io.ErrNoProgress,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			reader := &mockReader{
				data:  data,
				resps: test.readerResps,
			}

			getter := &mockGetter{
				t:             t,
				supportsRange: test.supportsRange,
				reader:        reader,
				resps:         test.getterResps,
				expectHeaders: test.getterExpectHeaders,
			}

			var (
				err     error
				n       int
				offset  int
				resData = make([]byte, len(data))
			)

			rrh, err := NewResetRetryHandler(ctx, getter)
			require.NoError(t, err, "making reader wrapper: %v", clues.ToCore(err))

			for err == nil && offset < len(data) {
				end := offset + readSize
				if end > len(data) {
					end = len(data)
				}

				n, err = rrh.Read(resData[offset:end])

				offset = offset + n
			}

			assert.Equal(t, test.expectData, data[:offset])

			if test.expectErr == nil {
				assert.NoError(t, err, clues.ToCore(err))
				return
			}

			assert.ErrorIs(t, err, test.expectErr, clues.ToCore(err))
		})
	}
}

func (suite *ResetRetryHandlerUnitSuite) TestIsRetriable() {
	table := []struct {
		name   string
		err    func() error
		expect bool
	}{
		{
			name:   "nil",
			err:    func() error { return nil },
			expect: false,
		},
		{
			name:   "Connection Reset Error",
			err:    func() error { return syscall.ECONNRESET },
			expect: true,
		},
		{
			name:   "Unexpected EOF Error",
			err:    func() error { return io.ErrUnexpectedEOF },
			expect: true,
		},
		{
			name:   "Not Retriable Error",
			err:    func() error { return assert.AnError },
			expect: false,
		},
		{
			name: "Chained Errors With No Retriables",
			err: func() error {
				return clues.Stack(assert.AnError, clues.New("another error"))
			},
			expect: false,
		},
		{
			name: "Chained Errors With ECONNRESET",
			err: func() error {
				return clues.Stack(assert.AnError, syscall.ECONNRESET, assert.AnError)
			},
			expect: true,
		},
		{
			name: "Chained Errors With ErrUnexpectedEOF",
			err: func() error {
				return clues.Stack(assert.AnError, io.ErrUnexpectedEOF, assert.AnError)
			},
			expect: true,
		},
		{
			name: "Wrapped ECONNRESET Error",
			err: func() error {
				return clues.Wrap(syscall.ECONNRESET, "wrapped error")
			},
			expect: true,
		},
		{
			name: "Wrapped ErrUnexpectedEOF Error",
			err: func() error {
				return clues.Wrap(io.ErrUnexpectedEOF, "wrapped error")
			},
			expect: true,
		},
		{
			name:   "Timeout - deadline exceeded",
			err:    func() error { return context.DeadlineExceeded },
			expect: true,
		},
		{
			name:   "Timeout - ctx canceled",
			err:    func() error { return context.Canceled },
			expect: true,
		},
		{
			name:   "Timeout - http timeout",
			err:    func() error { return http.ErrHandlerTimeout },
			expect: true,
		},
		{
			name: "Timeout - url error",
			err: func() error {
				return &url.Error{Err: context.DeadlineExceeded}
			},
			expect: true,
		},
		{
			name: "Timeout - OS timeout",
			err: func() error {
				return &os.PathError{Err: os.ErrDeadlineExceeded}
			},
			expect: true,
		},
		{
			name: "Timeout - net timeout",
			err: func() error {
				return &net.OpError{
					Op:  "read",
					Err: &os.PathError{Err: os.ErrDeadlineExceeded},
				}
			},
			expect: true,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, test.expect, isRetriable(test.err()))
		})
	}
}
