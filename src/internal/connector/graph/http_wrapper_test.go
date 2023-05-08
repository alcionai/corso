package graph

import (
	"net/http"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	khttp "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type HTTPWrapperIntgSuite struct {
	tester.Suite
}

func TestHTTPWrapperIntgSuite(t *testing.T) {
	suite.Run(t, &HTTPWrapperIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *HTTPWrapperIntgSuite) TestNewHTTPWrapper() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t  = suite.T()
		hw = NewHTTPWrapper()
	)

	resp, err := hw.Request(
		ctx,
		http.MethodGet,
		"https://www.corsobackup.io",
		nil,
		nil)

	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

type mwForceResp struct {
	err       error
	resp      *http.Response
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
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		uri  = "https://graph.microsoft.com"
		path = "/fnords/beaux/regard"
		url  = uri + path
	)

	// can't use gock for this, or else it'll short-circut the transport,
	// and thus skip all the middelware
	hdr := http.Header{}
	hdr.Set("Location", "localhost:99999999/smarfs")
	toResp := &http.Response{
		StatusCode: 302,
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
	require.NotNil(t, resp)
	// require.Equal(t, 1, calledCorrectly, "test server was called with expected path")
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
