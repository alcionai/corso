package m365

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/its"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type userIntegrationSuite struct {
	tester.Suite
	cli  client
	m365 its.M365IntgTestSetup
}

func TestUserIntegrationSuite(t *testing.T) {
	suite.Run(t, &userIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *userIntegrationSuite) SetupSuite() {
	var (
		t   = suite.T()
		err error
	)

	suite.m365 = its.GetM365(t)

	ctx, flush := tester.NewContext(t)
	defer flush()

	// will init the concurrency limiter
	suite.cli, err = NewM365Client(ctx, suite.m365.Acct)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *userIntegrationSuite) TestUsersCompat_HasNoInfo() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

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
			user:   suite.m365.User.ID,
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
			user:      suite.m365.User.ID,
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
			user: suite.m365.User.ID,
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
				require.NotNil(t, info)
				assert.Contains(t, info.ErrGetMailBoxSetting, api.ErrMailBoxNotFound)
			},
			// may seem odd, but we assume the user themselves
			// has already been vetted, which turns this into a
			// notFound error in the same way a mailboxNotFound
			// is handled.
			expectErr: require.NoError,
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
			userID:    suite.m365.User.ID,
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
		suite.Run(run.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			user, err := suite.cli.UserAssignedLicenses(ctx, run.userID)
			run.expectErr(t, err, clues.ToCore(err))
			assert.Equal(t, run.expect, user)
		})
	}
}
