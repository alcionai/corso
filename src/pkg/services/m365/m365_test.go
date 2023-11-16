package m365

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type M365IntgSuite struct {
	tester.Suite
}

func TestM365IntgSuite(t *testing.T) {
	suite.Run(t, &M365IntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{}),
	})
}

func (suite *userIntegrationSuite) TestNewM365Client_invalidCredentials() {
	table := []struct {
		name string
		acct func(t *testing.T) account.Account
	}{
		{
			name: "Invalid Credentials",
			acct: func(t *testing.T) account.Account {
				a, err := account.NewAccount(
					account.ProviderM365,
					account.M365Config{
						M365: credentials.M365{
							AzureClientID:     "Test",
							AzureClientSecret: "without",
						},
						AzureTenantID: "creds",
					})
				require.NoError(t, err, clues.ToCore(err))

				return a
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := NewM365Client(ctx, test.acct(t))
			assert.Error(t, err, clues.ToCore(err))
		})
	}
}
