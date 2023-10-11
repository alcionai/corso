package m365

import (
	"fmt"
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

func (suite *userIntegrationSuite) TestUsersCompat_HasNoInfo() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

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
	acct := tconfig.NewM365Account(t)
	userID := tconfig.M365UserID(t)

	table := []struct {
		name   string
		user   string
		expect bool
	}{
		{
			name:   "user with no mailbox",
			user:   "a53c26f7-5100-4acb-a910-4d20960b2c19", // User: testevents@10rqc2.onmicrosoft.com
			expect: false,
		},
		{
			name:   "user with mailbox",
			user:   userID,
			expect: true,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			enabled, err := UserHasMailbox(ctx, acct, test.user)
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect, enabled)
		})
	}
}

func (suite *userIntegrationSuite) TestUserHasDrive() {
	t := suite.T()
	acct := tconfig.NewM365Account(t)
	userID := tconfig.M365UserID(t)

	table := []struct {
		name      string
		user      string
		expect    bool
		expectErr require.ErrorAssertionFunc
	}{
		{
			name:      "user without drive",
			user:      "a53c26f7-5100-4acb-a910-4d20960b2c19", // User: testevents@10rqc2.onmicrosoft.com
			expect:    false,
			expectErr: require.NoError,
		},
		{
			name:      "user with drive",
			user:      userID,
			expect:    true,
			expectErr: require.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			enabled, err := UserHasDrives(ctx, acct, test.user)
			test.expectErr(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect, enabled)
		})
	}
}

func (suite *userIntegrationSuite) TestUserGetMailboxInfo() {
	t := suite.T()
	acct := tconfig.NewM365Account(t)
	userID := tconfig.M365UserID(t)

	table := []struct {
		name      string
		user      string
		expect    func(t *testing.T, info api.MailboxInfo)
		expectErr require.ErrorAssertionFunc
	}{
		{
			name: "shared mailbox",
			user: "bb1a2049-3fc1-4fdc-93b8-7a14f63dd0db", // User: neha-test-shared-mailbox@10rqc2.onmicrosoft.com
			expect: func(t *testing.T, info api.MailboxInfo) {
				require.NotNil(t, info)
				assert.Equal(t, "shared", info.Purpose)
			},
			expectErr: require.NoError,
		},
		{
			name: "user mailbox",
			user: userID,
			expect: func(t *testing.T, info api.MailboxInfo) {
				require.NotNil(t, info)
				assert.Equal(t, "user", info.Purpose)
			},
			expectErr: require.NoError,
		},
		{
			name: "user with no mailbox",
			user: "a53c26f7-5100-4acb-a910-4d20960b2c19", // User: testevents@10rqc2.onmicrosoft.com
			expect: func(t *testing.T, info api.MailboxInfo) {
				require.NotNil(t, info)
				assert.Contains(t, info.ErrGetMailBoxSetting, api.ErrMailBoxNotFound)
			},
			expectErr: require.NoError,
		},
		{
			name: "invalid user",
			user: uuid.NewString(),
			expect: func(t *testing.T, info api.MailboxInfo) {
				mi := api.MailboxInfo{
					ErrGetMailBoxSetting: []error{},
				}

				require.Equal(t, mi, info)
			},
			expectErr: require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			info, err := UserGetMailboxInfo(ctx, acct, test.user)
			test.expectErr(t, err, clues.ToCore(err))
			test.expect(t, info)
		})
	}
}

func (suite *userIntegrationSuite) TestUser_GetByID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tconfig.NewM365Account(t)

	users, err := UsersCompatNoInfo(ctx, acct)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, users)

	for _, s := range users {
		suite.Run("user_"+s.ID, func() {
			t := suite.T()
			reconciliationNeeded, err := UserIsLicenseReconciliationNeeded(ctx, acct, s.ID)
			assert.NoError(t, err, clues.ToCore(err))
			plans, err := UserAssignedPlansCount(ctx, acct, s.ID)
			assert.NoError(t, err, clues.ToCore(err))

			lic, err := UserAssignedLicenses(ctx, acct, s.ID)
			assert.NoError(t, err, clues.ToCore(err))
			fmt.Printf("user: %v \n", reconciliationNeeded)
			fmt.Printf("plans: %v \n", plans)
			fmt.Printf("lic: %v \n", lic)
		})
	}
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

			users, err := UsersCompatNoInfo(ctx, test.acct(t))
			assert.Empty(t, users, "returned some users")
			assert.NotNil(t, err)
		})
	}
}
