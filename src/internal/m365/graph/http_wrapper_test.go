package graph

import (
	"net/http"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	khttp "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/http2"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
)

type HTTPWrapperIntgSuite struct {
	tester.Suite
}

func TestHTTPWrapperIntgSuite(t *testing.T) {
	suite.Run(t, &HTTPWrapperIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *HTTPWrapperIntgSuite) TestNewHTTPWrapper() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	hw := NewHTTPWrapper()

	resp, err := hw.Request(
		ctx,
		http.MethodGet,
		"https://www.corsobackup.io",
		nil,
		nil)
	require.NoError(t, err, clues.ToCore(err))

	defer resp.Body.Close()

	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

type mwForceResp struct {
	err  error
	resp *http.Response
	// if alternate returns true, the middleware returns the
	// response and error returned by the func instead of the
	// resp and error saved in the struct.
	alternate func(*http.Request) (bool, *http.Response, error)
}

func (mw *mwForceResp) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	ok, r, e := mw.alternate(req)
	if ok {
		return r, e
	}

	return mw.resp, mw.err
}

type HTTPWrapperUnitSuite struct {
	tester.Suite
}

func TestHTTPWrapperUnitSuite(t *testing.T) {
	suite.Run(t, &HTTPWrapperUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *HTTPWrapperUnitSuite) TestNewHTTPWrapper_redirectMiddleware() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	url := "https://graph.microsoft.com/fnords/beaux/regard"

	hdr := http.Header{}
	hdr.Set("Location", "localhost:99999999/smarfs")

	toResp := &http.Response{
		StatusCode: http.StatusFound,
		Header:     hdr,
	}

	mwResp := mwForceResp{
		resp: toResp,
		alternate: func(req *http.Request) (bool, *http.Response, error) {
			if strings.HasSuffix(req.URL.String(), "smarfs") {
				return true, &http.Response{StatusCode: http.StatusOK}, nil
			}

			return false, nil, nil
		},
	}

	hw := NewHTTPWrapper(appendMiddleware(&mwResp))

	resp, err := hw.Request(ctx, http.MethodGet, url, nil, nil)
	require.NoError(t, err, clues.ToCore(err))

	defer resp.Body.Close()

	require.NotNil(t, resp)
	// require.Equal(t, 1, calledCorrectly, "test server was called with expected path")
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func (suite *HTTPWrapperUnitSuite) TestNewHTTPWrapper_http2StreamErrorRetries() {
	var (
		url       = "https://graph.microsoft.com/fnords/beaux/regard"
		streamErr = http2.StreamError{
			StreamID: 1,
			Code:     http2.ErrCodeEnhanceYourCalm,
			Cause:    assert.AnError,
		}
	)

	table := []struct {
		name          string
		retries       int
		expectRetries int
	}{
		{
			name:          "zero retries",
			retries:       0,
			expectRetries: 0,
		},
		{
			name:          "negative max",
			retries:       -1,
			expectRetries: 0,
		},
		{
			name:          "upper limit",
			retries:       9001,
			expectRetries: 5,
		},
		{
			name:          "four",
			retries:       4,
			expectRetries: 4,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			// -1 to account for the first try,
			// which isn't a retry.
			tries := -1

			mwResp := mwForceResp{
				err: streamErr,
				alternate: func(*http.Request) (bool, *http.Response, error) {
					tries++
					return false, nil, nil
				},
			}

			hw := NewHTTPWrapper(
				appendMiddleware(&mwResp),
				MaxConnectionRetries(test.retries))

			_, err := hw.Request(ctx, http.MethodGet, url, nil, nil)
			require.ErrorAs(t, err, &http2.StreamError{}, clues.ToCore(err))

			require.Equal(t, test.expectRetries, tries, "count of retries")
		})
	}
}
