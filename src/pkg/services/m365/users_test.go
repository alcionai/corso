package m365

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type userIntegrationSuite struct {
	tester.Suite
	cli client
}

func TestUserIntegrationSuite(t *testing.T) {
	suite.Run(t, &userIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *userIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tconfig.NewM365Account(t)

	var err error

	// will init the concurrency limiter
	suite.cli, err = NewM365Client(ctx, acct)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *userIntegrationSuite) TestUsersCompat_HasNoInfo() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	users, err := suite.cli.UsersCompatNoInfo(ctx)
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

			enabled, err := suite.cli.UserHasMailbox(ctx, test.user)
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect, enabled)
		})
	}
}

func (suite *userIntegrationSuite) TestUserHasDrive() {
	t := suite.T()
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

			enabled, err := suite.cli.UserHasDrives(ctx, test.user)
			test.expectErr(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect, enabled)
		})
	}
}

func (suite *userIntegrationSuite) TestUserGetMailboxInfo() {
	t := suite.T()
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

			info, err := suite.cli.UserGetMailboxInfo(ctx, test.user)
			test.expectErr(t, err, clues.ToCore(err))
			test.expect(t, info)
		})
	}
}

func (suite *userIntegrationSuite) TestUserAssignedLicenses() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)
	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	defer flush()

	runs := []struct {
		name      string
		userID    string
		expect    int
		expectErr require.ErrorAssertionFunc
	}{
		{
			name:      "user with no licenses",
			userID:    tconfig.UnlicensedM365UserID(t),
			expect:    0,
			expectErr: require.NoError,
		},
		{
			name:      "user with licenses",
			userID:    tconfig.M365UserID(t),
			expect:    2,
			expectErr: require.NoError,
		},
		{
			name:      "User does not exist",
			userID:    "fake",
			expect:    0,
			expectErr: require.Error,
		},
	}

	for _, run := range runs {
		t.Run(run.name, func(t *testing.T) {
			user, err := suite.cli.UserAssignedLicenses(ctx, run.userID)
			run.expectErr(t, err, clues.ToCore(err))
			assert.Equal(t, run.expect, user)
		})
	}
}
