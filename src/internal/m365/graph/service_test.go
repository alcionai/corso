package graph

import (
	"net/http"
	"syscall"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
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
		suite.fakeCredentials.AzureClientSecret)

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
				assert.Equal(t, defaultHTTPClientTimeout, c.Timeout, "default timeout")
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
				assert.Equal(t, defaultHTTPClientTimeout, c.Timeout, "default timeout")
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

			cli, cc := KiotaHTTPClient(test.opts...)
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
		suite.fakeCredentials.AzureClientSecret)
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
	count := 0

	// the panics should get caught and returned as errors
	alwaysECONNRESET := mwForceResp{
		err: syscall.ECONNRESET,
		alternate: func(req *http.Request) (bool, *http.Response, error) {
			count++
			return false, nil, nil
		},
	}

	adpt, err := CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret,
		appendMiddleware(&alwaysECONNRESET))
	require.NoError(t, err, clues.ToCore(err))

	// the query doesn't matter
	_, err = users.NewItemCalendarsItemEventsDeltaRequestBuilder(url, adpt).Get(ctx, nil)
	require.ErrorIs(t, err, syscall.ECONNRESET, clues.ToCore(err))
	require.Equal(t, 16, count, "number of retries")

	count = 0

	// the query doesn't matter
	_, err = NewService(adpt).Client().Users().Get(ctx, nil)
	require.ErrorIs(t, err, syscall.ECONNRESET, clues.ToCore(err))
	require.Equal(t, 16, count, "number of retries")
}
