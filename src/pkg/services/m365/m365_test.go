package m365_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365"
)

type M365IntegrationSuite struct {
	tester.Suite
}

func TestM365IntegrationSuite(t *testing.T) {
	suite.Run(t, &M365IntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *M365IntegrationSuite) TestUsers() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tester.NewM365Account(suite.T())

	users, err := m365.Users(ctx, acct, fault.New(true))
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

func (suite *M365IntegrationSuite) TestUsersCompat_HasNoInfo() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tester.NewM365Account(suite.T())

	users, err := m365.UsersCompatNoInfo(ctx, acct)
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

func (suite *M365IntegrationSuite) TestGetUserInfo() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct = tester.NewM365Account(t)
		uid  = tester.M365UserID(t)
	)

	info, err := m365.GetUserInfo(ctx, acct, uid)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, info)
	require.NotEmpty(t, info)

	expectEnabled := map[path.ServiceType]struct{}{
		path.ExchangeService: {},
		path.OneDriveService: {},
	}

	assert.NotEmpty(t, info.ServicesEnabled)
	assert.NotEmpty(t, info.Mailbox)
	assert.Equal(t, expectEnabled, info.ServicesEnabled)
	assert.Equal(t, "user", info.Mailbox.Purpose)
}

func (suite *M365IntegrationSuite) TestUserHasMailbox() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct = tester.NewM365Account(t)
		uid  = tester.M365UserID(t)
	)

	enabled, err := m365.UserHasMailbox(ctx, acct, uid)
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, enabled)
}

func (suite *M365IntegrationSuite) TestUserHasDrive() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct = tester.NewM365Account(t)
		uid  = tester.M365UserID(t)
	)

	enabled, err := m365.UserHasDrives(ctx, acct, uid)
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, enabled)
}

func (suite *M365IntegrationSuite) TestSites() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tester.NewM365Account(suite.T())

	sites, err := m365.Sites(ctx, acct, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, sites)

	for _, s := range sites {
		suite.Run("site", func() {
			t := suite.T()
			assert.NotEmpty(t, s.WebURL)
			assert.NotEmpty(t, s.ID)
			assert.NotEmpty(t, s.DisplayName)
		})
	}
}
