package discovery_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/connector/discovery"
	"github.com/alcionai/corso/src/pkg/connector/discovery/api"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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

	var (
		t    = suite.T()
		acct = tester.NewM365Account(t)
		errs = fault.New(true)
	)

	creds, err := acct.M365Config()
	require.NoError(t, err)

	cli, err := api.NewClient(creds)
	require.NoError(t, err)

	users, err := discovery.Users(ctx, cli.Users(), errs)
	assert.NoError(t, err, clues.ToCore(err))

	ferrs := errs.Errors()
	assert.Nil(t, ferrs.Failure)
	assert.Empty(t, ferrs.Recovered)
	assert.NotEmpty(t, users)
}

func (suite *DiscoveryIntegrationSuite) TestUsers_InvalidCredentials() {
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
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()
			acct := test.acct(t)

			creds, err := acct.M365Config()
			require.NoError(t, err)

			cli, err := api.NewClient(creds)
			require.NoError(t, err)

			users, err := discovery.Users(ctx, cli.Users(), fault.New(true))
			assert.Empty(t, users, "returned some users")
			assert.NotNil(t, err)
		})
	}
}

func (suite *DiscoveryIntegrationSuite) TestSites() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(t)
		errs = fault.New(true)
	)

	sites, err := discovery.Sites(ctx, acct, errs)
	assert.NoError(t, err, clues.ToCore(err))

	ferrs := errs.Errors()
	assert.Nil(t, ferrs.Failure)
	assert.Empty(t, ferrs.Recovered)
	assert.NotEmpty(t, sites)
}

func (suite *DiscoveryIntegrationSuite) TestSites_InvalidCredentials() {
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
			var (
				t    = suite.T()
				a    = test.acct(t)
				errs = fault.New(true)
			)

			sites, err := discovery.Sites(ctx, a, errs)
			assert.Empty(t, sites, "returned some sites")
			assert.NotNil(t, err)
		})
	}
}

func (suite *DiscoveryIntegrationSuite) TestUserInfo() {
	t := suite.T()
	acct := tester.NewM365Account(t)
	userID := tester.M365UserID(t)

	creds, err := acct.M365Config()
	require.NoError(t, err)

	cli, err := api.NewClient(creds)
	require.NoError(t, err)

	uapi := cli.Users()

	table := []struct {
		name   string
		user   string
		expect *api.UserInfo
	}{
		{
			name: "standard test user",
			user: userID,
			expect: &api.UserInfo{
				DiscoveredServices: map[path.ServiceType]struct{}{
					path.ExchangeService: {},
					path.OneDriveService: {},
				},
			},
		},
		{
			name: "user does not exist",
			user: uuid.NewString(),
			expect: &api.UserInfo{
				DiscoveredServices: map[path.ServiceType]struct{}{
					path.OneDriveService: {}, // currently statically populated
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()

			result, err := discovery.UserInfo(ctx, uapi, test.user)
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect, result)
		})
	}
}
