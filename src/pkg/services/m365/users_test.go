package m365

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type userIntegrationSuite struct {
	tester.Suite
	acct account.Account
}

func TestUserIntegrationSuite(t *testing.T) {
	suite.Run(t, &userIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *userIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	suite.acct = tconfig.NewM365Account(suite.T())
}

func (suite *userIntegrationSuite) TestUsers() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	users, err := Users(ctx, suite.acct, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, users)

	for _, u := range users {
		suite.Run("user_"+u.ID, func() {
			t := suite.T()

			assert.NotEmpty(t, u.ID)
			assert.NotEmpty(t, u.PrincipalName)
			assert.NotEmpty(t, u.Name)
			assert.NotEmpty(t, u.Info)
		})
	}
}

func (suite *userIntegrationSuite) TestUsersCompat_HasNoInfo() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tconfig.NewM365Account(suite.T())

	users, err := UsersCompatNoInfo(ctx, acct)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, users)

	for _, u := range users {
		suite.Run("user_"+u.ID, func() {
			t := suite.T()

			assert.NotEmpty(t, u.ID)
			assert.NotEmpty(t, u.PrincipalName)
			assert.NotEmpty(t, u.Name)
		})
	}
}

func (suite *userIntegrationSuite) TestUserHasMailbox() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct = tconfig.NewM365Account(t)
		uid  = tconfig.M365UserID(t)
	)

	enabled, err := UserHasMailbox(ctx, acct, uid)
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, enabled)
}

func (suite *userIntegrationSuite) TestUserHasDrive() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct = tconfig.NewM365Account(t)
		uid  = tconfig.M365UserID(t)
	)

	enabled, err := UserHasDrives(ctx, acct, uid)
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, enabled)
}

func (suite *userIntegrationSuite) TestUsers_InvalidCredentials() {
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

			users, err := Users(ctx, test.acct(t), fault.New(true))
			assert.Empty(t, users, "returned some users")
			assert.NotNil(t, err)
		})
	}
}

func (suite *userIntegrationSuite) TestGetUserInfo() {
	table := []struct {
		name      string
		user      string
		expect    *api.UserInfo
		expectErr require.ErrorAssertionFunc
	}{
		{
			name: "standard test user",
			user: tconfig.M365UserID(suite.T()),
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

			result, err := GetUserInfo(ctx, suite.acct, test.user)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, test.expect.ServicesEnabled, result.ServicesEnabled)
		})
	}
}

func (suite *userIntegrationSuite) TestGetUserInfo_userWithoutDrive() {
	userID := tconfig.M365UserID(suite.T())

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

			result, err := GetUserInfo(ctx, suite.acct, test.user)
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect.ServicesEnabled, result.ServicesEnabled)
			assert.Equal(t, test.expect.Mailbox.ErrGetMailBoxSetting, result.Mailbox.ErrGetMailBoxSetting)
			assert.Equal(t, test.expect.Mailbox.Purpose, result.Mailbox.Purpose)
		})
	}
}
