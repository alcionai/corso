package discovery_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/connector/discovery"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
)

type DiscoveryIntegrationSuite struct {
	tester.Suite
}

func TestDiscoveryIntegrationSuite(t *testing.T) {
	suite.Run(t, &DiscoveryIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *DiscoveryIntegrationSuite) TestUsers() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	acct := tester.NewM365Account(t)
	errs := fault.New(true)

	users, err := discovery.Users(ctx, acct, errs)
	assert.NoError(t, err, clues.ToCore(err))

	ferrs := errs.Errors()
	assert.NoError(t, ferrs.Failure, clues.ToCore(ferrs.Failure))
	assert.Empty(t, ferrs.Recovered)

	assert.Less(t, 0, len(users))
}

func (suite *DiscoveryIntegrationSuite) TestUsers_InvalidCredentials() {
	ctx, flush := tester.NewContext()
	defer flush()

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
						AzureTenantID: "data",
					},
				)
				require.NoError(t, err, clues.ToCore(err))

				return a
			},
		},
		{
			name: "Empty Credentials",
			acct: func(t *testing.T) account.Account {
				// intentionally swallowing the error here
				a, _ := account.NewAccount(account.ProviderM365)
				return a
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			a := test.acct(t)
			errs := fault.New(true)
			users, err := discovery.Users(ctx, a, errs)

			assert.Empty(t, users, "returned some users")
			assert.NotNil(t, err)
			// TODO(ashmrtn): Uncomment when fault package is used in discovery API.
			// assert.NotNil(t, errs.Err())
		})
	}
}
