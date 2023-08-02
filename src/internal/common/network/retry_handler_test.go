package network_test

import (
	"bytes"
	"context"
	"io"
	"syscall"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/network"
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

func (mg *mockGetter) SupportsRangeReq() bool {
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
			name: "OnlyFirstGetErrors NoRangeSupport",
			getterResps: map[int]getterResp{
				0: {
					err: syscall.ECONNRESET,
				},
			},
			expectData: data,
		},
		{
			name:          "OnlyFirstReadErrors RangeSupport",
			supportsRange: true,
			getterExpectHeaders: map[int]map[string]string{
				0: {"Range": "bytes=0-"},
				1: {"Range": "bytes=0-"},
			},
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
				0: {"Range": "bytes=0-"},
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
				0: {"Range": "bytes=0-"},
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
				0: {"Range": "bytes=0-"},
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
				0: {"Range": "bytes=0-"},
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
				0: {"Range": "bytes=0-"},
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
			getterResps: map[int]getterResp{
				1: {err: syscall.ECONNRESET},
			},
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
				0: {"Range": "bytes=0-"},
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
				0: {"Range": "bytes=0-"},
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
			getterExpectHeaders: map[int]map[string]string{
				0: {"Range": "bytes=0-"},
				1: {"Range": "bytes=0-"},
				2: {"Range": "bytes=0-"},
				3: {"Range": "bytes=0-"},
				4: {"Range": "bytes=0-"},
				5: {"Range": "bytes=0-"},
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

			rrh, err := network.NewResetRetryHandler(ctx, getter)
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
