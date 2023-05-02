package graph

import (
	"net/http"
	"testing"
	"time"

	"github.com/alcionai/clues"
	khttp "github.com/microsoft/kiota-http-go"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

func newBodylessTestMW(onIntercept func(), code int, err error) testMW {
	return testMW{
		err:         err,
		onIntercept: onIntercept,
		resp:        &http.Response{StatusCode: code},
	}
}

type testMW struct {
	err         error
	onIntercept func()
	resp        *http.Response
}

func (mw testMW) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	mw.onIntercept()
	return mw.resp, mw.err
}

// can't use graph/mock.CreateAdapter() due to circular references.
func mockAdapter(creds account.M365Config, mw khttp.Middleware) (*msgraphsdkgo.GraphRequestAdapter, error) {
	auth, err := GetAuth(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret)
	if err != nil {
		return nil, err
	}

	var (
		clientOptions = msgraphsdkgo.GetDefaultClientOptions()
		cc            = populateConfig(MinimumBackoff(10 * time.Millisecond))
		middlewares   = append(kiotaMiddlewares(&clientOptions, cc), mw)
		httpClient    = msgraphgocore.GetDefaultClient(&clientOptions, middlewares...)
	)

	httpClient.Timeout = 5 * time.Second

	cc.apply(httpClient)

	return msgraphsdkgo.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		auth,
		nil, nil,
		httpClient)
}

type RetryMWIntgSuite struct {
	tester.Suite
	creds account.M365Config
	srv   Servicer
	user  string
}

// We do end up mocking the actual request, but creating the rest
// similar to E2E suite
func TestRetryMWIntgSuite(t *testing.T) {
	suite.Run(t, &RetryMWIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *RetryMWIntgSuite) SetupSuite() {
	var (
		a   = tester.NewM365Account(suite.T())
		err error
	)

	suite.creds, err = a.M365Config()
	require.NoError(suite.T(), err, clues.ToCore(err))
}

func (suite *RetryMWIntgSuite) TestRetryMiddleware_Intercept_byStatusCode() {
	var (
		uri  = "https://graph.microsoft.com"
		path = "/v1.0/users/user/messages/foo"
		url  = uri + path
	)

	tests := []struct {
		name             string
		status           int
		expectRetryCount int
		mw               testMW
		expectErr        assert.ErrorAssertionFunc
	}{
		{
			name:             "200, no retries",
			status:           http.StatusOK,
			expectRetryCount: 0,
			expectErr:        assert.NoError,
		},
		{
			name:             "400, no retries",
			status:           http.StatusBadRequest,
			expectRetryCount: 0,
			expectErr:        assert.Error,
		},
		{
			// don't test 504: gets intercepted by graph client for long waits.
			name:             "502",
			status:           http.StatusBadGateway,
			expectRetryCount: defaultMaxRetries,
			expectErr:        assert.Error,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()
			called := 0
			mw := newBodylessTestMW(func() { called++ }, test.status, nil)

			adpt, err := mockAdapter(suite.creds, mw)
			require.NoError(t, err, clues.ToCore(err))

			// url doesn't fit the builder, but that shouldn't matter
			_, err = users.NewCountRequestBuilder(url, adpt).Get(ctx, nil)
			test.expectErr(t, err, clues.ToCore(err))

			// -1 because the non-retried call always counts for one, then
			// we increment based on the number of retry attempts.
			assert.Equal(t, test.expectRetryCount, called-1)
		})
	}
}
