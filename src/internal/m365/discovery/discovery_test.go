package discovery_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/discovery"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type DiscoveryIntgSuite struct {
	tester.Suite
}

func TestDiscoveryIntgSuite(t *testing.T) {
	suite.Run(t, &DiscoveryIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *DiscoveryIntgSuite) SetupSuite() {
	graph.InitializeConcurrencyLimiter(4)
}

func (suite *DiscoveryIntgSuite) TestUsers() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
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

func (suite *DiscoveryIntgSuite) TestUsers_InvalidCredentials() {
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
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

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

func (suite *DiscoveryIntgSuite) TestSites() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
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

func (suite *DiscoveryIntgSuite) TestSites_InvalidCredentials() {
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
				t          = suite.T()
				a          = test.acct(t)
				errs       = fault.New(true)
				ctx, flush = tester.NewContext(t)
			)

			defer flush()

			sites, err := discovery.Sites(ctx, a, errs)
			assert.Empty(t, sites, "returned some sites")
			assert.NotNil(t, err)
		})
	}
}

func (suite *DiscoveryIntgSuite) TestUserInfo() {
	t := suite.T()
	acct := tester.NewM365Account(t)

	creds, err := acct.M365Config()
	require.NoError(t, err)

	cli, err := api.NewClient(creds)
	require.NoError(t, err)

	uapi := cli.Users()

	table := []struct {
		name      string
		user      string
		expect    *api.UserInfo
		expectErr require.ErrorAssertionFunc
	}{
		{
			name: "standard test user",
			user: tester.M365UserID(t),
			expect: &api.UserInfo{
				ServicesEnabled: map[path.ServiceType]struct{}{
					path.ExchangeService: {},
					path.OneDriveService: {},
				},
				Mailbox: api.MailboxInfo{
					Purpose:              "user",
					ErrGetMailBoxSetting: nil,
				},
			},
			expectErr: require.NoError,
		},
		{
			name: "user does not exist",
			user: uuid.NewString(),
			expect: &api.UserInfo{
				ServicesEnabled: map[path.ServiceType]struct{}{},
				Mailbox:         api.MailboxInfo{},
			},
			expectErr: require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result, err := discovery.UserInfo(ctx, uapi, test.user)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, test.expect.ServicesEnabled, result.ServicesEnabled)
		})
	}
}

func (suite *DiscoveryIntgSuite) TestUserWithoutDrive() {
	t := suite.T()
	acct := tester.NewM365Account(t)
	userID := tester.M365UserID(t)

	table := []struct {
		name   string
		user   string
		expect *api.UserInfo
	}{
		{
			name: "user without drive and exchange",
			user: "a53c26f7-5100-4acb-a910-4d20960b2c19", // User: testevents@10rqc2.onmicrosoft.com
			expect: &api.UserInfo{
				ServicesEnabled: map[path.ServiceType]struct{}{},
				Mailbox: api.MailboxInfo{
					ErrGetMailBoxSetting: []error{api.ErrMailBoxSettingsNotFound},
				},
			},
		},
		{
			name: "user with drive and exchange",
			user: userID,
			expect: &api.UserInfo{
				ServicesEnabled: map[path.ServiceType]struct{}{
					path.ExchangeService: {},
					path.OneDriveService: {},
				},
				Mailbox: api.MailboxInfo{
					Purpose:              "user",
					ErrGetMailBoxSetting: []error{},
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result, err := discovery.GetUserInfo(ctx, acct, test.user, fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect.ServicesEnabled, result.ServicesEnabled)
			assert.Equal(t, test.expect.Mailbox.ErrGetMailBoxSetting, result.Mailbox.ErrGetMailBoxSetting)
			assert.Equal(t, test.expect.Mailbox.Purpose, result.Mailbox.Purpose)
		})
	}
}
