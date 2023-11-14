package graph

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"syscall"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/count"
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

	url := "https://graph.microsoft.com/fnords/beaux/regard"
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
		appendMiddleware(&alwaysECONNRESET))
	require.NoError(t, err, clues.ToCore(err))

	// the query doesn't matter
	_, err = users.NewItemCalendarsItemEventsDeltaRequestBuilder(url, adpt).Get(ctx, nil)
	require.ErrorIs(t, err, syscall.ECONNRESET, clues.ToCore(err))
	require.Equal(t, 16, retryInc, "number of retries")

	retryInc = 0

	// the query doesn't matter
	_, err = NewService(adpt).Client().Users().Get(ctx, nil)
	require.ErrorIs(t, err, syscall.ECONNRESET, clues.ToCore(err))
	require.Equal(t, 16, retryInc, "number of retries")
}

func requireParseableToReader(t *testing.T, thing serialization.Parsable) (int64, io.ReadCloser) {
	sw := kjson.NewJsonSerializationWriter()

	err := sw.WriteObjectValue("", thing)
	require.NoError(t, err, "serialize")

	content, err := sw.GetSerializedContent()
	require.NoError(t, err, "deserialize")

	return int64(len(content)), io.NopCloser(bytes.NewReader(content))
}

func (suite *GraphIntgSuite) TestAdapterWrap_retriesBadJWTToken() {
	var (
		t        = suite.T()
		retryInc = 0
		odErr    = odErrMsg(string(invalidAuthenticationToken), "string(invalidAuthenticationToken)")
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	// the panics should get caught and returned as errors
	alwaysBadJWT := mwForceResp{
		alternate: func(req *http.Request) (bool, *http.Response, error) {
			retryInc++

			l, b := requireParseableToReader(t, odErr)

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

	// Keeping this test in place for now as a showcase, even though
	// it kinda proves a failure to handle the error case.  On the bright
	// side, direct URL lookups are a rarity in corso, so the importance
	// is not as high.
	_, err = users.
		NewItemCalendarsItemEventsDeltaRequestBuilder("https://graph.microsoft.com/fnords/beaux/regard", adpt).
		Get(ctx, nil)
	assert.ErrorContains(
		t,
		err,
		"content type application/json does not have a factory registered to be parsed",
		clues.ToCore(err))
	assert.Equal(t, 1, retryInc, "number of retries")

	retryInc = 0

	// the query doesn't matter
	_, err = NewService(adpt).Client().Users().Get(ctx, nil)
	assert.True(t, IsErrBadJWTToken(err), clues.ToCore(err))
	assert.Equal(t, 4, retryInc, "number of retries")
}
