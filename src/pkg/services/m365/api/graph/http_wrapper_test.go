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
	"github.com/alcionai/corso/src/pkg/count"
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

	hw := NewHTTPWrapper(count.New())

	resp, err := hw.Request(
		ctx,
		http.MethodGet,
		"https://www.google.com",
		nil,
		nil,
		false)
	require.NoError(t, err, clues.ToCore(err))

	defer resp.Body.Close()

	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Test http wrapper config
	assert.Equal(t, httpWrapperRetryDelay, hw.retryDelay)
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

func (suite *HTTPWrapperIntgSuite) TestHTTPWrapper_Request_withAuth() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	a := tconfig.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	azureAuth, err := NewAzureAuth(m365)
	require.NoError(t, err, clues.ToCore(err))

	hw := NewHTTPWrapper(count.New(), AuthorizeRequester(azureAuth))

	// any request that requires authorization will do
	resp, err := hw.Request(
		ctx,
		http.MethodGet,
		"https://graph.microsoft.com/v1.0/users",
		nil,
		nil,
		true)
	require.NoError(t, err, clues.ToCore(err))

	defer resp.Body.Close()

	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// also validate that non-auth'd endpoints succeed
	resp, err = hw.Request(
		ctx,
		http.MethodGet,
		"https://www.google.com",
		nil,
		nil,
		true)
	require.NoError(t, err, clues.ToCore(err))

	defer resp.Body.Close()

	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

// ---------------------------------------------------------------------------
// unit
// ---------------------------------------------------------------------------

type HTTPWrapperUnitSuite struct {
	tester.Suite
}

func TestHTTPWrapperUnitSuite(t *testing.T) {
	suite.Run(t, &HTTPWrapperUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *HTTPWrapperUnitSuite) TestHTTPWrapper_Request_redirect() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	respHdr := http.Header{}
	respHdr.Set("Location", "localhost:99999999/smarfs")

	toResp := &http.Response{
		StatusCode: http.StatusFound,
		Header:     respHdr,
	}

	mwResp := mwForceResp{
		resp: toResp,
		alternate: func(req *http.Request) (bool, *http.Response, error) {
			if strings.HasSuffix(req.URL.String(), "smarfs") {
				assert.Equal(t, req.Header.Get("X-Test-Val"), "should-be-copied-to-redirect")
				return true, &http.Response{StatusCode: http.StatusOK}, nil
			}

			return false, nil, nil
		},
	}

	hw := NewHTTPWrapper(count.New(), appendMiddleware(&mwResp))

	resp, err := hw.Request(
		ctx,
		http.MethodGet,
		"https://graph.microsoft.com/fnords/beaux/regard",
		nil,
		map[string]string{"X-Test-Val": "should-be-copied-to-redirect"},
		false)
	require.NoError(t, err, clues.ToCore(err))

	defer resp.Body.Close()

	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func (suite *HTTPWrapperUnitSuite) TestHTTPWrapper_Request_http2StreamErrorRetries() {
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
				count.New(),
				appendMiddleware(&mwResp),
				MaxConnectionRetries(test.retries))

			// Configure retry delay to reduce test time. Retry delay doesn't
			// really matter here since all requests will be intercepted by
			// the test middleware.
			hw.retryDelay = 0

			_, err := hw.Request(ctx, http.MethodGet, url, nil, nil, false)
			require.ErrorAs(t, err, &http2.StreamError{}, clues.ToCore(err))
			require.Equal(t, test.expectRetries, tries, "count of retries")
		})
	}
}
