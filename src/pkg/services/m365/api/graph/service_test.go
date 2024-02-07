package graph

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"syscall"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/internal/tester/tconfig"
	"github.com/alcionai/canario/src/pkg/account"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/errs/core"
	graphTD "github.com/alcionai/canario/src/pkg/services/m365/api/graph/testdata"
)

type GraphIntgSuite struct {
	tester.Suite
	fakeCredentials account.M365Config
	credentials     account.M365Config
}

func TestGraphIntgSuite(t *testing.T) {
	suite.Run(t, &GraphIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *GraphIntgSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	fakeAcct := tconfig.NewFakeM365Account(t)
	acct := tconfig.NewM365Account(t)

	m365, err := fakeAcct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.fakeCredentials = m365

	m365, err = acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.credentials = m365

	InitializeConcurrencyLimiter(ctx, false, 0)
}

func (suite *GraphIntgSuite) TestCreateAdapter() {
	t := suite.T()
	adpt, err := CreateAdapter(
		suite.fakeCredentials.AzureTenantID,
		suite.fakeCredentials.AzureClientID,
		suite.fakeCredentials.AzureClientSecret,
		count.New())

	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, adpt)

	aw := adpt.(*adapterWrap)
	assert.Equal(t, adapterRetryDelay, aw.retryDelay, "default retry delay")
}

func (suite *GraphIntgSuite) TestHTTPClient() {
	table := []struct {
		name        string
		opts        []Option
		check       func(*testing.T, *http.Client)
		checkConfig func(*testing.T, *clientConfig)
	}{
		{
			name: "no options",
			opts: []Option{},
			check: func(t *testing.T, c *http.Client) {
				assert.Equal(t, defaultHTTPClientTimeout, c.Timeout, "default timeout")
			},
			checkConfig: func(t *testing.T, c *clientConfig) {
				assert.Equal(t, defaultDelay, c.minDelay, "default delay")
				assert.Equal(t, defaultMaxRetries, c.maxRetries, "max retries")
				assert.Equal(t, defaultMaxRetries, c.maxConnectionRetries, "max connection retries")
			},
		},
		{
			name: "configured options",
			opts: []Option{
				NoTimeout(),
				MaxRetries(4),
				MaxConnectionRetries(2),
				MinimumBackoff(999 * time.Millisecond),
			},
			check: func(t *testing.T, c *http.Client) {
				// FIXME: Change to 0 one upstream issue is fixed
				assert.Equal(t, time.Duration(48*time.Hour), c.Timeout, "unlimited timeout")
			},
			checkConfig: func(t *testing.T, c *clientConfig) {
				assert.Equal(t, 999*time.Millisecond, c.minDelay, "minimum delay")
				assert.Equal(t, 4, c.maxRetries, "max retries")
				assert.Equal(t, 2, c.maxConnectionRetries, "max connection retries")
			},
		},
		{
			name: "below minimums",
			opts: []Option{
				NoTimeout(),
				MaxRetries(-1),
				MaxConnectionRetries(-1),
				MinimumBackoff(0),
			},
			check: func(t *testing.T, c *http.Client) {
				// FIXME: Change to 0 one upstream issue is fixed
				assert.Equal(t, time.Duration(48*time.Hour), c.Timeout, "unlimited timeout")
			},
			checkConfig: func(t *testing.T, c *clientConfig) {
				assert.Equal(t, 100*time.Millisecond, c.minDelay, "minimum delay")
				assert.Equal(t, 0, c.maxRetries, "max retries")
				assert.Equal(t, 0, c.maxConnectionRetries, "max connection retries")
			},
		},
		{
			name: "above maximums",
			opts: []Option{
				NoTimeout(),
				MaxRetries(9001),
				MaxConnectionRetries(9001),
				MinimumBackoff(999 * time.Second),
			},
			check: func(t *testing.T, c *http.Client) {
				// FIXME: Change to 0 one upstream issue is fixed
				assert.Equal(t, time.Duration(48*time.Hour), c.Timeout, "unlimited timeout")
			},
			checkConfig: func(t *testing.T, c *clientConfig) {
				assert.Equal(t, 5*time.Second, c.minDelay, "minimum delay")
				assert.Equal(t, 5, c.maxRetries, "max retries")
				assert.Equal(t, 5, c.maxConnectionRetries, "max connection retries")
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cli, cc := KiotaHTTPClient(count.New(), test.opts...)
			assert.NotNil(t, cli)
			test.check(t, cli)
			test.checkConfig(t, cc)
		})
	}
}

func (suite *GraphIntgSuite) TestSerializationEndPoint() {
	t := suite.T()
	adpt, err := CreateAdapter(
		suite.fakeCredentials.AzureTenantID,
		suite.fakeCredentials.AzureClientID,
		suite.fakeCredentials.AzureClientSecret,
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	serv := NewService(adpt)
	email := models.NewMessage()
	subject := "TestSerializationEndPoint"
	email.SetSubject(&subject)

	byteArray, err := serv.Serialize(email)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, byteArray)
	t.Log(string(byteArray))
}

func (suite *GraphIntgSuite) TestAdapterWrap_catchesPanic() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	url := "https://graph.microsoft.com/fnords/beaux/regard"

	// the panics should get caught and returned as errors
	alwaysPanicMiddleware := mwForceResp{
		alternate: func(req *http.Request) (bool, *http.Response, error) {
			panic(clues.New("intentional panic"))
		},
	}

	adpt, err := CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret,
		count.New(),
		appendMiddleware(&alwaysPanicMiddleware))
	require.NoError(t, err, clues.ToCore(err))

	// Set retry delay to a low value to reduce test runtime.
	aw := adpt.(*adapterWrap)
	aw.retryDelay = 10 * time.Millisecond

	// the query doesn't matter
	_, err = users.NewItemCalendarsItemEventsDeltaRequestBuilder(url, adpt).Get(ctx, nil)
	require.Error(t, err, clues.ToCore(err))
	require.Contains(t, err.Error(), "panic", clues.ToCore(err))

	// the query doesn't matter
	_, err = NewService(adpt).Client().Users().Get(ctx, nil)
	require.Error(t, err, clues.ToCore(err))
	require.Contains(t, err.Error(), "panic", clues.ToCore(err))
}

func (suite *GraphIntgSuite) TestAdapterWrap_retriesConnectionClose() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	retryInc := 0

	// the panics should get caught and returned as errors
	alwaysECONNRESET := mwForceResp{
		err: syscall.ECONNRESET,
		alternate: func(req *http.Request) (bool, *http.Response, error) {
			retryInc++
			return false, nil, nil
		},
	}

	adpt, err := CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret,
		count.New(),
		appendMiddleware(&alwaysECONNRESET),
		// Configure retry middlewares so that they don't retry on connection reset.
		// Those middlewares have their own tests to verify retries.
		MaxRetries(-1))
	require.NoError(t, err, clues.ToCore(err))

	// Retry delay doesn't really matter here since all requests will be intercepted
	// by the test middleware. Set it to 0 to reduce test runtime.
	aw := adpt.(*adapterWrap)
	aw.retryDelay = 0

	// the query doesn't matter
	_, err = NewService(adpt).Client().Users().Get(ctx, nil)
	require.ErrorIs(t, err, syscall.ECONNRESET, clues.ToCore(err))
	require.Equal(t, 4, retryInc, "number of retries")
}

func (suite *GraphIntgSuite) TestAdapterWrap_retriesBadJWTToken() {
	var (
		t        = suite.T()
		retryInc = 0
		odErr    = graphTD.ODataErr(string(invalidAuthenticationToken))
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	// the panics should get caught and returned as errors
	alwaysBadJWT := mwForceResp{
		alternate: func(req *http.Request) (bool, *http.Response, error) {
			retryInc++

			l, b := graphTD.ParseableToReader(t, odErr)

			header := http.Header{}
			header.Set("Content-Length", strconv.Itoa(int(l)))
			header.Set("Content-Type", "application/json")

			resp := &http.Response{
				Body:          b,
				ContentLength: l,
				Header:        header,
				Proto:         req.Proto,
				Request:       req,
				// avoiding 401 for the test to escape extraneous code paths in graph client
				// shouldn't affect the result
				StatusCode: http.StatusMethodNotAllowed,
			}

			return true, resp, nil
		},
	}

	adpt, err := CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret,
		count.New(),
		appendMiddleware(&alwaysBadJWT))
	require.NoError(t, err, clues.ToCore(err))

	// Retry delay doesn't really matter here since all requests will be intercepted
	// by the test middleware. Set it to 0 to reduce test runtime.
	aw := adpt.(*adapterWrap)
	aw.retryDelay = 0

	// When run locally this may fail. Not sure why it works in github but not locally.
	// Pester keepers if it bothers you.
	_, err = users.
		NewItemCalendarsItemEventsDeltaRequestBuilder("https://graph.microsoft.com/fnords/beaux/regard", adpt).
		Get(ctx, nil)
	assert.ErrorIs(t, err, core.ErrAuthTokenExpired, clues.ToCore(err))
	assert.Equal(t, 4, retryInc, "number of retries")

	retryInc = 0

	// the query doesn't matter
	_, err = NewService(adpt).Client().Users().Get(ctx, nil)
	assert.ErrorIs(t, err, core.ErrAuthTokenExpired, clues.ToCore(err))
	assert.Equal(t, 4, retryInc, "number of retries")
}

// TestAdapterWrap_retriesInvalidRequest tests adapter retries for graph 400
// invalidRequest errors. It also tests that retries are only done for GET
// requests.
func (suite *GraphIntgSuite) TestAdapterWrap_retriesInvalidRequest() {
	var (
		t        = suite.T()
		retryInc = 0
		// Formulate a graph error response which is parseable to odata error.
		graphResp = map[string]any{
			"error": map[string]any{
				"code":    invalidRequest,
				"message": "Invalid request",
				"innerError": map[string]any{
					"date":              "2024-01-01T18:00:00",
					"request-id":        "rid",
					"client-request-id": "cid",
				},
			},
		}
	)

	serialized, err := json.Marshal(graphResp)
	require.NoError(t, err, clues.ToCore(err))

	l := int64(len(serialized))

	// Set up a test middleware to always return a graph 400 invalidRequest
	// response.
	returnsGraphResp := mwForceResp{
		alternate: func(req *http.Request) (bool, *http.Response, error) {
			retryInc++

			header := http.Header{}
			header.Set("Content-Length", strconv.Itoa(int(l)))
			header.Set("Content-Type", "application/json")

			resp := &http.Response{
				Body:          io.NopCloser(bytes.NewReader(serialized)),
				ContentLength: l,
				Header:        header,
				Proto:         req.Proto,
				Request:       req,
				StatusCode:    http.StatusBadRequest, // Required retry condition
			}

			return true, resp, nil
		},
	}

	adpt, err := CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret,
		count.New(),
		appendMiddleware(&returnsGraphResp))
	require.NoError(t, err, clues.ToCore(err))

	// Retry delay doesn't really matter here since all requests will be intercepted
	// by the test middleware. Set it to 0 to reduce test runtime.
	aw := adpt.(*adapterWrap)
	aw.retryDelay = 0

	table := []struct {
		name            string
		apiRequest      func(t *testing.T, ctx context.Context) error
		expectedRetries int
	}{
		{
			name: "GET request, retried",
			apiRequest: func(t *testing.T, ctx context.Context) error {
				_, err = NewService(adpt).Client().Users().Get(ctx, nil)
				return err
			},
			expectedRetries: 4,
		},
		{
			name: "POST request, no retry",
			apiRequest: func(t *testing.T, ctx context.Context) error {
				u := models.NewUser()

				cfg := users.UsersRequestBuilderPostRequestConfiguration{}

				_, err = NewService(adpt).Client().Users().Post(ctx, u, &cfg)
				return err
			},
			expectedRetries: 1,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			tt := suite.T()
			retryInc = 0

			ctx, flush := tester.NewContext(tt)
			defer flush()

			err := test.apiRequest(tt, ctx)
			assert.True(tt, IsErrInvalidRequest(err), clues.ToCore(err))
			assert.Equal(tt, test.expectedRetries, retryInc, "number of retries")
		})
	}
}
