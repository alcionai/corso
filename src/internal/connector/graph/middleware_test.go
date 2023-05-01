package graph

import (
	"net/http"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type RetryMWIntgSuite struct {
	tester.Suite
	credentials account.M365Config
	srv         Servicer
	user        string
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
	t := suite.T()

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.credentials = m365

	adp, err := CreateAdapter(
		m365.AzureTenantID,
		m365.AzureClientID,
		m365.AzureClientSecret,
		MinimumBackoff(10*time.Millisecond))
	require.NoError(t, err)

	suite.srv = NewService(adp)
}

func (suite *RetryMWIntgSuite) TestRetryMiddleware_Intercept_byStatusCode() {
	var (
		uri  = "https://graph.microsoft.com"
		path = "/v1.0/users/user/messages/foo"
	)

	tests := []struct {
		name       string
		status     int
		retryCount int
	}{
		{
			name:       "400, no retries",
			status:     http.StatusBadRequest,
			retryCount: 0,
		},
		{
			name:       "502",
			status:     http.StatusBadGateway,
			retryCount: defaultMaxRetries,
		},
		{
			name:       "504",
			status:     http.StatusGatewayTimeout,
			retryCount: defaultMaxRetries,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()

			defer gock.Off()
			gock.New(uri).Get(path).Reply(test.status)

			_, err := suite.srv.Client().UsersById("user").MessagesById("foo").Get(ctx, nil)
			assert.Error(t, err)

			// hacky, but not much better way to catch the retry count
			ce := clues.ToCore(err)
			assert.Equal(t, test.retryCount, ce.Values["retry_count"])
		})
	}
}
