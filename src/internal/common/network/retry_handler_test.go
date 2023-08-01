package network_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"syscall"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/network"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

type readResp struct {
	read int
	err  error
}

type mockReader struct {
	r    io.Reader
	data []byte
	// Associate return values for Read with calls. Allows partial reads as well.
	// A value of nil in this slice means completing the request completely with
	// no errors (i.e. all bytes requested are returned or as many as possible and
	// EOF).
	resps     []*readResp
	callCount int
}

func (mr *mockReader) Read(p []byte) (int, error) {
	defer func() {
		mr.callCount++
	}()

	if mr.r == nil {
		mr.reset(0)
	}

	if len(mr.resps) == 0 {
		n, err := mr.r.Read(p)
		return n, clues.Stack(err).OrNil()
	}

	resp := mr.resps[mr.callCount]
	if resp == nil {
		n, err := mr.r.Read(p)
		return n, clues.Stack(err).OrNil()
	}

	if resp.read == 0 {
		return resp.read, clues.Stack(resp.err).OrNil()
	}

	n, err := mr.r.Read(p[:resp.read])

	if resp.err != nil {
		return n, clues.Stack(resp.err)
	}

	return n, clues.Stack(err).OrNil()
}

func (mr *mockReader) reset(n int) {
	mr.r = bytes.NewBuffer(mr.data[n:])
}

type getterResp struct {
	status int
	offset int
	err    error
}

type mockGetter struct {
	t *testing.T
	// We assume a single URL with possibly different headers for all calls.
	expectURL     string
	reader        *mockReader
	resps         []*getterResp
	expectHeaders []map[string]string
	callCount     int
}

func (mg *mockGetter) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	defer func() {
		mg.callCount++
	}()

	assert.Equal(mg.t, mg.expectURL, url)

	expectHeaders := map[string]string{}
	if len(mg.expectHeaders) > 0 {
		expectHeaders = mg.expectHeaders[mg.callCount]
	}

	if expectHeaders == nil {
		expectHeaders = map[string]string{}
	}

	assert.Equal(mg.t, expectHeaders, headers)

	resp := getterResp{}

	if len(mg.resps) > 0 {
		// Alright if we end up with the default for resp because we assume
		// resetting the reader and returning no error is the usual.
		resp = ptr.Val(mg.resps[mg.callCount])
	}

	if resp.status == 0 {
		resp.status = http.StatusOK
	}

	if resp.offset >= 0 {
		mg.t.Logf("resetting reader to offset %d\n", resp.offset)
		mg.reader.reset(resp.offset)
	}

	return &http.Response{
		StatusCode: resp.status,
		Body:       io.NopCloser(mg.reader),
	}, clues.Stack(resp.err).OrNil()
}

type ResetRetryHandlerUnitSuite struct {
	tester.Suite
}

func TestResetRetryHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &ResetRetryHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ResetRetryHandlerUnitSuite) TestResetRetryHandler() {
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	url := "https://www.corsobackup.io"
	// Pick a smaller read size so we can see how things will act if we have a
	// "chunked" set of data.
	readSize := 4

	totalReadCalls := (len(data) / readSize) + 1
	if len(data)%readSize != 0 {
		totalReadCalls++
	}

	table := []struct {
		name                string
		supportsRange       bool
		getterResps         []*getterResp
		getterExpectHeaders []map[string]string
		// First entry should represent the first read which is through the original
		// underlying reader.
		readerResps     []*readResp
		expectData      []byte
		hasErr          bool
		expectErr       error
		expectErrLabels []string
	}{
		{
			name:       "NoErrorsNeverCallsGetter",
			expectData: data,
		},
		{
			name: "OnlyFirstReadErrors NoRangeSupport",
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls)
				r[0] = &readResp{
					read: 0,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data,
		},
		{
			name:          "OnlyFirstReadErrors RangeSupport",
			supportsRange: true,
			getterExpectHeaders: []map[string]string{
				{"Range": "bytes=0-"},
			},
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls)
				r[0] = &readResp{
					read: 0,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data,
		},
		{
			name: "ErrorInMiddle NoRangeSupport",
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+1)
				r[3] = &readResp{
					read: 0,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data,
		},
		{
			name:          "ErrorInMiddle RangeSupport",
			supportsRange: true,
			getterResps:   []*getterResp{{offset: 12}},
			getterExpectHeaders: []map[string]string{
				{"Range": "bytes=12-"},
			},
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls)
				r[3] = &readResp{
					read: 0,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data,
		},
		{
			name: "MultipleErrorsInMiddle NoRangeSupport",
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+3)
				r[3] = &readResp{
					read: 0,
					err:  syscall.ECONNRESET,
				}
				r[7] = &readResp{
					read: 0,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data,
		},
		{
			name:          "MultipleErrorsInMiddle RangeSupport",
			supportsRange: true,
			getterResps: []*getterResp{
				{offset: 12},
				{offset: 20},
			},
			getterExpectHeaders: []map[string]string{
				{"Range": "bytes=12-"},
				{"Range": "bytes=20-"},
			},
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+1)
				r[3] = &readResp{
					read: 0,
					err:  syscall.ECONNRESET,
				}
				r[6] = &readResp{
					read: 0,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data,
		},
		{
			name: "ShortReadWithError NoRangeSupport",
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+2)
				r[3] = &readResp{
					read: readSize / 2,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data,
		},
		{
			name:          "ShortReadWithError RangeSupport",
			supportsRange: true,
			getterResps: []*getterResp{
				{offset: 14},
			},
			getterExpectHeaders: []map[string]string{
				{"Range": "bytes=14-"},
			},
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+1)
				r[3] = &readResp{
					read: readSize / 2,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data,
		},
		{
			name: "ErrorAtEndOfRead NoRangeSupport",
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+2)
				r[3] = &readResp{
					read: readSize,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data,
		},
		{
			name:          "ErrorAtEndOfRead RangeSupport",
			supportsRange: true,
			getterResps: []*getterResp{
				{offset: 16},
			},
			getterExpectHeaders: []map[string]string{
				{"Range": "bytes=16-"},
			},
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+1)
				r[3] = &readResp{
					read: readSize,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data,
		},
		{
			name: "UnexpectedError NoRangeSupport",
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+2)
				r[3] = &readResp{
					read: 0,
					err:  assert.AnError,
				}

				return r
			}(),
			expectData: data[:12],
			expectErr:  assert.AnError,
		},
		{
			name:          "UnexpectedError RangeSupport",
			supportsRange: true,
			getterResps: []*getterResp{
				{offset: 12},
			},
			getterExpectHeaders: []map[string]string{
				{"Range": "bytes=12-"},
			},
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+1)
				r[3] = &readResp{
					read: 0,
					err:  assert.AnError,
				}

				return r
			}(),
			expectData: data[:12],
			expectErr:  assert.AnError,
		},
		{
			name: "BadStatusCode NoRangeSupport",
			getterResps: []*getterResp{
				{status: http.StatusNotFound},
			},
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+2)
				r[3] = &readResp{
					read: 0,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data[:12],
			hasErr:     true,
			expectErrLabels: []string{
				graph.LabelStatus(http.StatusNotFound),
			},
		},
		{
			name:          "BadStatusCode RangeSupport",
			supportsRange: true,
			getterResps: []*getterResp{
				{status: http.StatusNotFound},
			},
			getterExpectHeaders: []map[string]string{
				{"Range": "bytes=12-"},
			},
			readerResps: func() []*readResp {
				r := make([]*readResp, totalReadCalls+1)
				r[3] = &readResp{
					read: 0,
					err:  syscall.ECONNRESET,
				}

				return r
			}(),
			expectData: data[:12],
			hasErr:     true,
			expectErrLabels: []string{
				graph.LabelStatus(http.StatusNotFound),
			},
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
				expectURL:     url,
				reader:        reader,
				resps:         test.getterResps,
				expectHeaders: test.getterExpectHeaders,
			}

			var (
				err     error
				n       int
				offset  int
				resData = make([]byte, len(data))
				rrh     = network.NewRetryResetHandler(
					ctx,
					getter,
					url,
					io.NopCloser(reader),
					test.supportsRange)
			)

			for err == nil && offset < len(data) {
				end := offset + readSize
				if end > len(data) {
					end = len(data)
				}

				n, err = rrh.Read(resData[offset:end])

				offset = offset + n
				t.Logf("read %d bytes\n", n)
			}

			assert.Equal(t, test.expectData, data[:offset])

			if !test.hasErr && test.expectErr == nil {
				assert.NoError(t, err, clues.ToCore(err))
				return
			} else if test.hasErr {
				// We got an error but can't check for the exact type.
				assert.Error(t, err)
			} else {
				assert.ErrorIs(t, err, test.expectErr, clues.ToCore(err))
			}

			for _, l := range test.expectErrLabels {
				assert.True(t, clues.HasLabel(err, l), "expecting label %s in error", l)
			}
		})
	}
}
