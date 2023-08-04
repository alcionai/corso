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
		name  string
		opts  []Option
		check func(*testing.T, *http.Client)
	}{
		{
			name: "no options",
			opts: []Option{},
			check: func(t *testing.T, c *http.Client) {
				assert.Equal(t, defaultHTTPClientTimeout, c.Timeout, "default timeout")
			},
		},
		{
			name: "no timeout",
			opts: []Option{NoTimeout()},
			check: func(t *testing.T, c *http.Client) {
				// FIXME: Change to 0 one upstream issue is fixed
				assert.Equal(t, time.Duration(48*time.Hour), c.Timeout, "unlimited timeout")
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cli := KiotaHTTPClient(test.opts...)
			assert.NotNil(t, cli)
			test.check(t, cli)
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
	require.Equal(t, 12, count, "number of retries")

	count = 0

	// the query doesn't matter
	_, err = NewService(adpt).Client().Users().Get(ctx, nil)
	require.ErrorIs(t, err, syscall.ECONNRESET, clues.ToCore(err))
	require.Equal(t, 12, count, "number of retries")
}
